package storage

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/school"
	"skillspark/internal/utils"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type LocationRepository interface {
	GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, *errs.HTTPError)
	CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, *errs.HTTPError)
}

type SchoolRepository interface {
	GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, *errs.HTTPError)
}

type ChildRepository interface {
	GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError)
	GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, *errs.HTTPError)
	UpdateChildByID(ctx context.Context, child *models.UpdateChildInput) (*models.Child, *errs.HTTPError)
	CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, *errs.HTTPError)
	DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError)
}

type Repository struct {
	db       *pgxpool.Pool
	Location LocationRepository
	School   SchoolRepository
	Child    ChildRepository
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
		Location: location.NewLocationRepository(db),
		School:   school.NewSchoolRepository(db),
		Child:    child.NewChildRepository(db),
	}
}
