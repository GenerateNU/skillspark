package routes

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

// Health check types
type HealthOutput struct {
	Body struct {
		Status  string `json:"status" doc:"Health status" example:"ok"`
		Version string `json:"version" doc:"API version" example:"1.0.0"`
	}
}

func SetupBaseRoutes(api huma.API) {
	// Health check endpoint
	huma.Register(api, huma.Operation{
		OperationID: "health-check",
		Method:      http.MethodGet,
		Path:        "/api/v1/health",
		Summary:     "Health check",
		Description: "Check if the API is running and healthy",
		Tags:        []string{"Health"},
	}, func(ctx context.Context, input *struct{}) (*HealthOutput, error) {
		resp := &HealthOutput{}
		resp.Body.Status = "ok"
		resp.Body.Version = "1.0.0"
		return resp, nil
	})
}
