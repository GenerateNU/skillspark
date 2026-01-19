package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (c *ChildRepository) GetChildByID(ctx context.Context, childID uuid.UUID) (*models.Child, *errs.HTTPError) {
	return nil, nil
}
