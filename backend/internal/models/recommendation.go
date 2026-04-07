package models

import (
	"time"

	"github.com/google/uuid"
)

type GetRecommendationsByChildIDInput struct {
	AcceptLanguage string          `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
	ChildID        uuid.UUID       `path:"child_id"`
	Page           int             `query:"page" minimum:"1" default:"1"`
	Limit          int             `query:"limit" minimum:"1" maximum:"100" default:"10"`
	Latitude       OptionalFloat64 `query:"lat"`
	Longitude      OptionalFloat64 `query:"lng"`
	RadiusKm       float64         `query:"radius_km"`
	MinDate        time.Time       `query:"min_date"`
	MaxDate        time.Time       `query:"max_date"`
}

type GetRecommendationsByChildIDOutput struct {
	Body []Event `json:"body" doc:"List of recommended events for the child"`
}

type RecommendationFilters struct {
	Page      int             `query:"page" minimum:"1" default:"1"`
	Limit     int             `query:"limit" minimum:"1" maximum:"100" default:"10"`
	Latitude  OptionalFloat64 `query:"lat"`
	Longitude OptionalFloat64 `query:"lng"`
	RadiusKm  float64         `query:"radius_km"`
	MinDate   time.Time       `query:"min_date"`
	MaxDate   time.Time       `query:"max_date"`
}
