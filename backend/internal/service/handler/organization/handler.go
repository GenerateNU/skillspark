package organization

import "skillspark/internal/storage"

type Handler struct {
	OrganizationRepository storage.OrganizationRepository
	LocationRepository     storage.LocationRepository
}

func NewHandler(orgRepo storage.OrganizationRepository, locRepo storage.LocationRepository) *Handler {
	return &Handler{
		OrganizationRepository: orgRepo,
		LocationRepository:     locRepo,
	}
}