# Huma API Documentation Guide

## Table of Contents

- [Introduction](#introduction)
- [Project Structure](#project-structure)
- [Quick Start](#quick-start)
- [Defining Models](#defining-models)
- [Route Registration](#route-registration)
- [Validation Tags](#validation-tags)
- [OpenAPI Metadata](#openapi-metadata)
- [Error Responses](#error-responses)
- [Examples](#examples)

---

## Introduction

Huma is a modern REST API framework for Go that automatically generates OpenAPI 3.1 documentation from your code. Key features:

- **Type-safe endpoints** - Define request/response types, get validation and docs automatically
- **Automatic OpenAPI generation** - No manual YAML writing required
- **Struct tag validation** - Use Go struct tags for input validation
- **Beautiful documentation** - Scalar UI served at `/docs`

---

## Project Structure

```
internal/
├── service/
│   ├── app.go                    # Main app setup
│   ├── routes/
│   │   ├── locations.go          # Location route registration
│   │   ├── users.go              # User route registration
│   │   └── products.go           # Product route registration
│   └── handler/
│       ├── location/handler.go   # Location business logic
│       ├── user/handler.go       # User business logic
│       └── product/handler.go    # Product business logic
└── models/
    ├── location.go               # Location request/response types
    ├── user.go                   # User request/response types
    └── product.go                # Product request/response types
```

---

## Quick Start

### 1. Define Your Models

**`internal/models/product.go`:**

```go
package models

// Input types
type CreateProductInput struct {
    Body struct {
        Name  string  `json:"name" minLength:"3" maxLength:"100" doc:"Product name"`
        Price float64 `json:"price" minimum:"0.01" doc:"Product price in USD"`
        Stock int     `json:"stock" minimum:"0" doc:"Available stock"`
    }
}

type GetProductInput struct {
    ID string `path:"id" doc:"Product ID"`
}

type ListProductsInput struct {
    Page  int `query:"page" default:"1" minimum:"1" doc:"Page number"`
    Limit int `query:"limit" default:"20" minimum:"1" maximum:"100" doc:"Items per page"`
}

// Output types
type ProductOutput struct {
    Body struct {
        ID        string  `json:"id" doc:"Product ID"`
        Name      string  `json:"name" doc:"Product name"`
        Price     float64 `json:"price" doc:"Product price"`
        Stock     int     `json:"stock" doc:"Available stock"`
        CreatedAt string  `json:"created_at" format:"date-time" doc:"Creation timestamp"`
    }
}

type ListProductsOutput struct {
    Body struct {
        Products []ProductOutput `json:"products" doc:"List of products"`
        Total    int             `json:"total" doc:"Total count"`
    }
}
```

### 2. Create Route Setup

**`internal/service/routes/products.go`:**

```go
package routes

import (
    "net/http"
    "skillspark/internal/models"
    "skillspark/internal/service/handler/product"
    "skillspark/internal/storage"
    
    "github.com/danielgtaylor/huma/v2"
)

func SetupProductRoutes(api huma.API, repo *storage.Repository) {
    h := product.NewHandler(repo.Product)
    
    // POST /api/v1/products
    huma.Register(api, huma.Operation{
        OperationID: "create-product",
        Method:      http.MethodPost,
        Path:        "/api/v1/products",
        Summary:     "Create a product",
        Description: "Creates a new product in the catalog",
        Tags:        []string{"Products"},
    }, h.Create)
    
    // GET /api/v1/products/{id}
    huma.Register(api, huma.Operation{
        OperationID: "get-product",
        Method:      http.MethodGet,
        Path:        "/api/v1/products/{id}",
        Summary:     "Get a product",
        Description: "Retrieves a single product by ID",
        Tags:        []string{"Products"},
    }, h.GetByID)
    
    // GET /api/v1/products
    huma.Register(api, huma.Operation{
        OperationID: "list-products",
        Method:      http.MethodGet,
        Path:        "/api/v1/products",
        Summary:     "List products",
        Description: "Returns a paginated list of products",
        Tags:        []string{"Products"},
    }, h.List)
}
```

### 3. Register Routes in App

**`internal/service/app.go`:**

```go
func setupHumaRoutes(api huma.API, repo *storage.Repository) {
    routes.SetupProductRoutes(api, repo)
    routes.SetupUserRoutes(api, repo)
    routes.SetupLocationRoutes(api, repo)
}
```

### 4. View Documentation

```bash
make dev          # Start the server
make api-gen      # Generate OpenAPI spec
make api-preview  # Open Scalar docs at http://localhost:8080/docs
```

---

## Defining Models

All request/response types go in `internal/models/` organized by resource.

### Model File Structure

**`internal/models/[resource].go`:**

```go
package models

// Input types for each endpoint
type Create[Resource]Input struct { ... }
type Get[Resource]Input struct { ... }
type List[Resource]Input struct { ... }
type Update[Resource]Input struct { ... }
type Delete[Resource]Input struct { ... }

// Output types
type [Resource]Output struct { ... }
type List[Resource]Output struct { ... }
```

### Input Type Anatomy

```go
type MyInput struct {
    // Path parameters - extracted from URL path
    ID string `path:"id" doc:"Resource ID"`
    
    // Query parameters - extracted from URL query string
    Page  int    `query:"page" default:"1" doc:"Page number"`
    Search string `query:"search,omitempty" doc:"Search query"`
    
    // Headers - extracted from HTTP headers
    Authorization string `header:"Authorization" doc:"Bearer token"`
    
    // Request body - JSON payload
    Body struct {
        Name  string `json:"name" minLength:"3" doc:"Resource name"`
        Email string `json:"email" format:"email" doc:"Email address"`
    }
}
```

### Output Type Anatomy

```go
type MyOutput struct {
    // HTTP status code (optional, defaults by method)
    Status int
    
    // Response body - becomes JSON
    Body struct {
        ID        string `json:"id" doc:"Resource ID"`
        Name      string `json:"name" doc:"Resource name"`
        CreatedAt string `json:"created_at" format:"date-time" doc:"Timestamp"`
    }
}
```

### Empty Input/Output

```go
// No input parameters (like health checks or list all)
func(ctx context.Context, input *struct{}) (*MyOutput, error)

// No output (like DELETE endpoints)
func(ctx context.Context, input *MyInput) (*struct{}, error)
```

---

## Route Registration

### Basic Registration

```go
huma.Register(api, huma.Operation{
    OperationID: "unique-operation-id",  // Unique identifier
    Method:      http.MethodGet,         // HTTP method
    Path:        "/api/v1/resource",     // URL path
    Summary:     "Short description",    // Brief summary
    Description: "Longer description",   // Detailed explanation
    Tags:        []string{"TagName"},    // Groups in docs
}, handlerFunction)
```

### Operation Metadata

```go
huma.Operation{
    OperationID: "get-user-profile",
    Method:      http.MethodGet,
    Path:        "/api/v1/users/{id}/profile",
    Summary:     "Get user profile",
    Description: "Retrieves detailed profile information for a specific user",
    Tags:        []string{"Users", "Profiles"},
    
    // Mark as deprecated
    Deprecated: true,
    
    // Link to external docs
    ExternalDocs: &huma.ExternalDocs{
        URL:         "https://docs.example.com/users",
        Description: "User API Documentation",
    },
}
```

### Default Response Codes

Huma automatically sets response codes based on HTTP method:

- `GET` → 200 OK
- `POST` → 201 Created
- `PUT`/`PATCH` → 200 OK
- `DELETE` → 204 No Content

### Custom Response Codes

```go
type MyOutput struct {
    Status int  // Set custom status code
    Body struct {
        Message string `json:"message"`
    }
}

func handler(ctx context.Context, input *MyInput) (*MyOutput, error) {
    resp := &MyOutput{}
    resp.Status = http.StatusAccepted  // 202
    resp.Body.Message = "Accepted"
    return resp, nil
}
```

---

## Validation Tags

Huma uses struct tags for automatic validation. Invalid requests return `400 Bad Request`.

### String Validation

```go
type MyInput struct {
    Body struct {
        // Length constraints
        Username string `json:"username" minLength:"3" maxLength:"30" doc:"Username"`
        
        // Format validation
        Email   string `json:"email" format:"email" doc:"Email address"`
        Website string `json:"website" format:"uri" doc:"Website URL"`
        Date    string `json:"date" format:"date" doc:"Date (YYYY-MM-DD)"`
        Time    string `json:"time" format:"date-time" doc:"ISO 8601 datetime"`
        
        // Pattern matching (regex)
        Code string `json:"code" pattern:"^[A-Z0-9]{6}$" doc:"6-char code"`
        
        // Enum (allowed values)
        Status string `json:"status" enum:"active,inactive,pending" doc:"Status"`
    }
}
```

**Available String Formats:**

- `format:"email"` - Email address
- `format:"uri"` - URL/URI
- `format:"date"` - Date (YYYY-MM-DD)
- `format:"date-time"` - ISO 8601 datetime
- `format:"uuid"` - UUID format

### Number Validation

```go
type MyInput struct {
    Body struct {
        // Range constraints
        Age      int     `json:"age" minimum:"18" maximum:"120" doc:"Age"`
        Price    float64 `json:"price" minimum:"0.01" doc:"Price"`
        Discount float64 `json:"discount" minimum:"0" maximum:"100" doc:"Discount %"`
        
        // Exclusive bounds (> instead of >=)
        Rating float64 `json:"rating" minimum:"0" exclusiveMinimum:"true" doc:"Rating > 0"`
        
        // Multiple of
        Quantity int `json:"quantity" multipleOf:"5" doc:"Must be multiple of 5"`
    }
}
```

### Array Validation

```go
type MyInput struct {
    Body struct {
        // Array length
        Tags []string `json:"tags" minItems:"1" maxItems:"10" doc:"1-10 tags"`
        
        // Unique items
        IDs []string `json:"ids" uniqueItems:"true" doc:"Unique IDs"`
    }
}
```

### Optional Fields

```go
type MyInput struct {
    Body struct {
        // Required (default)
        Name string `json:"name" doc:"Required field"`
        
        // Optional with omitempty
        Bio string `json:"bio,omitempty" doc:"Optional field"`
        
        // Optional with pointer
        Age *int `json:"age,omitempty" doc:"Optional age"`
    }
}
```

### Default Values

```go
type MyInput struct {
    Page  int    `query:"page" default:"1" doc:"Page number"`
    Limit int    `query:"limit" default:"20" doc:"Items per page"`
    Sort  string `query:"sort" default:"created_at" doc:"Sort field"`
}
```

---

## OpenAPI Metadata

### Documentation Tags

Every field should have a `doc:` tag that appears in the OpenAPI spec:

```go
type MyInput struct {
    Body struct {
        Name  string `json:"name" doc:"User's full name"`
        Email string `json:"email" format:"email" doc:"Email for notifications"`
        Age   int    `json:"age" minimum:"18" doc:"Must be 18 or older"`
    }
}
```

### Example Values

Add `example:` tags to show example values in docs:

```go
type ProductOutput struct {
    Body struct {
        ID    string  `json:"id" example:"prod_123abc" doc:"Product ID"`
        Name  string  `json:"name" example:"Wireless Mouse" doc:"Product name"`
        Price float64 `json:"price" example:"29.99" doc:"Price in USD"`
    }
}
```

### Grouping with Tags

Use `Tags` to group related endpoints in documentation:

```go
// All user endpoints
huma.Operation{
    Tags: []string{"Users"},
    // ...
}

// Admin-only endpoints
huma.Operation{
    Tags: []string{"Admin", "Users"},
    // ...
}

// Public endpoints
huma.Operation{
    Tags: []string{"Public"},
    // ...
}
```

### Deprecation

Mark endpoints as deprecated:

```go
huma.Operation{
    OperationID: "legacy-endpoint",
    Method:      http.MethodGet,
    Path:        "/api/v1/legacy",
    Summary:     "Legacy endpoint",
    Deprecated:  true,  // Shows as deprecated in docs
    Tags:        []string{"Deprecated"},
}
```

---

## Error Responses

Return Huma errors to generate proper OpenAPI error responses:

```go
import "github.com/danielgtaylor/huma/v2"

func handler(ctx context.Context, input *MyInput) (*MyOutput, error) {
    // Return Huma errors for automatic OpenAPI documentation
    return nil, huma.Error404NotFound("User not found")
}
```

### Common Error Functions

```go
huma.Error400BadRequest("Invalid input")
huma.Error401Unauthorized("Authentication required")
huma.Error403Forbidden("Access denied")
huma.Error404NotFound("Resource not found")
huma.Error409Conflict("Resource already exists")
huma.Error422UnprocessableEntity("Validation failed")
huma.Error500InternalServerError("Server error")
```

### Error with Details

```go
return nil, huma.Error400BadRequest("Validation failed",
    huma.ErrorDetail{
        Location: "body.email",
        Message:  "Email already exists",
        Value:    input.Body.Email,
    },
)
```

---

## Examples

### Example 1: CRUD Endpoints

**Models:**

```go
// internal/models/task.go
package models

type CreateTaskInput struct {
    Body struct {
        Title    string `json:"title" minLength:"3" maxLength:"200" doc:"Task title"`
        Priority string `json:"priority" enum:"low,medium,high" default:"medium" doc:"Priority"`
        DueDate  string `json:"due_date,omitempty" format:"date" doc:"Due date (YYYY-MM-DD)"`
    }
}

type GetTaskInput struct {
    ID string `path:"id" doc:"Task ID"`
}

type ListTasksInput struct {
    Status   string `query:"status,omitempty" enum:"todo,in_progress,done" doc:"Filter by status"`
    Priority string `query:"priority,omitempty" enum:"low,medium,high" doc:"Filter by priority"`
    Page     int    `query:"page" default:"1" minimum:"1" doc:"Page number"`
    Limit    int    `query:"limit" default:"20" minimum:"1" maximum:"100" doc:"Items per page"`
}

type UpdateTaskInput struct {
    ID   string `path:"id" doc:"Task ID"`
    Body struct {
        Title    *string `json:"title,omitempty" minLength:"3" maxLength:"200" doc:"Task title"`
        Priority *string `json:"priority,omitempty" enum:"low,medium,high" doc:"Priority"`
        Status   *string `json:"status,omitempty" enum:"todo,in_progress,done" doc:"Status"`
    }
}

type DeleteTaskInput struct {
    ID string `path:"id" doc:"Task ID"`
}

type TaskOutput struct {
    Body struct {
        ID        string `json:"id" doc:"Task ID"`
        Title     string `json:"title" doc:"Task title"`
        Priority  string `json:"priority" doc:"Task priority"`
        Status    string `json:"status" doc:"Task status"`
        DueDate   string `json:"due_date,omitempty" format:"date" doc:"Due date"`
        CreatedAt string `json:"created_at" format:"date-time" doc:"Creation timestamp"`
    }
}

type ListTasksOutput struct {
    Body struct {
        Tasks []TaskOutput `json:"tasks" doc:"List of tasks"`
        Total int          `json:"total" doc:"Total count"`
        Page  int          `json:"page" doc:"Current page"`
    }
}
```

**Routes:**

```go
// internal/service/routes/tasks.go
package routes

import (
    "net/http"
    "skillspark/internal/models"
    "skillspark/internal/service/handler/task"
    "skillspark/internal/storage"
    "github.com/danielgtaylor/huma/v2"
)

func SetupTaskRoutes(api huma.API, repo *storage.Repository) {
    h := task.NewHandler(repo.Task)
    
    huma.Register(api, huma.Operation{
        OperationID: "create-task",
        Method:      http.MethodPost,
        Path:        "/api/v1/tasks",
        Summary:     "Create a task",
        Tags:        []string{"Tasks"},
    }, h.Create)
    
    huma.Register(api, huma.Operation{
        OperationID: "get-task",
        Method:      http.MethodGet,
        Path:        "/api/v1/tasks/{id}",
        Summary:     "Get a task",
        Tags:        []string{"Tasks"},
    }, h.GetByID)
    
    huma.Register(api, huma.Operation{
        OperationID: "list-tasks",
        Method:      http.MethodGet,
        Path:        "/api/v1/tasks",
        Summary:     "List tasks",
        Tags:        []string{"Tasks"},
    }, h.List)
    
    huma.Register(api, huma.Operation{
        OperationID: "update-task",
        Method:      http.MethodPatch,
        Path:        "/api/v1/tasks/{id}",
        Summary:     "Update a task",
        Tags:        []string{"Tasks"},
    }, h.Update)
    
    huma.Register(api, huma.Operation{
        OperationID: "delete-task",
        Method:      http.MethodDelete,
        Path:        "/api/v1/tasks/{id}",
        Summary:     "Delete a task",
        Tags:        []string{"Tasks"},
    }, h.Delete)
}
```

### Example 2: Search with Filters

**Models:**

```go
// internal/models/product.go
type SearchProductsInput struct {
    Query      string   `query:"q" minLength:"2" maxLength:"100" doc:"Search query"`
    Category   string   `query:"category,omitempty" doc:"Filter by category"`
    MinPrice   *float64 `query:"min_price,omitempty" minimum:"0" doc:"Minimum price"`
    MaxPrice   *float64 `query:"max_price,omitempty" minimum:"0" doc:"Maximum price"`
    InStock    *bool    `query:"in_stock,omitempty" doc:"Filter in-stock items"`
    SortBy     string   `query:"sort_by" default:"relevance" enum:"relevance,price_asc,price_desc,newest" doc:"Sort order"`
    Page       int      `query:"page" default:"1" minimum:"1" doc:"Page number"`
    Limit      int      `query:"limit" default:"24" minimum:"1" maximum:"100" doc:"Items per page"`
}

type SearchProductsOutput struct {
    Body struct {
        Products []ProductOutput `json:"products" doc:"Matching products"`
        Total    int             `json:"total" doc:"Total results"`
        Page     int             `json:"page" doc:"Current page"`
    }
}
```

**Routes:**

```go
huma.Register(api, huma.Operation{
    OperationID: "search-products",
    Method:      http.MethodGet,
    Path:        "/api/v1/products/search",
    Summary:     "Search products",
    Description: "Full-text search with filtering and sorting",
    Tags:        []string{"Products", "Search"},
}, h.Search)
```

### Example 3: Nested Resources

**Models:**

```go
// internal/models/comment.go
type CreateCommentInput struct {
    PostID string `path:"post_id" doc:"Post ID"`
    Body struct {
        Content string `json:"content" minLength:"1" maxLength:"1000" doc:"Comment content"`
    }
}

type ListCommentsInput struct {
    PostID string `path:"post_id" doc:"Post ID"`
    Page   int    `query:"page" default:"1" doc:"Page number"`
    Limit  int    `query:"limit" default:"20" doc:"Items per page"`
}

type CommentOutput struct {
    Body struct {
        ID        string `json:"id" doc:"Comment ID"`
        PostID    string `json:"post_id" doc:"Post ID"`
        Content   string `json:"content" doc:"Comment content"`
        CreatedAt string `json:"created_at" format:"date-time" doc:"Creation time"`
    }
}
```

**Routes:**

```go
huma.Register(api, huma.Operation{
    OperationID: "create-post-comment",
    Method:      http.MethodPost,
    Path:        "/api/v1/posts/{post_id}/comments",
    Summary:     "Create a comment",
    Tags:        []string{"Posts", "Comments"},
}, h.CreateComment)

huma.Register(api, huma.Operation{
    OperationID: "list-post-comments",
    Method:      http.MethodGet,
    Path:        "/api/v1/posts/{post_id}/comments",
    Summary:     "List comments",
    Tags:        []string{"Posts", "Comments"},
}, h.ListComments)
```

---

## Quick Reference

### Validation Tags

| Tag | Type | Example | Description |
|-----|------|---------|-------------|
| `json:"name"` | All | `json:"email"` | JSON field name |
| `path:"name"` | String/Int | `path:"user_id"` | URL path parameter |
| `query:"name"` | All | `query:"page"` | URL query parameter |
| `header:"name"` | String | `header:"Authorization"` | HTTP header |
| `doc:"text"` | All | `doc:"User email"` | OpenAPI description |
| `example:"val"` | All | `example:"john@example.com"` | Example value |
| `default:"val"` | All | `default:"10"` | Default value |
| `minLength:"n"` | String | `minLength:"3"` | Min string length |
| `maxLength:"n"` | String | `maxLength:"100"` | Max string length |
| `pattern:"regex"` | String | `pattern:"^[a-z]+$"` | Regex pattern |
| `format:"type"` | String | `format:"email"` | String format |
| `enum:"a,b,c"` | String | `enum:"active,inactive"` | Allowed values |
| `minimum:"n"` | Number | `minimum:"0"` | Minimum value |
| `maximum:"n"` | Number | `maximum:"100"` | Maximum value |
| `minItems:"n"` | Array | `minItems:"1"` | Min array length |
| `maxItems:"n"` | Array | `maxItems:"10"` | Max array length |

### HTTP Methods

```go
http.MethodGet     // GET
http.MethodPost    // POST
http.MethodPut     // PUT
http.MethodPatch   // PATCH
http.MethodDelete  // DELETE
```

### Naming Conventions

```go
// ✅ Good - Consistent kebab-case
OperationID: "create-user"
OperationID: "get-user"
OperationID: "list-users"
OperationID: "update-user"
OperationID: "delete-user"

// ❌ Bad - Inconsistent
OperationID: "createUser"
OperationID: "user-get"
OperationID: "users"
```

---

## Checklist for New Endpoints

- [ ] Create input/output types in `internal/models/[resource].go`
- [ ] Add validation tags (`minLength`, `format`, `enum`, etc.)
- [ ] Add `doc:` tags to all fields
- [ ] Register endpoint in `internal/service/routes/[resource].go`
- [ ] Set unique `OperationID`
- [ ] Add clear `Summary` and `Description`
- [ ] Add appropriate `Tags`
- [ ] Test endpoint
- [ ] Run `make api-gen` to update OpenAPI spec
- [ ] Check docs at `http://localhost:8080/docs`

---

## Resources

- **Huma Documentation**: https://huma.rocks
- **OpenAPI 3.1 Spec**: https://spec.openapis.org/oas/v3.1.0
- **Scalar UI**: https://github.com/scalar/scalar

As always reach out to your TLs for questions <3
