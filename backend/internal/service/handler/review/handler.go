package review

import (
	"skillspark/internal/storage"
)

type Handler struct {
	RegistrationRepository storage.RegistrationRepository
	ReviewRepository       storage.ReviewRepository
	GuardianRepository     storage.GuardianRepository
	EventRepository        storage.EventRepository
}

func NewHandler(registrationRepository storage.RegistrationRepository, reviewRepository storage.ReviewRepository, guardianRepository storage.GuardianRepository, eventRepository storage.EventRepository) *Handler {
	return &Handler{
		RegistrationRepository: registrationRepository,
		ReviewRepository:       reviewRepository,
		GuardianRepository:     guardianRepository,
		EventRepository:        eventRepository,
	}
}
