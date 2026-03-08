package registration

import (
	"skillspark/internal/notification"
	"skillspark/internal/storage"
)

type Handler struct {
	RegistrationRepository    storage.RegistrationRepository
	EventOccurrenceRepository storage.EventOccurrenceRepository
	GuardianRepository        storage.GuardianRepository
	ChildRepository           storage.ChildRepository
	NotificationService       *notification.Service
}

func NewHandler(registrationRepo storage.RegistrationRepository, childRepo storage.ChildRepository, guardianRepo storage.GuardianRepository, eventOccurrenceRepo storage.EventOccurrenceRepository, notifService *notification.Service) *Handler {
	return &Handler{
		RegistrationRepository:    registrationRepo,
		ChildRepository:           childRepo,
		GuardianRepository:        guardianRepo,
		EventOccurrenceRepository: eventOccurrenceRepo,
		NotificationService:       notifService,
	}
}
