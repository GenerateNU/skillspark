package manager

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	ManagerRepository  storage.ManagerRepository
	db *pgxpool.Pool
	config             config.Supabase
}

func NewHandler(managerRepository storage.ManagerRepository, db *pgxpool.Pool, config config.Supabase) *Handler {
	return &Handler{
		ManagerRepository:  managerRepository,
		db: db,
		config: config,
	}
}
