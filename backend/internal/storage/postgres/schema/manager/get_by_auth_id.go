package manager

import (
	"context"
	"errors"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *ManagerRepository) GetManagerByAuthID(ctx context.Context, authID string) (*models.Manager, error) {
	query, err := schema.ReadSQLBaseScript("manager/sql/get_by_auth_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, authID)
	var manager models.Manager
	err = row.Scan(&manager.ID, &manager.UserID, &manager.OrganizationID, &manager.Role,
		&manager.CreatedAt, &manager.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Manager", "auth_id", authID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch manager by auth_id: ", err.Error())
		return nil, &err
	}

	return &manager, nil
}
