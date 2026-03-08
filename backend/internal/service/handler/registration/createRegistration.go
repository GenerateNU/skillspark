package registration

import (
	"context"
	"fmt"
	"log/slog"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
	_, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID)
	if err != nil {
		return nil, errs.BadRequest("Invalid event_occurrence_id: event occurrence does not exist")
	}

	if _, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID); err != nil {
		return nil, errs.BadRequest("Invalid child_id: child does not exist")
	}

	_, err = h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID)
	if err != nil {
		return nil, errs.BadRequest("Invalid guardian_id: guardian does not exist")
	}

	registration, err := h.RegistrationRepository.CreateRegistration(ctx, input)
	if err != nil {
		return nil, err
	}

	if h.NotificationService != nil {
		subject := "Registration Confirmed"
		body := fmt.Sprintf(
			"Your child has been successfully registered for %s on %s.",
			registration.Body.EventName,
			registration.Body.OccurrenceStartTime.Format("January 2, 2006 at 3:04 PM"),
		)
		myEmail := "bobbypalazzi@gmail.com"
		if notifErr := h.NotificationService.SendNotification(ctx, &models.SendNotificationInput{
			NotificationType: models.NotificationTypeEmail,
			RecipientEmail:   &myEmail,
			Subject:          &subject,
			Body:             body,
		}); notifErr != nil {
			slog.Error("failed to send registration confirmation notification", "error", notifErr)
		}
	}

	return registration, nil
}
