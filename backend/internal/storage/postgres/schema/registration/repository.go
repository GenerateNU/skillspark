package registration

import "github.com/jackc/pgx/v5/pgxpool"

type RegistrationRepository struct {
	db *pgxpool.Pool
}

func NewRegistrationRepository(db *pgxpool.Pool) *RegistrationRepository {
	return &RegistrationRepository{db: db}
}
