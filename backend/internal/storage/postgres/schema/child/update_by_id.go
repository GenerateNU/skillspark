package child

import (
	"context"
	"errors"
	"net/http"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema"
	"skillspark/internal/storage/postgres/testutil"
	"skillspark/internal/utils"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func (r *ChildRepository) UpdateChildByID(ctx context.Context, childID uuid.UUID, child *models.UpdateChildInput) (*models.Child, error) {

	query, err := schema.ReadSQLBaseScript("child/sql/update_by_id.sql")
	if err != nil {
		errr := errs.InternalServerError("Failed to read base query: ", err.Error())
		return nil, &errr
	}

	row := r.db.QueryRow(
		ctx,
		query,
		child.Body.Name,
		child.Body.SchoolID,
		child.Body.BirthMonth,
		child.Body.BirthYear,
		child.Body.Interests,
		child.Body.GuardianID,
		childID)

	var updatedChild models.Child
	err = row.Scan(
		&updatedChild.ID,
		&updatedChild.Name,
		&updatedChild.SchoolID,
		&updatedChild.SchoolName,
		&updatedChild.BirthMonth,
		&updatedChild.BirthYear,
		&updatedChild.Interests,
		&updatedChild.GuardianID,
		&updatedChild.CreatedAt,
		&updatedChild.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err := errs.NotFound("Child", "id", childID)
			return nil, &err
		}
		err := errs.InternalServerError("Failed to fetch child by id: ", err.Error())
		return nil, &err
	}

	return &updatedChild, nil
}

func TestChildRepository_UpdateChildByID_NotFound(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	testDB := testutil.SetupTestDB(t)
	repo := NewChildRepository(testDB)
	ctx := context.Background()

	nonExistentID := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	input := &models.UpdateChildInput{}
	input.Body.Name = utils.PtrString("Updated Name")
	input.Body.BirthMonth = utils.PtrInt(5)
	input.Body.BirthYear = utils.PtrInt(2019)

	child, err := repo.UpdateChildByID(ctx, nonExistentID, input)

	assert.Nil(t, child)
	assert.NotNil(t, err)

	httpErr, ok := err.(*errs.HTTPError)
	assert.True(t, ok, "expected *errs.HTTPError")

	assert.Equal(t, http.StatusNotFound, httpErr.Code)
}
