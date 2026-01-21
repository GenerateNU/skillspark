package organization

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	createorganization "skillspark/internal/storage/postgres/schema/organization/createorganization"
	deleteorganization "skillspark/internal/storage/postgres/schema/organization/deleteorganization"
	getallorganizationspaginated "skillspark/internal/storage/postgres/schema/organization/getallorganizationspaginated"
	getorganizationbyid "skillspark/internal/storage/postgres/schema/organization/getorganizationbyid"
	updateorganization "skillspark/internal/storage/postgres/schema/organization/updateorganization"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewOrganizationRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) CreateOrganization(ctx context.Context, org *models.Organization) *errs.HTTPError {
	return createorganization.Execute(ctx, r.db, org)
}

func (r *Repository) GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, *errs.HTTPError) {
	return getorganizationbyid.Execute(ctx, r.db, id)
}

func (r *Repository) GetAllOrganizations(ctx context.Context, offset, pageSize int) ([]models.Organization, int, *errs.HTTPError) {
	return getallorganizationspaginated.Execute(ctx, r.db, offset, pageSize, nil, nil)
}

func (r *Repository) UpdateOrganization(ctx context.Context, org *models.Organization) *errs.HTTPError {
	return updateorganization.Execute(ctx, r.db, org)
}

func (r *Repository) DeleteOrganization(ctx context.Context, id uuid.UUID) *errs.HTTPError {
	return deleteorganization.Execute(ctx, r.db, id)
}