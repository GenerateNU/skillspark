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

func (r *ManagerRepository) GetManagerByOrgID(ctx context.Context, org_id uuid.UUID) (*models.Manager, *errs.HTTPError) {
	query, err := schema.ReadSQLBaseScript("manager/sql/get_by_org_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query, org_id)
	var manager models.Manager
	err = row.Scan(&manager.ID, &manager.UserID, &manager.OrganizationID, &manager.Role,
		&manager.CreatedAt, &manager.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Location", "id", org_id)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch location by id: ", err.Error())
		return nil, &err
	}

	return &manager, nil
}
