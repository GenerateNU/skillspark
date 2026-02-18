package review

import (
	"skillspark/internal/storage"
	translations "skillspark/internal/translation"
)

type Handler struct {
	RegistrationRepository storage.RegistrationRepository
	ReviewRepository       storage.ReviewRepository
	GuardianRepository     storage.GuardianRepository
	EventRepository        storage.EventRepository
	TranslateClient        translations.TranslationInterface
}

func NewHandler(registrationRepository storage.RegistrationRepository, reviewRepository storage.ReviewRepository, guardianRepository storage.GuardianRepository, eventRepository storage.EventRepository, translateClient translations.TranslationInterface) *Handler {
	return &Handler{
		RegistrationRepository: registrationRepository,
		ReviewRepository:       reviewRepository,
		GuardianRepository:     guardianRepository,
		EventRepository:        eventRepository,
		TranslateClient:        translateClient,
	}
}
