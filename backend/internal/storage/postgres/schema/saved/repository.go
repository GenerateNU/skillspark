package saved

import "github.com/jackc/pgx/v5/pgxpool"

type SavedRepository struct {
	db *pgxpool.Pool
}

func NewSavedRepository(db *pgxpool.Pool) *SavedRepository {
	return &SavedRepository{db: db}
}
