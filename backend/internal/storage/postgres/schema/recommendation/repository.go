package recommendation

import (
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sql/*.sql
var SqlRecommendationFiles embed.FS

type RecommendationRepository struct {
	db *pgxpool.Pool
}

func NewRecommendationRepository(db *pgxpool.Pool) *RecommendationRepository {
	return &RecommendationRepository{db: db}
}
