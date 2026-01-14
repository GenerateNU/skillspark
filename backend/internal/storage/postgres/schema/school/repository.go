package school

import "github.com/jackc/pgx/v5/pgxpool"

type SchoolRepository struct {
	db *pgxpool.Pool
}

func NewSchoolRepository(db *pgxpool.Pool) *SchoolRepository {
	return &SchoolRepository{db: db}
}
