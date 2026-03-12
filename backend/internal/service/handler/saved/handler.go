package saved

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	SavedRepository storage.SavedRepository
	db              *pgxpool.Pool
	config          config.Supabase
}

func NewHandler(savedRepository storage.SavedRepository, db *pgxpool.Pool, config config.Supabase) *Handler {
	return &Handler{
		SavedRepository: savedRepository,
		db:              db,
		config:          config,
	}
}
