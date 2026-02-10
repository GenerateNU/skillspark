package guardianpaymentmethod

import "github.com/jackc/pgx/v5/pgxpool"

type GuardianPaymentMethodRepository struct {
	db *pgxpool.Pool
}

func NewGuardianPaymentMethodRepository(db *pgxpool.Pool) *GuardianPaymentMethodRepository {
	return &GuardianPaymentMethodRepository{db: db}
}