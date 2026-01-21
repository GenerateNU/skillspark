package storage

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	locationSchema "skillspark/internal/storage/postgres/schema/location/getLocationByID"
	organizationSchema "skillspark/internal/storage/postgres/schema/organization"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type LocationRepository interface {
	GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, *errs.HTTPError)
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, org *models.Organization) *errs.HTTPError
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, *errs.HTTPError)
	GetAllOrganizations(ctx context.Context, offset, pageSize int) ([]models.Organization, int, *errs.HTTPError)
	UpdateOrganization(ctx context.Context, org *models.Organization) *errs.HTTPError
	DeleteOrganization(ctx context.Context, id uuid.UUID) *errs.HTTPError
}

type Repository struct {
	db       *pgxpool.Pool
	Location LocationRepository
	Organization OrganizationRepository
}

// Close closes the database connection pool
func (r *Repository) Close() error {
	r.db.Close()
	return nil
}

// GetDB returns the underlying pgxpool.Pool instance
func (r *Repository) GetDB() *pgxpool.Pool {
	return r.db
}

// NewRepository creates a new Repository instance with the given database pool
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db:       db,
		Location: locationSchema.NewLocationRepository(db),
		Organization: organizationSchema.NewOrganizationRepository(db),
	}
}
