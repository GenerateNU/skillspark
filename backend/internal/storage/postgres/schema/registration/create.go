package registration

import (
	"context"
	"log/slog"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
)

func (r *RegistrationRepository) CreateRegistration(ctx context.Context, input *models.CreateRegistrationData) (*models.CreateRegistrationOutput, error) {
	tx, err := r.db.Begin(ctx)

	var titleEN string
	var titleTH *string

	if err != nil {
		return nil, err
	}

	query, err := schema.ReadSQLBaseScript("create.sql", SqlRegistrationFiles)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(ctx, query,
		input.ChildID,
		input.GuardianID,
		input.EventOccurrenceID,
		input.Status,
	)

	var createdRegistration models.CreateRegistrationOutput

	err = row.Scan(
		&createdRegistration.Body.ID,
		&createdRegistration.Body.ChildID,
		&createdRegistration.Body.GuardianID,
		&createdRegistration.Body.EventOccurrenceID,
		&createdRegistration.Body.Status,
		&createdRegistration.Body.CreatedAt,
		&createdRegistration.Body.UpdatedAt,
		&createdRegistration.Body.StripeCustomerID,
		&createdRegistration.Body.OrgStripeAccountID,
		&createdRegistration.Body.Currency,
		&createdRegistration.Body.PaymentIntentStatus,
		&createdRegistration.Body.CancelledAt,
		&createdRegistration.Body.StripePaymentIntentID,
		&createdRegistration.Body.TotalAmount,
		&createdRegistration.Body.ProviderAmount,
		&createdRegistration.Body.PlatformFeeAmount,
		&createdRegistration.Body.PaidAt,
		&createdRegistration.Body.StripePaymentMethodID,
		&titleEN,
		&titleTH,
		&createdRegistration.Body.OccurrenceStartTime,
	)

	if err != nil {
		errr := errs.InternalServerError("Failed to create registration: ", err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, &errr
	}

	incrementEventOccurrenceQuery, err := schema.ReadSQLBaseScript("change_event_occurrence_by.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, &errr
	}

	_, err = tx.Exec(ctx, incrementEventOccurrenceQuery, input.EventOccurrenceID, 1)
	if err != nil {
		errr := errs.InternalServerError("Failed to increment event occurrence attendee count: ", err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, &errr
	}

	if err := tx.Commit(ctx); err != nil {
		slog.Error("Failed to commit transaction: " + err.Error())
		if err := tx.Rollback(ctx); err != nil {
			slog.Error("Failed to rollback transaction: " + err.Error())
		}
		return nil, errs.InternalServerError("Failed to commit transaction: ", err.Error())
	}

	switch input.AcceptLanguage {
	case "th-TH":
		createdRegistration.Body.EventName = *titleTH
	case "en-US":
		createdRegistration.Body.EventName = titleEN
	}
	return &createdRegistration, nil
}
