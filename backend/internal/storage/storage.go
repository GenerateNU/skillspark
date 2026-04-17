package storage

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/child"
	emergencycontact "skillspark/internal/storage/postgres/schema/emergency-contact"
	"skillspark/internal/storage/postgres/schema/event"
	eventoccurrence "skillspark/internal/storage/postgres/schema/event-occurrence"
	"skillspark/internal/storage/postgres/schema/guardian"
	"skillspark/internal/storage/postgres/schema/location"
	"skillspark/internal/storage/postgres/schema/manager"
	notification "skillspark/internal/storage/postgres/schema/notification"
	"skillspark/internal/storage/postgres/schema/organization"
	"skillspark/internal/storage/postgres/schema/recommendation"
	"skillspark/internal/storage/postgres/schema/registration"
	"skillspark/internal/storage/postgres/schema/review"
	"skillspark/internal/storage/postgres/schema/saved"
	"skillspark/internal/storage/postgres/schema/school"
	"skillspark/internal/storage/postgres/schema/user"
	"skillspark/internal/utils"
	"time"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type LocationRepository interface {
	GetLocationByID(ctx context.Context, id uuid.UUID) (*models.Location, error)
	CreateLocation(ctx context.Context, location *models.CreateLocationInput) (*models.Location, error)
	GetAllLocations(ctx context.Context, pagination utils.Pagination) ([]models.Location, error)
	GetLocationByOrganizationID(ctx context.Context, orgID uuid.UUID) (*models.Location, error)
}

type SchoolRepository interface {
	CreateSchool(ctx context.Context, school *models.CreateSchoolInput) (*models.School, error)
	GetAllSchools(ctx context.Context, pagination utils.Pagination) ([]models.School, error)
}

type OrganizationRepository interface {
	CreateOrganization(ctx context.Context, org *models.CreateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error)
	GetOrganizationByID(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.Organization, error)
	GetAllOrganizations(ctx context.Context, pagination utils.Pagination, AcceptLanguage string) ([]models.Organization, error)
	UpdateOrganization(ctx context.Context, org *models.UpdateOrganizationDBInput, PfpS3Key *string) (*models.Organization, error)
	DeleteOrganization(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.Organization, error)
	GetEventOccurrencesByOrganizationID(ctx context.Context, organization_id uuid.UUID, AcceptLanguage string) ([]models.EventOccurrence, error)
	SetStripeAccountID(ctx context.Context, orgID uuid.UUID, stripeAccountID string) (*models.Organization, error)
	SetStripeAccountStatus(ctx context.Context, stripeAccountID string, activated bool) (*models.Organization, error)
}

type ManagerRepository interface {
	GetManagerByID(ctx context.Context, id uuid.UUID) (*models.Manager, error)
	GetManagerByUserID(ctx context.Context, userID uuid.UUID) (*models.Manager, error)
	GetManagerByOrgID(ctx context.Context, org_id uuid.UUID) (*models.Manager, error)
	GetManagerByAuthID(ctx context.Context, authID string) (*models.Manager, error)
	DeleteManager(ctx context.Context, id uuid.UUID, tx pgx.Tx) (*models.Manager, error)
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
	SetStripeCustomerID(ctx context.Context, guardianID uuid.UUID, stripeCustomerID string) (*models.Guardian, error)
	DeleteGuardian(ctx context.Context, id uuid.UUID, tx pgx.Tx) (*models.Guardian, error)
	GetGuardianNotificationPreferences(ctx context.Context, ids []uuid.UUID) (map[uuid.UUID]models.GuardianNotificationPreferences, error)
}

type EventRepository interface {
	CreateEvent(ctx context.Context, location *models.CreateEventDBInput, HeaderImageS3Key *string) (*models.Event, error)
	UpdateEvent(ctx context.Context, location *models.UpdateEventDBInput, HeaderImageS3Key *string) (*models.Event, error)
	DeleteEvent(ctx context.Context, id uuid.UUID) error
	GetEventOccurrencesByEventID(ctx context.Context, event_id uuid.UUID, AcceptLanguage string) ([]models.EventOccurrence, error)
	GetEventByID(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.Event, error)
}

type ChildRepository interface {
	GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
	GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, error)
	UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, error)
	CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, error)
	DeleteChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, error)
}

type EventOccurrenceRepository interface {
	GetAllEventOccurrences(ctx context.Context, pagination utils.Pagination, AcceptLanguage string, filters models.GetAllEventOccurrencesFilter) ([]models.EventOccurrence, error)
	GetTrendingEventOccurrences(ctx context.Context, input *models.GetTrendingEventOccurrencesInput) ([]models.EventOccurrence, error)
	GetEventOccurrenceByID(ctx context.Context, id uuid.UUID, AcceptLanguage string) (*models.EventOccurrence, error)
	CreateEventOccurrence(ctx context.Context, input *models.CreateEventOccurrenceInput) (*models.EventOccurrence, error)
	UpdateEventOccurrence(ctx context.Context, input *models.UpdateEventOccurrenceInput, tx *pgx.Tx) (*models.EventOccurrence, error)
	CancelEventOccurrence(ctx context.Context, id uuid.UUID) error
}

