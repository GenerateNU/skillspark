package guardian

import (
	"context"
	"log/slog"

	"skillspark/internal/auth"
	"skillspark/internal/models"
)

func (h *Handler) DeleteGuardian(ctx context.Context, input *models.DeleteGuardianInput) (*models.Guardian, error) {
	// transaction so that database guardian and Supabase auth user are always deleted together
	tx, txErr := h.db.Begin(ctx)
	
	defer tx.Conn().Close(ctx)

	if txErr != nil {
		slog.Error("Failed to start transaction: " + txErr.Error())
		return nil, txErr
	}

	guardian, err := h.GuardianRepository.DeleteGuardian(ctx, input.ID)
	if err != nil {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			slog.Error("Failed to roll back: " + rollBackErr.Error())
			return nil, rollBackErr
		}
		return nil, err
	}

	// delete Supabase auth user
	deleteErr := auth.SupabaseDeleteUser(&h.config, guardian.AuthID)
	if deleteErr != nil {
		rollBackErr := tx.Rollback(ctx)
		if rollBackErr != nil {
			slog.Error("Failed to roll back: " + rollBackErr.Error())
			return nil, rollBackErr
		}
		slog.Error("Roll back: " + deleteErr.Error())
		return nil, deleteErr
	}

	commitErr := tx.Commit(ctx)
	if commitErr != nil {
		slog.Error("Failed to commit: " + commitErr.Error())
		return nil, commitErr
	}
	return guardian, nil
}
