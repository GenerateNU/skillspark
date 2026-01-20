# Backend Development Guide

## Overview

This guide explains how the SkillSpark backend is organized and how to create new endpoints. The architecture is organized by domain objects (sessions, users, skills, etc.) with clear separation of concerns across layers.

## Architecture Layers

```
API Request
    ↓
Handler (Business Logic & Validation)
    ↓
Repository Interface (Contract)
    ↓
Repository Implementation (Database Logic)
    ↓
Database
```

## Directory Structure

```
/internal
├── /service
│   ├── server.go              # Route configuration
│   └── /handler
│       └── /[object]          # One folder per domain object
│           ├── handler.go     # Handler struct and constructor
│           ├── get_all.go     # GET /objects
│           ├── get_by_id.go   # GET /objects/:id
│           ├── create.go      # POST /objects
│           ├── update.go      # PATCH /objects/:id
│           └── delete.go      # DELETE /objects/:id
├── /storage
│   ├── storage.go             # Repository interfaces
│   └── /postgres
│       ├── storage.go         # Database connection (rarely edited)
│       └── /schema
│           └── [object].go    # Repository implementation
└── /models
    └── [object].go            # Data models, inputs, outputs
```

---

## Creating a New Endpoint: Step-by-Step

### Step 1: Define Models (`/internal/models/[object].go`)

Create your database model and input/output types:

```go
// Database model - matches DB schema
type Session struct {
    ID        uuid.UUID `json:"id" db:"id"`
    Title     string    `json:"title" db:"title"`
    // ... other fields
}

// Input for creating
type CreateSessionInput struct {
    Title string `json:"title" validate:"required,min=1,max=200"`
    Date  string `json:"date" validate:"required,datetime=2006-01-02"`
}

// Input for updating (use pointers for optional fields)
type UpdateSessionInput struct {
    Title *string `json:"title,omitempty" validate:"omitempty,min=1,max=200"`
    Date  *string `json:"date,omitempty" validate:"omitempty,datetime=2006-01-02"`
}

// Filter for queries
type SessionFilter struct {
    UserID *uuid.UUID `json:"user_id,omitempty"`
    Status *string    `json:"status,omitempty"`
}
```

**Key Points:**

- Use `json` tags for API, `db` tags for database, `validate` tags for validation
- See [xvalidator.md](./xvalidator.md) for validation details
- Use pointers in update types for optional fields

---

### Step 2: Define Repository Interface (`/internal/storage/storage.go`)

Add your interface and register it in the Repository struct:

```go
type SessionRepository interface {
    CreateSession(ctx context.Context, input *models.CreateSessionInput) (*models.Session, error)
    GetSessions(ctx context.Context, pagination utils.Pagination, filter *models.SessionFilter) ([]models.Session, error)
    GetSessionByID(ctx context.Context, id uuid.UUID) (*models.Session, error)
    PatchSession(ctx context.Context, id uuid.UUID, input *models.UpdateSessionInput) (*models.Session, error)
    DeleteSession(ctx context.Context, id uuid.UUID) error
}

type Repository struct {
    // ... existing
    Session SessionRepository
    // ... more repos
}

func NewRepository(db *pgxpool.Pool) *Repository {
    return &Repository{
        // ...
        Session: schema.NewSessionRepository(db),
    }
}
```

---

### Step 3: Implement Repository (`/internal/storage/postgres/schema/[object].go`)

Implement all database operations:

```go
type SessionRepository struct {
    db *pgxpool.Pool
}

func (r *SessionRepository) CreateSession(ctx context.Context, input *models.CreateSessionInput) (*models.Session, error) {
    // INSERT query with RETURNING
    // Use errs.InternalServerError for errors
}

func (r *SessionRepository) GetSessions(ctx context.Context, pagination utils.Pagination, filter *models.SessionFilter) ([]models.Session, error) {
    // Build dynamic WHERE clause from filter
    // Add LIMIT/OFFSET from pagination
    // Use pgx.CollectRows for multiple results
}

func (r *SessionRepository) GetSessionByID(ctx context.Context, id uuid.UUID) (*models.Session, error) {
    // SELECT WHERE id = $1
    // Use pgx.CollectExactlyOneRow
    // Return errs.NotFound if not found
}

func (r *SessionRepository) PatchSession(ctx context.Context, id uuid.UUID, input *models.UpdateSessionInput) (*models.Session, error) {
    // Build dynamic UPDATE SET clause
    // Check if any fields provided
    // Return updated record with RETURNING
}

func (r *SessionRepository) DeleteSession(ctx context.Context, id uuid.UUID) error {
    // DELETE WHERE id = $1
}

func NewSessionRepository(db *pgxpool.Pool) *SessionRepository {
    return &SessionRepository{db}
}
```

**Key Points:**

- Use `pgx.CollectRows` for multiple results
- Use `pgx.CollectExactlyOneRow` for single results
- Always use `errs` package for errors
- Build dynamic queries for filters/updates
- Always `defer rows.Close()` after queries

---

### Step 4: Create Handler (`/internal/service/handler/[object]/`)

#### `handler.go` - Create handler struct

```go
type Handler struct {
    sessionRepository storage.SessionRepository
    validator         *xvalidator.XValidator
}

func NewHandler(sessionRepository storage.SessionRepository) *Handler {
    return &Handler{
        sessionRepository: sessionRepository,
        validator:         xvalidator.Validator,
    }
}
```

