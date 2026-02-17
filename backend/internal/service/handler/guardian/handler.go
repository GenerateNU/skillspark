package guardian

import (
	"skillspark/internal/storage" 
	"skillspark/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	db 	*pgxpool.Pool
	config             config.Supabase
}

func NewHandler(guardianRepository storage.GuardianRepository, db *pgxpool.Pool, config config.Supabase) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		db: db,
		config: config,
	}
}
