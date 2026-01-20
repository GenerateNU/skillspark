package child

import (
	"context"
	"skillspark/internal/errs"
	"skillspark/internal/models"
	repomocks "skillspark/internal/storage/repo-mocks"
	"skillspark/internal/utils"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_GetChildByID(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		mockSetup func(*repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name: "successful get child by id - Emily",
			id:   "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("GetChildByID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).
					Return(&models.Child{
						ID:         uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						Name:       "Emily Johnson",
						SchoolID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						BirthMonth: 3,
						BirthYear:  2016,
						Interests:  []string{"science", "technology", "math"},
						GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "child not found",
			id:   "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("GetChildByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()
			input := &models.ChildIDInput{ID: uuid.MustParse(tt.id)}

			child, err := handler.GetChildByID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, child)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, child)
				assert.Equal(t, tt.id, child.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_GetChildrenByParentID(t *testing.T) {
	tests := []struct {
		name      string
		parentID  string
		mockSetup func(*repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name:     "children exist for guardian",
			parentID: "11111111-1111-1111-1111-111111111111",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("GetChildrenByParentID", mock.Anything, uuid.MustParse("11111111-1111-1111-1111-111111111111")).
					Return([]models.Child{
						{
							ID:         uuid.MustParse("30000000-0000-0000-0000-000000000001"),
							Name:       "Emily Johnson",
							SchoolID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
							BirthMonth: 3,
							BirthYear:  2016,
							Interests:  []string{"science", "technology", "math"},
							GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
							CreatedAt:  time.Now(),
							UpdatedAt:  time.Now(),
						},
					}, nil)
			},
			wantErr: false,
		},
		{
			name:     "no children found",
			parentID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("GetChildrenByParentID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "guardian_id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()
			input := &models.GuardianIDInput{ID: uuid.MustParse(tt.parentID)}

			children, err := handler.GetChildrenByParentID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, children)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, children)
				assert.Equal(t, 1, len(children))
				assert.Equal(t, "Emily Johnson", children[0].Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_UpdateChildByID(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.UpdateChildInput
		mockSetup func(*repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name: "successful update child",
			input: func() *models.UpdateChildInput {
				in := &models.UpdateChildInput{}
				in.Body.Name = utils.PtrString("Updated Emily")
				in.Body.BirthMonth = utils.PtrInt(5)
				in.Body.BirthYear = utils.PtrInt(2017)
				in.Body.Interests = &[]string{"science", "math"}
				return in
			}(),
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On(
					"UpdateChildByID",
					mock.Anything,
					mock.Anything,
					mock.AnythingOfType("*models.UpdateChildInput"),
				).Return(&models.Child{
					Name:       "Updated Emily",
					SchoolID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
					BirthMonth: 5,
					BirthYear:  2017,
					Interests:  []string{"science", "math"},
					GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
					CreatedAt:  time.Now(),
					UpdatedAt:  time.Now(),
				}, nil)
			},

			wantErr: false,
		},
		{
			name: "update non-existent child",
			input: func() *models.UpdateChildInput {
				in := &models.UpdateChildInput{}
				return in
			}(),
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On(
					"UpdateChildByID",
					mock.Anything, // context.Context
					mock.Anything, // uuid.UUID
					mock.AnythingOfType("*models.UpdateChildInput"),
				).Return(nil, &errs.HTTPError{
					Code:    errs.NotFound("Child", "id", "00000000-0000-0000-0000-000000000000").Code,
					Message: "Not found",
				})
			},

			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			child, err := handler.UpdateChildByID(ctx, tt.input.ID, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, child)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, child)
				assert.Equal(t, *tt.input.Body.Name, child.Name)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_DeleteChildByID(t *testing.T) {
	tests := []struct {
		name      string
		childID   string
		mockSetup func(*repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name:    "successful delete child",
			childID: "30000000-0000-0000-0000-000000000001",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("DeleteChildByID", mock.Anything, uuid.MustParse("30000000-0000-0000-0000-000000000001")).
					Return(&models.Child{
						ID:         uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						Name:       "Emily Johnson",
						SchoolID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						BirthMonth: 3,
						BirthYear:  2016,
						Interests:  []string{"science", "technology", "math"},
						GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name:    "delete non-existent child",
			childID: "00000000-0000-0000-0000-000000000000",
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("DeleteChildByID", mock.Anything, uuid.MustParse("00000000-0000-0000-0000-000000000000")).
					Return(nil, &errs.HTTPError{
						Code:    errs.NotFound("Child", "id", "00000000-0000-0000-0000-000000000000").Code,
						Message: "Not found",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()
			input := &models.ChildIDInput{ID: uuid.MustParse(tt.childID)}

			child, err := handler.DeleteChildByID(ctx, input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, child)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, child)
				assert.Equal(t, tt.childID, child.ID.String())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestHandler_CreateChild(t *testing.T) {
	tests := []struct {
		name      string
		input     *models.CreateChildInput
		mockSetup func(*repomocks.MockChildRepository)
		wantErr   bool
	}{
		{
			name: "successful create child",
			input: func() *models.CreateChildInput {
				in := &models.CreateChildInput{}
				in.Body.Name = "Emily Johnson"
				in.Body.SchoolID = uuid.MustParse("20000000-0000-0000-0000-000000000001")
				in.Body.BirthMonth = 3
				in.Body.BirthYear = 2016
				in.Body.Interests = []string{"science", "technology", "math"}
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				return in
			}(),
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("CreateChild", mock.Anything, mock.AnythingOfType("*models.CreateChildInput")).
					Return(&models.Child{
						ID:         uuid.MustParse("30000000-0000-0000-0000-000000000001"),
						Name:       "Emily Johnson",
						SchoolID:   uuid.MustParse("20000000-0000-0000-0000-000000000001"),
						BirthMonth: 3,
						BirthYear:  2016,
						Interests:  []string{"science", "technology", "math"},
						GuardianID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
						CreatedAt:  time.Now(),
						UpdatedAt:  time.Now(),
					}, nil)
			},
			wantErr: false,
		},
		{
			name: "create child with missing school - error",
			input: func() *models.CreateChildInput {
				in := &models.CreateChildInput{}
				in.Body.Name = "John Doe"
				// missing school ID
				in.Body.BirthMonth = 5
				in.Body.BirthYear = 2017
				in.Body.Interests = []string{"math"}
				in.Body.GuardianID = uuid.MustParse("11111111-1111-1111-1111-111111111111")
				return in
			}(),
			mockSetup: func(m *repomocks.MockChildRepository) {
				m.On("CreateChild", mock.Anything, mock.AnythingOfType("*models.CreateChildInput")).
					Return(nil, &errs.HTTPError{
						Code:    errs.BadRequest("Missing school ID").Code,
						Message: "Missing school ID",
					})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			mockRepo := new(repomocks.MockChildRepository)
			tt.mockSetup(mockRepo)

			handler := NewHandler(mockRepo)
			ctx := context.Background()

			child, err := handler.CreateChild(ctx, tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, child)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, child)
				assert.Equal(t, tt.input.Body.Name, child.Name)
				assert.Equal(t, tt.input.Body.SchoolID, child.SchoolID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
