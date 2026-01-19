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

func (r *ManagerRepository) DeleteManager(ctx context.Context, id uuid.UUID) (*models.Manager, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("location/sql/delete.sql")
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &err
	}

	row := r.db.QueryRow(ctx, query, id)

	var deletedManager models.Manager

	err = row.Scan(&deletedManager.ID, &deletedManager.UserID, &deletedManager.OrganizationID, &deletedManager.Role,
		&deletedManager.CreatedAt, &deletedManager.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &deletedManager, nil
		}
	}

	errvoid := errs.InternalServerError("Failed to Delete Manager ", err.Error())
	return nil, &errvoid
}
