package child

import "github.com/jackc/pgx/v5/pgxpool"

type ChildRepository struct {
	db *pgxpool.Pool
}

func NewChildRepository(db *pgxpool.Pool) *ChildRepository {
	return &ChildRepository{db: db}
}
