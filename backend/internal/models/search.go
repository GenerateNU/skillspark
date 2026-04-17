package models

type SearchEventsInput struct {
	Query          string `query:"q" doc:"Search query string" minLength:"1"`
	Page           int    `query:"page" minimum:"1" default:"1" doc:"Page number (starts at 1)"`
	Limit          int    `query:"limit" minimum:"1" maximum:"100" default:"10" doc:"Number of results per page"`
	AcceptLanguage string `header:"Accept-Language" default:"en-US" enum:"en-US,th-TH"`
}

type SearchEventsOutput struct {
	Body []Event
}
