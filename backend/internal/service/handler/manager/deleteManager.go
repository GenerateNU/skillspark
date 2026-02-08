package manager

import (
	"context"
	"log/slog"
	"skillspark/internal/models"
	"skillspark/internal/auth"

	"github.com/google/uuid"
)

// DeleteManager handles DELETE /manager
func (h *Handler) DeleteManager(ctx context.Context, input *models.DeleteManagerInput) (*models.Manager, error) {
	id, _ := uuid.Parse(input.ID.String())	

	// transaction so that database guardian and Supabase auth user are always deleted together
	tx, txErr := h.db.Begin(ctx)
	if txErr != nil {
		slog.Error("Failed to start transaction: " + txErr.Error())
		return nil, txErr
	}

	manager, err := h.ManagerRepository.DeleteManager(ctx, id)
	if err != nil {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			slog.Error("Failed to roll back: " + rollBackErr.Error())
			return nil, rollBackErr
		}
		return nil, err
	}

	// delete Supabase auth user
	deleteErr := auth.SupabaseDeleteUser(&h.config, manager.AuthID)
	if deleteErr != nil {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			slog.Error("Failed to roll back: " + rollBackErr.Error())
			return nil, rollBackErr
		}
		return nil, deleteErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		slog.Error("Failed to commit: " + commitErr.Error())
		return nil, commitErr
	}
	return manager, nil
}
