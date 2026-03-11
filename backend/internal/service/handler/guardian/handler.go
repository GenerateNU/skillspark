package guardian

import (
	"skillspark/internal/config"
	"skillspark/internal/storage"
	"skillspark/internal/stripeClient"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	GuardianRepository storage.GuardianRepository
	StripeClient       stripeClient.StripeClientInterface
	db                 *pgxpool.Pool
	config             config.Supabase
}

func NewHandler(guardianRepository storage.GuardianRepository, db *pgxpool.Pool, sc stripeClient.StripeClientInterface, config config.Supabase) *Handler {
	return &Handler{
		GuardianRepository: guardianRepository,
		db:                 db,
		config:             config,
		StripeClient:       sc,
	}
}