type RegistrationRepository interface {
	CreateRegistration(ctx context.Context, input *models.CreateRegistrationWithPaymentData) (*models.CreateRegistrationOutput, error)
	GetRegistrationByID(ctx context.Context, input *models.GetRegistrationByIDInput, tx *pgx.Tx) (*models.GetRegistrationByIDOutput, error)
	GetRegistrationByPaymentIntentID(ctx context.Context, paymentIntentID string, AcceptLanguage string) (*models.Registration, error)
	GetRegistrationsByChildID(ctx context.Context, input *models.GetRegistrationsByChildIDInput) (*models.GetRegistrationsByChildIDOutput, error)
	GetRegistrationsByGuardianID(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error)
	GetRegistrationsByEventOccurrenceID(ctx context.Context, input *models.GetRegistrationsByEventOccurrenceIDInput) (*models.GetRegistrationsByEventOccurrenceIDOutput, error)
	GetRegistrationsForCapture(ctx context.Context, startWindow time.Time, endWindow time.Time) ([]models.Registration, error)
	UpdateRegistration(ctx context.Context, input *models.UpdateRegistrationInput) (*models.UpdateRegistrationOutput, error)
	CancelRegistration(ctx context.Context, input *models.CancelRegistrationInput) (*models.CancelRegistrationOutput, error)
	UpdateRegistrationPaymentStatus(ctx context.Context, input *models.UpdateRegistrationPaymentStatusInput) (*models.UpdateRegistrationPaymentStatusOutput, error)
}

type ReviewRepository interface {
	CreateReview(ctx context.Context, input *models.CreateReviewDBInput) (*models.Review, error)
	GetReviewsByGuardianID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error)
	GetReviewsByEventID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination, sortBy string) ([]models.Review, error)
	GetReviewsByOrganizationID(ctx context.Context, id uuid.UUID, AcceptLanguage string, pagination utils.Pagination) ([]models.Review, error)
	DeleteReview(ctx context.Context, id uuid.UUID) error
	GetAggregateReviews(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error)
	GetAggregateReviewsForOrganization(ctx context.Context, id uuid.UUID) (*models.ReviewAggregate, error)
	GetEventReviewsForOrganization(ctx context.Context, id uuid.UUID, pagination utils.Pagination, AcceptLanguage string, sortBy string) ([]models.SimpleReviewAggregate, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.CreateUserInput) (*models.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UpdateUserInput) (*models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (*models.User, error)
	UsernameExists(ctx context.Context, username string) (bool, error)
}

type NotificationRepository interface {
	CreateScheduledNotification(ctx context.Context, input *models.CreateScheduledNotificationInput) (*models.Notification, error)
	GetPendingNotifications(ctx context.Context) ([]models.Notification, error)
	UpdateNotificationStatus(ctx context.Context, id uuid.UUID, status models.NotificationStatus) (*models.Notification, error)
}

type SavedRepository interface {
	CreateSaved(ctx context.Context, saved *models.CreateSavedInput) (*models.Saved, error)
	DeleteSaved(ctx context.Context, id uuid.UUID) error
	GetByGuardianID(ctx context.Context, user_id uuid.UUID, pagination utils.Pagination, AcceptLanguage string) ([]models.Saved, error)
}

type EmergencyContactRepository interface {
	CreateEmergencyContact(ctx context.Context, emergencyContact *models.CreateEmergencyContactInput) (*models.CreateEmergencyContactOutput, error)
	UpdateEmergencyContact(ctx context.Context, emergencyContact *models.UpdateEmergencyContactInput) (*models.UpdateEmergencyContactOutput, error)
	GetEmergencyContactByGuardianID(ctx context.Context, guardian_id uuid.UUID) ([]*models.EmergencyContact, error)
	DeleteEmergencyContact(ctx context.Context, guardian_id uuid.UUID) (*models.DeleteEmergencyContactOutput, error)
}

type RecommendationRepository interface {
	GetRecommendationsByChildID(ctx context.Context, childInterests []string, childBirthYear int, acceptLanguage string, pagination utils.Pagination, filters models.RecommendationFilters) ([]models.Event, error)
}

type Repository struct {
	db               *pgxpool.Pool
	Location         LocationRepository
	Organization     OrganizationRepository
	School           SchoolRepository
	Manager          ManagerRepository
	Guardian         GuardianRepository
	Event            EventRepository
	Child            ChildRepository
	EventOccurrence  EventOccurrenceRepository
	Registration     RegistrationRepository
	Review           ReviewRepository
	User             UserRepository
	Notification     NotificationRepository
	Saved            SavedRepository
	EmergencyContact EmergencyContactRepository
	Recommendation   RecommendationRepository
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
		db:               db,
		Location:         location.NewLocationRepository(db),
		Organization:     organization.NewOrganizationRepository(db),
		School:           school.NewSchoolRepository(db),
		Manager:          manager.NewManagerRepository(db),
		Guardian:         guardian.NewGuardianRepository(db),
		Event:            event.NewEventRepository(db),
		Child:            child.NewChildRepository(db),
		EventOccurrence:  eventoccurrence.NewEventOccurrenceRepository(db),
		User:             user.NewUserRepository(db),
		Registration:     registration.NewRegistrationRepository(db),
		Review:           review.NewReviewRepository(db),
		Notification:     notification.NewNotificationRepository(db),
		Saved:            saved.NewSavedRepository(db),
		EmergencyContact: emergencycontact.NewEmergencyContactRepository(db),
		Recommendation:   recommendation.NewRecommendationRepository(db),
	}
}
