package user

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/storage/postgres/schema"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (bool, error) {
	query, err := schema.ReadSQLBaseScript("get_by_username.sql", SqlUserFiles)
	if err != nil {
		err := errs.InternalServerError("Failed to read base query: ", err.Error())
		return false, &err
	}

	row := r.db.QueryRow(ctx, query, username)

	var id string
	err = row.Scan(&id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return false, nil
		}
		err := errs.InternalServerError("Failed to check username: ", err.Error())
		return false, &err
	}

	return true, nil
}
