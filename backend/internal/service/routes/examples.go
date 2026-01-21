package routes

import (
	"context"
	"net/http"
	"skillspark/internal/storage"
	"time"

	"github.com/danielgtaylor/huma/v2"
)

// Example request/response types for testing
type GetGreetingInput struct {
	Name string `path:"name" maxLength:"30" doc:"Name to greet"`
}

type GetGreetingOutput struct {
	Body struct {
		Message   string `json:"message" doc:"Greeting message"`
		Timestamp string `json:"timestamp" doc:"Server timestamp"`
	}
}

type CreateUserInput struct {
	Body struct {
		Email    string `json:"email" format:"email" doc:"User email address"`
		Username string `json:"username" minLength:"3" maxLength:"50" doc:"Username"`
		FullName string `json:"full_name,omitempty" doc:"Full name (optional)"`
	}
}

type UserOutput struct {
	Body struct {
		ID       string `json:"id" doc:"User ID"`
		Email    string `json:"email" doc:"User email"`
		Username string `json:"username" doc:"Username"`
		FullName string `json:"full_name,omitempty" doc:"Full name"`
	}
}

type ListUsersOutput struct {
	Body struct {
		Users []struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		} `json:"users"`
		Total int `json:"total"`
	}
}

type ErrorOutput struct {
	Body struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}
}

func SetupExamplesRoutes(api huma.API, repo *storage.Repository) {
	// Example 1: Simple greeting endpoint with path parameter
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/api/v1/greet/{name}",
		Summary:     "Get a greeting",
		Description: "Returns a personalized greeting message",
		Tags:        []string{"Examples"},
	}, func(ctx context.Context, input *GetGreetingInput) (*GetGreetingOutput, error) {
		resp := &GetGreetingOutput{}
		resp.Body.Message = "Hello, " + input.Name + "!"
		resp.Body.Timestamp = time.Now().Format(time.RFC3339)
		return resp, nil
	})

	// Example 2: Create user endpoint with request body validation
	huma.Register(api, huma.Operation{
		OperationID: "create-user",
		Method:      http.MethodPost,
		Path:        "/api/v1/users",
		Summary:     "Create a new user",
		Description: "Creates a new user account with the provided information",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *CreateUserInput) (*UserOutput, error) {
		// Simulate user creation
		resp := &UserOutput{}
		resp.Body.ID = "user_123456"
		resp.Body.Email = input.Body.Email
		resp.Body.Username = input.Body.Username
		resp.Body.FullName = input.Body.FullName
		return resp, nil
	})

	// Example 3: List users endpoint
	huma.Register(api, huma.Operation{
		OperationID: "list-users",
		Method:      http.MethodGet,
		Path:        "/api/v1/users",
		Summary:     "List all users",
		Description: "Returns a list of all registered users",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct{}) (*ListUsersOutput, error) {
		resp := &ListUsersOutput{}
		resp.Body.Users = []struct {
			ID       string `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			{ID: "user_1", Email: "alice@example.com", Username: "alice"},
			{ID: "user_2", Email: "bob@example.com", Username: "bob"},
		}
		resp.Body.Total = len(resp.Body.Users)
		return resp, nil
	})

	// Example 4: Get user by ID with path parameter
	huma.Register(api, huma.Operation{
		OperationID: "get-user",
		Method:      http.MethodGet,
		Path:        "/api/v1/users/{id}",
		Summary:     "Get user by ID",
		Description: "Returns a single user by their ID",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"User ID"`
	}) (*UserOutput, error) {
		resp := &UserOutput{}
		resp.Body.ID = input.ID
		resp.Body.Email = "user@example.com"
		resp.Body.Username = "testuser"
		resp.Body.FullName = "Test User"
		return resp, nil
	})

	// Example 5: Delete user endpoint
	huma.Register(api, huma.Operation{
		OperationID: "delete-user",
		Method:      http.MethodDelete,
		Path:        "/api/v1/users/{id}",
		Summary:     "Delete a user",
		Description: "Deletes a user account by ID",
		Tags:        []string{"Users"},
	}, func(ctx context.Context, input *struct {
		ID string `path:"id" doc:"User ID to delete"`
	}) (*struct{}, error) {
		// Simulate deletion (return empty response with 204 status)
		return &struct{}{}, nil
	})
}
