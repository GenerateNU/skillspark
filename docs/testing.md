
# 🧪 Skillspark Go Testing Guide (Mocks + DB)

This guide explains how to test **handlers, routes, and database repositories** in Skillspark.

---

## 1. **Testing Types**

| Type                           | Purpose                                                                  | Examples                                |
| ------------------------------ | ------------------------------------------------------------------------ | --------------------------------------- |
| **Handler Tests (Unit)**       | Test business logic without HTTP or DB                                   | `TestHandler_CreateLocation`            |
| **Route/API Tests (Mock)**     | Test HTTP endpoints with Huma validation using **mocked DB**             | `TestHumaValidation_CreateLocation`     |
| **Database Integration Tests** | Test repository methods against a real Postgres test DB (TestContainers) | `TestLocationRepository_Create_NewYork` |

---

## 2. **Core Concepts**

* **Mocks**: Replace the database with `repomocks.MockLocationRepository` to control behavior.
* **Huma validation**: Handles schema validation (e.g., required fields, UUID format).
* **TestContainers**: Run **isolated Postgres databases** in Docker for integration tests.
* **t.Parallel()**: Allows running tests in parallel to speed up execution.
* **Template DBs**: Use a template database to speed up per-test DB setup instead of truncating tables.

---

## 3. **Mock-Based Route Testing (Fast, No DB)**

1. **Setup API**: Use `setupTestAPI(mockRepo)` to get a Fiber app + Huma API.
2. **Mock repository methods**:

```go
mockRepo := new(repomocks.MockLocationRepository)
mockRepo.On("CreateLocation", mock.Anything, mock.AnythingOfType("*models.CreateLocationInput")).
    Return(&models.Location{ID: uuid.New(), City: "NY"}, nil)
```

3. **Send HTTP request**:

```go
payload := map[string]interface{}{
    "latitude": 40.7128,
    "longitude": -74.0060,
    "city": "New York",
}
bodyBytes, _ := json.Marshal(payload)

req, _ := http.NewRequest(http.MethodPost, "/api/v1/locations", bytes.NewBuffer(bodyBytes))
req.Header.Set("Content-Type", "application/json")
resp, _ := app.Test(req)
```

4. **Assert results**:

```go
assert.Equal(t, http.StatusOK, resp.StatusCode)
mockRepo.AssertExpectations(t)
```

**Tip:** Huma returns **422** for schema validation errors (not 400). Ensure payload matches the struct.

---

## 4. **Database Repository Tests (Integration)**

* Uses **TestContainers** to run a fresh Postgres DB.
* Tests can safely run in **parallel** using `SetupTestDB(t)`.

```go
testDB := testutil.SetupTestDB(t)
repo := NewLocationRepository(testDB)
ctx := context.Background()

input := &models.CreateLocationInput{
    Body: models.CreateLocationBody{
        Latitude: 40.7128,
        Longitude: -74.0060,
        City: "New York",
        State: "NY",
        ZipCode: "10001",
        Country: "USA",
        Address: "123 Broadway",
    },
}

loc, err := repo.CreateLocation(ctx, input)
assert.Nil(t, err)
assert.Equal(t, "New York", loc.City)

// Retrieve to verify
retrieved, err := repo.GetLocationByID(ctx, loc.ID)
assert.Equal(t, loc.ID, retrieved.ID)
```

**Tips**:

* Skip DB tests in **short mode**: `if testing.Short() { t.Skip() }`
* Each test gets a **fresh DB cloned from a template**.
* Cleanup happens automatically using `t.Cleanup`.
* No need for truncation or reseeding — faster setup.
* Use `t.Parallel()` to run multiple DB tests concurrently.

---

## 5. **Writing Tests With Mocks vs DB**

| Feature      | Mock Test                | DB Test                               |
| ------------ | ------------------------ | ------------------------------------- |
| Speed        | Fast                     | Slower (but template DB speeds it up) |
| Isolation    | Fully controlled         | Real database                         |
| Coverage     | Handler + validation     | Repository + SQL                      |
| Dependencies | None                     | Docker required                       |
| Setup        | `setupTestAPI(mockRepo)` | `testutil.SetupTestDB(t)`             |

**Recommended Workflow**:

1. Use **mock tests** for route validation and handler logic.
2. Use **DB tests** for SQL correctness and repository integration.

---

## 6. **Best Practices**

1. **t.Parallel()** everywhere to speed up tests.
2. **Wrap range variables** in subtests:

```go
for _, tt := range tests {
    tt := tt // capture range variable
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        ...
    })
}
```

3. **Only mock external dependencies** in handler or route tests.
4. **Always assert mock expectations**: `mockRepo.AssertExpectations(t)`
5. **Use real DB tests for migrations, seeds, and queries**.
6. **Skip DB tests in short mode**: `go test -short`

---

## 7. **Running Tests (Makefile)**

The Makefile provides unified commands for running tests:

```bash
# Run all tests (mocks + DB)
make test

# Run unit tests only (fast, mocks)
make test-unit

# Run database integration tests
make test-db

# Run tests with coverage report
make test-coverage

# Run a specific test
make test-one TEST=TestHumaValidation_CreateLocation

# Clean test cache and coverage files
make test-clean

# Skip slow DB tests in short mode
go test -short ./...
```

**Coverage HTML**: `coverage.html` is generated for viewing in a browser.

**Coverage as percentage only**:

```bash
make coverage-percentage
```

---

## 8. **Common Pitfalls**

| Problem                              | Cause                        | Fix                                                            |
| ------------------------------------ | ---------------------------- | -------------------------------------------------------------- |
| `0 out of 1 expectation(s) were met` | Mock not called              | Ensure `mockSetup` is correct and payload is valid             |
| `expected 400, got 422`              | Huma schema validation       | Use correct JSON fields; Huma returns 422 for invalid payloads |
| Huma handler not called              | Wrong payload format         | Send flat JSON matching the Go struct                          |
| DB test fails                        | Docker container not running | Ensure TestContainers is installed and Docker is running       |
