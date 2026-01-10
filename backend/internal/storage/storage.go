package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repository provides methods to interact with the database
type Repository struct {
	db *pgxpool.Pool
}

// Close closes the database connection pool
func (r *Repository) Close() error {
	r.db.Close()
	return nil
}

// GetDB returns the underlying pgxpool.Pool instance
func (r *Repository) GetDB() *pgxpool.Pool {
	return r.db
}

// NewRepository creates a new Repository instance with the given database pool
func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}
