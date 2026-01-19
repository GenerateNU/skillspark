package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
)

func (c *ChildRepository) CreateChild(ctx context.Context, child *models.CreateChildInput) (*models.Child, *errs.HTTPError) {
	return nil, nil
}
