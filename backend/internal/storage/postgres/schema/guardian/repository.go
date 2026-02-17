package guardian

import "github.com/jackc/pgx/v5/pgxpool"

type GuardianRepository struct {
	db *pgxpool.Pool
}

func NewGuardianRepository(db *pgxpool.Pool) *GuardianRepository {
	return &GuardianRepository{db: db}
}

func (r *GuardianRepository) GetDB() *pgxpool.Pool {
	return r.db
}