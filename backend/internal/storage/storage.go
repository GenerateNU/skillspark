package storage

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/manager"
	"skillspark/internal/storage/postgres/schema/organization"
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

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, org *models.CreateOrganizationInput, PfpS3Key *string) (*models.Organization, *errs.HTTPError)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, *errs.HTTPError)
	GetAllOrganizations(ctx context.Context, pagination utils.Pagination) ([]models.Organization, *errs.HTTPError)
	UpdateOrganization(ctx context.Context, org *models.UpdateOrganizationInput) (*models.Organization, *errs.HTTPError)
	DeleteOrganization(ctx context.Context, id uuid.UUID) (*models.Organization, *errs.HTTPError)
}

type ManagerRepository interface {
	GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, error)
	GetManagerByOrgID(ctx context.Context, org_id uuid.UUID) (*models.Manager, error)
	DeleteManager(ctx context.Context, id uuid.UUID) (*models.Manager, error)
	CreateManager(ctx context.Context, manager *models.CreateManagerInput) (*models.Manager, error)
	PatchManager(ctx context.Context, manager *models.PatchManagerInput) (*models.Manager, error)
}

type GuardianRepository interface {
	CreateGuardian(ctx context.Context, guardian *models.CreateGuardianInput) (*models.Guardian, error)
	GetGuardianByChildID(ctx context.Context, childID uuid.UUID) (*models.Guardian, error)
	GetGuardianByID(ctx context.Context, id uuid.UUID) (*models.Guardian, error)
	GetGuardianByUserID(ctx context.Context, userID uuid.UUID) (*models.Guardian, error)
	UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, error)
	DeleteGuardian(ctx context.Context, id uuid.UUID) (*models.Guardian, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, location *models.CreateEventInput, HeaderImageS3Key *string) (*models.Event, error)
	UpdateEvent(ctx context.Context, location *models.UpdateEventInput) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID) ([]models.EventOccurrence, error)
}

type ChildRepository interface {
	GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
	GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, error)
	UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, error)
	CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, error)
	DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
}

type EventOccurrenceRepository interface {
	GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination) ([]models.EventOccurrence, error)
	GetEventOccurrenceByID(ctx context.Context, id uuid.UUID) (*models.EventOccurrence, error)
	CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error)
}

type Repository struct {
	db              *pgxpool.Pool
	Location        LocationRepository
	Organization    OrganizationRepository
	School          SchoolRepository
	Manager         ManagerRepository
	Guardian        GuardianRepository
	Event           EventRepository
	Child           ChildRepository
	EventOccurrence EventOccurrenceRepository
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
		db:              db,
		Location:        location.NewLocationRepository(db),
		Organization:    organization.NewOrganizationRepository(db),
		School:          school.NewSchoolRepository(db),
		Manager:         manager.NewManagerRepository(db),
		Guardian:        guardian.NewGuardianRepository(db),
		Event:           event.NewEventRepository(db),
		Child:           child.NewChildRepository(db),
		EventOccurrence: eventoccurrence.NewEventOccurrenceRepository(db),
	}
}
