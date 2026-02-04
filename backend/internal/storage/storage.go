package storage

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/manager"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/schema/registration"
	"skillspark/internal/storage/postgres/schema/school"
	"skillspark/internal/storage/postgres/schema/user"
	"skillspark/internal/utils"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type LocationRepository interface {
	GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, error)
	GetAllLocations(ctx context.Context, pagination utils.Pagination) ([]models.Location, error)
}

type SchoolRepository interface {
	CreateSchool(ctx context.Context, school *models.CreateSchoolInput) (*models.School, error)
	GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error)
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, org *models.CreateOrganizationInput, PfpS3Key *string) (*models.Organization, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID) (*models.Organization, error)
	GetAllOrganizations(ctx context.Context, pagination utils.Pagination) ([]models.Organization, error)
	UpdateOrganization(ctx context.Context, org *models.UpdateOrganizationInput, PfpS3Key *string) (*models.Organization, error)
	DeleteOrganization(ctx context.Context, id uuid.UUID) (*models.Organization, error)
	GetEventOccurrencesByOrganizationID(ctx context.Context, organization_id uuid.UUID) ([]models.EventOccurrence, error)
}

type ManagerRepository interface {
	GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, error)
	GetManagerByUserID(ctx context.Context, userID uuid.UUID) (*models.Manager, error)
	GetManagerByOrgID(ctx context.Context, org_id uuid.UUID) (*models.Manager, error)
	GetManagerByAuthID(ctx context.Context, authID string) (*models.Manager, error)
	DeleteManager(ctx context.Context, id uuid.UUID) (*models.Manager, error)
	CreateManager(ctx context.Context, manager *models.CreateManagerInput) (*models.Manager, error)
	PatchManager(ctx context.Context, manager *models.PatchManagerInput) (*models.Manager, error)
}

type GuardianRepository interface {
	CreateGuardian(ctx context.Context, guardian *models.CreateGuardianInput) (*models.Guardian, error)
	GetGuardianByChildID(ctx context.Context, childID uuid.UUID) (*models.Guardian, error)
	GetGuardianByID(ctx context.Context, id uuid.UUID) (*models.Guardian, error)
	GetGuardianByUserID(ctx context.Context, userID uuid.UUID) (*models.Guardian, error)
	GetGuardianByAuthID(ctx context.Context, authID string) (*models.Guardian, error)
	UpdateGuardian(ctx context.Context, guardian *models.UpdateGuardianInput) (*models.Guardian, error)
	DeleteGuardian(ctx context.Context, id uuid.UUID) (*models.Guardian, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, location *models.CreateEventInput, HeaderImageS3Key *string) (*models.Event, error)
	UpdateEvent(ctx context.Context, location *models.UpdateEventInput, HeaderImageS3Key *string) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID) ([]models.EventOccurrence, error)
	GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error)
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
	UpdateEventOccurrence(ctx context.Context, input *models.UpdateEventOccurrenceInput) (*models.EventOccurrence, error)
	CancelEventOccurrence(ctx context.Context, id uuid.UUID) error
}

type RegistrationRepository interface {
	CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error)
	GetRegistrationByID(ctx context.Context, input *models.GetRegistrationByIDInput) (*models.GetRegistrationByIDOutput, error)
	GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error)
	GetRegistrationsByGuardianID(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error)
	GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error)
	UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error)
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
	Registration    RegistrationRepository
	User            UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.CreateUserInput) (*models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UpdateUserInput) (*models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (*models.User, error)
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
		User:            user.NewUserRepository(db),
		Registration:    registration.NewRegistrationRepository(db),
	}
}
