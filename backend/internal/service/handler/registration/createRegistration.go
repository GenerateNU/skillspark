package registration

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"skillspark/internal/models"
	"time"
)

func (h *Handler) CreateRegistration(ctx context.Context, input *models.CreateRegistrationInput) (*models.CreateRegistrationOutput, error) {
	eventOccurrence, err := h.EventOccurrenceRepository.GetEventOccurrenceByID(ctx, input.Body.EventOccurrenceID, "en-US")
	if err != nil {
		return nil, err
	}

	if eventOccurrence.StartTime.Before(time.Now()) {
		return nil, errors.New("event occurrence has already started")
	}

	if eventOccurrence.CurrEnrolled >= eventOccurrence.MaxAttendees {
		return nil, errors.New("event occurrence has reached max registration")
	}

	child, err := h.ChildRepository.GetChildByID(ctx, input.Body.ChildID)
	if err != nil {
		return nil, err
	}

	if child.GuardianID != input.Body.GuardianID {
		return nil, errors.New("child does not belong to the specified guardian")
	}

	guardian, err := h.GuardianRepository.GetGuardianByID(ctx, input.Body.GuardianID)
	if err != nil {
		return nil, err
	}
	if guardian.StripeCustomerID == nil {
		return nil, errors.New("guardian must have a Stripe Customer ID before registering")
	}

	regData := &models.CreateRegistrationData{
		AcceptLanguage:    input.AcceptLanguage,
		ChildID:           input.Body.ChildID,
		GuardianID:        input.Body.GuardianID,
		EventOccurrenceID: input.Body.EventOccurrenceID,
		Status:            input.Body.Status,
	}

	registration, err := h.RegistrationRepository.CreateRegistration(ctx, regData)
	if err != nil {
		return nil, err
	}

	if h.NotificationService != nil && guardian.EmailNotifications {
		subject := "Registration Confirmed"
		body := fmt.Sprintf(
			"Your child has been successfully registered for %s on %s.",
			registration.Body.EventName,
			registration.Body.OccurrenceStartTime.Format("January 2, 2006 at 3:04 PM"),
		)
		if notifErr := h.NotificationService.SendNotification(ctx, &models.SendNotificationInput{
			NotificationType: models.NotificationTypeEmail,
			RecipientEmail:   &guardian.Email,
			Subject:          &subject,
			Body:             body,
		}); notifErr != nil {
			slog.Error("failed to send registration confirmation notification", "error", notifErr)
		}
	}

	return registration, nil
}
