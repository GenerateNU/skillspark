package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"

	"github.com/google/uuid"
)

func (c *ChildRepository) GetChildrenByParentID(ctx context.Context, parentID uuid.UUID) ([]models.Child, *errs.HTTPError) {
	return nil, nil
}
