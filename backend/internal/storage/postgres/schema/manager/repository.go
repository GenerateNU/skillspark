package manager

import "github.com/jackc/pgx/v5/pgxpool"

type ManagerRepository struct {
	db *pgxpool.Pool
}

func NewManagerRepository(db *pgxpool.Pool) *ManagerRepository {
	return &ManagerRepository{db: db}
}