package emergencycontact

import (
	"context"
	"fmt"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (h *Handler) DeleteEmergencyContact(ctx context.Context, id uuid.UUID) (*models.DeleteEmergencyContactOutput, error) {
	id, err := uuid.Parse(id.String())
	if err != nil {
		fmt.Println("failure at start")
		fmt.Println(err.Error())
		return nil, errs.BadRequest("Invalid ID format")
	}

	emergencyContact, httpErr := h.EmergencyContactRepository.DeleteEmergencyContact(ctx, id)
	if httpErr != nil {
		fmt.Println("failure at repo step")
		fmt.Println(httpErr.Error())
		return nil, httpErr
	}
	return emergencyContact, nil
}
