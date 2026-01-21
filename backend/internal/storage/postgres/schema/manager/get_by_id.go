package manager

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *ManagerRepository) GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, error) {
	query, err := schema.ReadSQLBaseScript("manager/sql/get_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, id)
	var manager models.Manager
	err = row.Scan(&manager.ID, &manager.UserID, &manager.OrganizationID, &manager.Role,
		&manager.CreatedAt, &manager.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Location", "id", id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch manager by id: ", err.Error())
		return nil, &err
	}

	return &manager, nil
}
