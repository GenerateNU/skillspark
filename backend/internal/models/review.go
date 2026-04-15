package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID             uuid.UUID  `json:"id" db:"id" doc:"Unique review identifier"`
	RegistrationID uuid.UUID  `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
	GuardianID     *uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian. Null when the review was submitted anonymously."`
	EventID        uuid.UUID  `json:"event_id" db:"event_id" doc:"ID of the event"`
	Description    string     `json:"description" db:"description" doc:"The review text"`
	Categories     []string   `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	Rating         int        `json:"rating" db:"rating" doc:"Rating left with the review, can be 1-5 inclusive"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at" doc:"Timestamp when registration was created"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at" doc:"Timestamp when registration was last updated"`
}

type CreateReviewDBBody struct {
	RegistrationID uuid.UUID  `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
	GuardianID     *uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian. Omit or set to null for an anonymous review."`
	Description_EN string     `json:"description_en" db:"description_en" doc:"The review text"`
	Description_TH *string    `json:"description_th" db:"description_th" doc:"The review text"`
	Categories     []string   `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	Rating         int        `json:"rating" db:"rating" doc:"Rating left with the review, can be 1-5 inclusive"`
}

type CreateReviewDBInput struct {
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Body           CreateReviewDBBody
}

type GetReviewInput struct {
	ID             uuid.UUID  `json:"id" db:"id" doc:"Unique review identifier"`
	RegistrationID uuid.UUID  `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
	GuardianID     *uuid.UUID `json:"guardian_id" db:"guardian_id" doc:"ID of the guardian. Null for anonymous reviews."`
	EventID        uuid.UUID  `json:"event_id" db:"event_id" doc:"ID of the event"`
	Rating         int        `json:"rating" db:"rating" doc:"Rating left with the review, can be 1-5 inclusive"`
	Description_EN string     `json:"description_en" db:"description" doc:"The review text"`
	Description_TH *string    `json:"description_th" db:"description" doc:"The review text"`
	Categories     []string   `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
	CreatedAt      time.Time  `json:"created_at" db:"created_at" doc:"Timestamp when registration was created"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at" doc:"Timestamp when registration was last updated"`
}

type CreateReviewInput struct {
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Body           struct {
		RegistrationID uuid.UUID  `json:"registration_id" db:"registration_id" doc:"ID of the linked registration"`
		GuardianID     *uuid.UUID `json:"guardian_id,omitempty" db:"guardian_id" doc:"ID of the guardian. Omit or set to null for an anonymous review."`
		Description    string     `json:"description" db:"description" doc:"The review text"`
		Categories     []string   `json:"categories" db:"categories" doc:"Review categories for this review, can be one of fun, engaging, interesting or informative."`
		Rating         int        `json:"rating" db:"rating" doc:"Rating left with the review, can be 1-5 inclusive"`
	}
}

type GetReviewsByGuardianIDInput struct {
	ID             uuid.UUID `path:"id"`
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Page           int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize       int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetReviewsByEventIDInput struct {
	ID             uuid.UUID `path:"id"`
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Page           int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize       int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
	SortBy         string    `query:"sort_by" default:"most_recent" enum:"most_recent,highest,lowest" doc:"Sort order for reviews"`
}
type GetReviewsByOrganizationIDInput struct {
	ID             uuid.UUID `path:"id"`
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	Page           int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize       int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
}

type GetReviewsAggregateInput struct {
	ID uuid.UUID `path:"id"`
}

type ReviewsAggregateOutput struct {
	Body ReviewAggregate `json:"body" doc:"Aggregate review data including breakdown by rating, total reviews, and average rating"`
}

type DeleteReviewInput struct {
	ID uuid.UUID `path:"id"`
}

type ReviewsOutput struct {
	Body []Review `json:"body" doc:"List of reviews"`
}

type DeleteReviewOutput struct {
	Body struct {
		Message string `json:"message" doc:"Success message"`
	} `json:"body"`
}

type CreateReviewOutput struct {
	Body Review
}

type ReviewRatingCount struct {
	Rating      int `json:"rating" db:"rating"`
	ReviewCount int `json:"review_count" db:"review_count"`
}

type ReviewAggregate struct {
	EventID       uuid.UUID           `json:"event_id"`
	TotalReviews  int                 `json:"total_reviews"`
	AverageRating float64             `json:"average_rating"`
	Breakdown     []ReviewRatingCount `json:"breakdown"`
}

type SimpleReviewAggregate struct {
	EventID       uuid.UUID `json:"event_id"`
	TotalReviews  int       `json:"total_reviews"`
	AverageRating float64   `json:"average_rating"`
	Event         Event     `json:"event" db:"-"`
}

type GetEventReviewsForOrganizationInput struct {
	ID       uuid.UUID `path:"id"`
	Page     int       `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	PageSize int       `query:"page_size" minimum:"1" maximum:"100" default:"10" doc:"Number of items per page"`
	SortBy   string    `query:"sort_by" default:"most_rated" enum:"most_rated,highest,lowest"`
}

type GetEventReviewsForOrganizationOutput struct {
	Body []SimpleReviewAggregate `json:"body" doc:"List of review aggregates"`
}

type OrgReviewSummary struct {
	TotalReviews  int                 `json:"total_reviews"`
	AverageRating float64             `json:"average_rating"`
	Breakdown     []ReviewRatingCount `json:"breakdown"`
}
