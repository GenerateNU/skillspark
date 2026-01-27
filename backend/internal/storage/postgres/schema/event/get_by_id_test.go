package event

import (
	"context"
	"testing"
	"skillspark/internal/storage/postgres/testutil"

	"github.com/stretchr/testify/assert"
	"github.com/google/uuid"
)

func TestEventRepository_GetEventById(t *testing.T){
	if testing.Short() {
		t.Skip("Skipping database test in short mode")
	}

	category_arr1 := []string{"science","technology"}
	category_arr2 := []string{"science"}

	testDB := testutil.SetupTestDB(t)
	repo := NewEventRepository(testDB)
	ctx := context.Background()

	// check that get by id works for 2 different events
	event1, err := repo.GetEventByID(ctx, uuid.MustParse("60000000-0000-0000-0000-000000000001"))
	assert.Nil(t, err)
	assert.NotNil(t, event1)
	assert.Equal(t, "Junior Robotics Workshop", event1.Title)
	assert.Equal(t, "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!", event1.Description)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), event1.OrganizationID)
	assert.Equal(t, 8, *event1.AgeRangeMin)
	assert.Equal(t, 12, *event1.AgeRangeMax)
	assert.Equal(t, category_arr1, event1.Category)
	assert.Equal(t, "events/robotics_workshop.jpg", *event1.HeaderImageS3Key)

	event2, err := repo.GetEventByID(ctx, uuid.MustParse("60000000-0000-0000-0000-000000000002"))
	assert.Nil(t, err)
	assert.NotNil(t, event2)
	assert.Equal(t, "Chemistry for Kids", event2.Title)
	assert.Equal(t, "Exciting chemistry experiments that are safe and fun. Discover reactions, make slime, and learn about molecules!", event2.Description)
	assert.Equal(t, uuid.MustParse("40000000-0000-0000-0000-000000000001"), event2.OrganizationID)
	assert.Equal(t, 7, *event2.AgeRangeMin)
	assert.Equal(t, 10, *event2.AgeRangeMax)
	assert.Equal(t, category_arr2, event2.Category)
	assert.Equal(t, "events/chemistry_kids.jpg", *event2.HeaderImageS3Key)
}