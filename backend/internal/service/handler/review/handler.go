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
	OrganizationRepository storage.OrganizationRepository
	TranslateClient        translations.TranslationInterface
}

func NewHandler(registrationRepository storage.RegistrationRepository, reviewRepository storage.ReviewRepository, guardianRepository storage.GuardianRepository, eventRepository storage.EventRepository, organizationRepository storage.OrganizationRepository, translateClient translations.TranslationInterface) *Handler {
	return &Handler{
		RegistrationRepository: registrationRepository,
		ReviewRepository:       reviewRepository,
		GuardianRepository:     guardianRepository,
		EventRepository:        eventRepository,
		OrganizationRepository: organizationRepository,
		TranslateClient:        translateClient,
	}
}
