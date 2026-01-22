package storage

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/school"
	"skillspark/internal/utils"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type LocationRepository interface {
	GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, error)
}

type SchoolRepository interface {
	GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, location *models.CreateEventInput) (*models.Event, error)
	UpdateEvent(ctx context.Context, location *models.UpdateEventInput) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
}

type ChildRepository interface {
	GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
	GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, error)
	UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, error)
	CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, error)
	DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
}

type Repository struct {
	db       *pgxpool.Pool
	Location LocationRepository
	School   SchoolRepository
	Event    EventRepository
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
		Event:    event.NewEventRepository(db),
		Child:    child.NewChildRepository(db),
	}
}
