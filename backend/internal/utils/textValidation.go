package utils

import "github.com/jackc/pgx/v5/pgtype"

func TextPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}
