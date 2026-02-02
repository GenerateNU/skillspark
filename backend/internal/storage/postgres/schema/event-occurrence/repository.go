package eventoccurrence

import "github.com/jackc/pgx/v5/pgxpool"

type EventOccurrenceRepository struct {
	db *pgxpool.Pool
}

func NewEventOccurrenceRepository(db *pgxpool.Pool) *EventOccurrenceRepository {
	return &EventOccurrenceRepository{db: db}
}
