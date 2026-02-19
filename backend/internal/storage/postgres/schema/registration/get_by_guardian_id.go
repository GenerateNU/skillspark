package registration

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *RegistrationRepository) GetRegistrationsByGuardianID(ctx context.Context, input *models.GetRegistrationsByGuardianIDInput) (*models.GetRegistrationsByGuardianIDOutput, error) {
	query, err := schema.ReadSQLBaseScript("get_by_guardian_id.sql", SqlRegistrationFiles)
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	rows, err := r.db.Query(ctx, query, input.GuardianID)
	if err != nil {
		errr := errs.InternalServerError("Failed to get registrations by guardian id: ", err.Error())
		return nil, &errr
	}
	defer rows.Close()

	registrations, err := pgx.CollectRows(rows, scanRegistration)
	if err != nil {
		errr := errs.InternalServerError("Failed to collect registrations: ", err.Error())
		return nil, &errr
	}

	var output models.GetRegistrationsByGuardianIDOutput

	output.Body.Registrations = registrations

	return &output, nil
}
