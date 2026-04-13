package emergencycontact

import (
	"skillspark/internal/storage"
)

type Handler struct {
	EmergencyContactRepository storage.EmergencyContactRepository
}

func NewHandler(emergencyContactRepository storage.EmergencyContactRepository) *Handler {
	return &Handler{
		EmergencyContactRepository: emergencyContactRepository,
	}
}
