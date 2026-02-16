package event

import (
	"context"
	"skillspark/internal/models"
	"skillspark/internal/storage/postgres/schema/organization"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func CreateTestEvent(
	t *testing.T,
	ctx context.Context,
	db *pgxpool.Pool,
) *models.Event {
	t.Helper()

	repo := NewEventRepository(db)

	organization := organization.CreateTestOrganization(t, ctx, db)

	input := &models.CreateEventDBInput{}
	ageMin := 8
	ageMax := 12

	ptrTitle := "เวิร์คช็อปหุ่นยนต์รุ่นเยาว์"
	ptrDesc := "เรียนรู้พื้นฐานของหุ่นยนต์ด้วยโปรเจ็กต์ LEGO Mindstorms ที่เน้นการลงมือปฏิบัติจริง สร้างและเขียนโปรแกรมหุ่นยนต์ของคุณเอง!"

	input.Body.Title_EN = "Junior Robotics Workshop"
	input.Body.Title_TH = &ptrTitle
	input.Body.Description_EN = "Learn the basics of robotics with hands-on LEGO Mindstorms projects. Build and program your own robots!"
	input.Body.Description_TH = &ptrDesc
	input.Body.OrganizationID = organization.ID
	input.Body.AgeRangeMin = &ageMin
	input.Body.AgeRangeMax = &ageMax
	input.Body.Category = []string{"science", "technology"}

	event, err := repo.CreateEvent(ctx, input, nil)

	require.NoError(t, err)
	require.NotNil(t, event)
	return event
}
