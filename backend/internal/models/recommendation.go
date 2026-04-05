package models

import "github.com/google/uuid"

type GetRecommendationsByChildIDInput struct {
	AcceptLanguage string    `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ChildID        uuid.UUID `path:"child_id"`
}

type GetRecommendationsByChildIDOutput struct {
	Body []EventOccurrence `json:"body" doc:"List of recommended event occurrences for the child"`
}