#### Create separate files for each method

**`create.go`:**

1. Parse JSON body with `c.BodyParser()`
2. Validate with `h.validator.Validate()` - returns `errs.InvalidRequestData()` if errors
3. Add custom validation for complex business rules (e.g., date comparisons)
4. Call repository method
5. Handle specific database errors
6. Return `fiber.StatusCreated` with created object

**`get_all.go`:**

1. Parse pagination with `utils.ParsePagination(c)`
2. Parse filters from query params (`c.Query()`, `c.QueryInt()`)
3. Call repository method
4. Return `fiber.StatusOK` with results

**`get_by_id.go`:**

1. Parse UUID from `c.Params("id")`
2. Validate UUID format
3. Call repository method
4. Return `fiber.StatusOK` with result

**`update.go`:**

1. Parse UUID from params
2. Parse JSON body
3. Validate with xvalidator
4. Add custom validation
5. Call repository method
6. Return `fiber.StatusOK` with updated object

**`delete.go`:**

1. Parse UUID from params
2. Call repository method
3. Return `fiber.StatusNoContent`

---

### Step 5: Register Routes (`/internal/service/server.go`)

Add your handler and routes in `SetupApp`:

```go
sessionHandler := session.NewHandler(repo.Session)
apiV1.Route("/sessions", func(r fiber.Router) {
    r.Post("/", sessionHandler.CreateSession)
    r.Get("/", sessionHandler.GetSessions)
    r.Get("/:id", sessionHandler.GetSessionByID)
    r.Patch("/:id", sessionHandler.PatchSession)
    r.Delete("/:id", sessionHandler.DeleteSession)
})
```

---

## Validation Strategy

### Use XValidator For

- Required fields
- String length, numeric ranges
- Email/URL/UUID formats
- Enum values (`oneof`)

See [xvalidator.md](./xvalidator.md) for complete validation guide.

### Use Custom Validation For

- **Cross-field validation:** "Start date must be before end date"
- **Business rules:** "Cannot schedule on weekends"
- **Database checks:** "User must exist"
- **Complex logic:** "Age 18+ OR parental consent"

Add custom validation in handlers after xvalidator but before repository call.

---

## Common Patterns

**Parse URL params:**

```go
id, err := uuid.Parse(c.Params("id"))
if err != nil {
    return errs.BadRequest("Invalid ID format")
}
```

**Parse query params:**

```go
search := c.Query("search")
page := c.QueryInt("page", 1) // with default
```

**Handle pagination:**

```go
pagination := utils.ParsePagination(c)
```

**Error handling:**

```go
if err != nil {
    slog.Error("Operation failed", "error", err)
    return err // Repository already returns proper error type
}
```

---

## Checklist for New Endpoints

- [ ] **Models** (`/internal/models/[object].go`)
  - Database model, Create/Update inputs, Filter struct
  
- [ ] **Repository Interface** (`/internal/storage/storage.go`)
  - Define interface, add to Repository struct, initialize in NewRepository

- [ ] **Repository Implementation** (`/internal/storage/postgres/schema/[object].go`)
  - Implement all methods with proper error handling

- [ ] **Handler** (`/internal/service/handler/[object]/`)
  - `handler.go` + separate files for each operation
  - Parse inputs, validate, call repo, return response

- [ ] **Routes** (`/internal/service/server.go`)
  - Create handler instance and register routes

- [ ] **Test** with curl or API client

---

## Running the App Locally

### Step 1: Start Local Supabase

From the project root, start the local Supabase instance:

```bash
cd backend/internal
bun run supabase start
```

Once started, the CLI will output connection details including URLs and keys. Keep this output handy for the next step.

---

### Step 2: Configure Environment

Create `backend/.local.env` using `backend/.env.template` as a reference:

```
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=host.docker.internal
DB_PORT=54322
DB_NAME=postgres
PORT=8080
SUPABASE_URL=http://host.docker.internal:54321
SUPABASE_ANON_KEY=sb_publishable_ACJWlzQHlZjBrEguHvfOxg_3BJgxAaH
SUPABASE_SERVICE_ROLE_KEY=sb_secret_N7UND0UgjKTVK-Uodkm0Hg_xSvEMPvz
DB_SSLMODE=disable
```

Replace the placeholder values with the keys from your `supabase start` output.

> **Important:** The Supabase output may show `127.0.0.1` for URLs and hosts. Replace any instance of `127.0.0.1` with `host.docker.internal` — otherwise Docker containers will reference their own localhost rather than your machine.

---

### Step 3: Set Environment Mode

In the root `.env` file, ensure the `ENVIRONMENT` variable is set:

```
ENVIRONMENT=development
```

This tells the backend to load `backend/.local.env`. Setting it to `production` would load `backend/.env` instead.

---

### Step 4: Start the Application

From the project root:

```bash
make up
```

This spins up both the backend and frontend.

---

### Step 5: Verify It's Working

1. Open Supabase Studio at `http://127.0.0.1:54323` to view your local database
2. Create a test entity via the API
3. Confirm the data appears in Supabase Studio

---

### Stopping the Application

To stop the Docker containers:

```bash
make down
```

To stop the local Supabase instance:

```bash
cd backend/internal
bun run supabase stop
```

--- 

## Summary

This architecture provides clear separation between API handling, business logic, and data access. Each domain object is self-contained, making the codebase easy to navigate and maintain. Follow this pattern for all new endpoints to maintain consistency.
