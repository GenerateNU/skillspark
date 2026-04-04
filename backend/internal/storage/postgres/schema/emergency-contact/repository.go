package emergencycontact

import "github.com/jackc/pgx/v5/pgxpool"

type EmergencyContactRepository struct {
	db *pgxpool.Pool
}

func NewEmergencyContactRepository(db *pgxpool.Pool) *EmergencyContactRepository {
	return &EmergencyContactRepository{db: db}
}
