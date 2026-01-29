package registration

import "skillspark/internal/storage"

type Handler struct {
	RegistrationRepository storage.RegistrationRepository
	EventOccurrenceRepository storage.EventOccurrenceRepository
	GuardianRepository storage.GuardianRepository
	ChildRepository storage.ChildRepository
}

func NewHandler(registrationRepo storage.RegistrationRepository, childRepo storage.ChildRepository, guardianRepo storage.GuardianRepository, eventOccurrenceRepo storage.EventOccurrenceRepository) *Handler {
	return &Handler{
		RegistrationRepository: registrationRepo,
		ChildRepository:     childRepo,
		GuardianRepository: guardianRepo,
		EventOccurrenceRepository: eventOccurrenceRepo,
	}
}