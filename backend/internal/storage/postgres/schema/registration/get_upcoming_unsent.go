package registration

import (
	"context"
	_ "embed"
	"time"

	"skillspark/internal/errs"
	"skillspark/internal/models"
)

//go:embed sql/get_upcoming_unsent.sql
var getUpcomingUnsentSQL string

type GetUpcomingUnsentRegistrationsInput struct {
	WindowStart time.Time `json:"window_start"`
	WindowEnd   time.Time `json:"window_end"`
}

type RegistrationWithGuardian struct {
	Registration  models.Registration `json:"registration"`
	GuardianEmail *string             `json:"guardian_email"`
	GuardianName  *string             `json:"guardian_name"`
	LineAccountID *string             `json:"line_account_id"`
}

type GetUpcomingUnsentRegistrationsOutput struct {
	Body struct {
		Registrations []RegistrationWithGuardian `json:"registrations"`
	} `json:"body"`
}

// GetUpcomingUnsentRegistrations retrieves registrations for events starting within a time window that haven't had a reminder sent.
func (r *RegistrationRepository) GetUpcomingUnsentRegistrations(ctx context.Context, input *GetUpcomingUnsentRegistrationsInput) (*GetUpcomingUnsentRegistrationsOutput, error) {
	rows, err := r.db.Query(ctx, getUpcomingUnsentSQL, input.WindowStart, input.WindowEnd)
	if err != nil {
		err := errs.InternalServerError("Failed to query upcoming unsent registrations: ", err.Error())
		return nil, &err
	}
	defer rows.Close()

	var output GetUpcomingUnsentRegistrationsOutput
	output.Body.Registrations = []RegistrationWithGuardian{}

	for rows.Next() {
		var item RegistrationWithGuardian
		// Ensure we scan into the correct fields matching the SELECT order
		err := rows.Scan(
			&item.Registration.ID,
			&item.Registration.ChildID,
			&item.Registration.GuardianID,
			&item.Registration.EventOccurrenceID,
			&item.Registration.Status,
			&item.Registration.CreatedAt,
			&item.Registration.UpdatedAt,
			&item.Registration.EventName,
			&item.Registration.OccurrenceStartTime,
			&item.Registration.ReminderSent,
			&item.GuardianEmail,
			&item.GuardianName,
			&item.LineAccountID,
		)
		if err != nil {
			err := errs.InternalServerError("Failed to scan registration row: ", err.Error())
			return nil, &err
		}
		output.Body.Registrations = append(output.Body.Registrations, item)
	}

	if rows.Err() != nil {
		err := errs.InternalServerError("Row iteration error: ", rows.Err().Error())
		return nil, &err
	}

	return &output, nil
}
