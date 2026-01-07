# XValidator Package Documentation

## Overview

`xvalidator` is a Go validation wrapper around the popular `go-playground/validator` library that provides business-level request validation with user-friendly error messages. It simplifies the process of validating incoming API requests and converting technical validation errors into human-readable messages.

## Purpose

- **Business-Level Validation**: Validate request data before it reaches your business logic or database
- **User-Friendly Errors**: Convert technical validation tags into clear, actionable error messages
- **Type Safety**: Leverage Go's struct tags for declarative validation rules
- **Centralized Validation**: Maintain consistent validation logic across your application

## Package Structure

### Core Components

#### `XValidator`

The main validator struct that wraps `validator.Validate`.

```go
type XValidator struct {
    Validator *validator.Validate
}
```

#### `ErrorResponse`

Represents a single validation error with detailed information.

```go
type ErrorResponse struct {
    Error       bool        // Always true when validation fails
    FailedField string      // Name of the struct field that failed
    Tag         string      // Validation tag that failed (e.g., "required", "min:3")
    Value       interface{} // The actual value that failed validation
}
```

#### `GlobalErrorHandlerResp`

Standard error response format for API responses.

```go
type GlobalErrorHandlerResp struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}
```

## How to Use

### Step 1: Define Your Input Struct with Validation Tags

Use struct tags to declare validation rules on your input models:

```go
type CreateStudentInput struct {
    FirstName   string   `json:"first_name" validate:"required,min=1,max=100"`
    LastName    string   `json:"last_name" validate:"required,min=1,max=100"`
    DOB         *string  `json:"dob,omitempty" validate:"omitempty,datetime=2006-01-02"`
    TherapistID string   `json:"therapist_id" validate:"required,uuid"`
    Grade       *int     `json:"grade,omitempty" validate:"omitempty,oneof=0 1 2 3 4 5 6 7 8 9 10 11 12"`
    IEP         []string `json:"iep,omitempty"`
    SchoolID    int      `json:"school_id" validate:"required,min=1"`
}
```

### Step 2: Validate in Your Handler

Use the global `Validator` instance in your HTTP handlers:

```go
func (h *Handler) AddStudent(c *fiber.Ctx) error {
    var req models.CreateStudentInput
    
    // Parse request body
    if err := c.BodyParser(&req); err != nil {
        return errs.InvalidJSON("Invalid JSON format")
    }
    
    // Validate using xvalidator
    if validationErrors := xvalidator.Validator.Validate(req); len(validationErrors) > 0 {
        return errs.InvalidRequestData(xvalidator.ConvertToMessages(validationErrors))
    }
    
    // Continue with business logic...
    // Your validated data is now safe to use
}
```

### Step 3: Handle Validation Errors

The `ConvertToMessages` function transforms technical errors into user-friendly messages:

```go
// Input: validationErrors from Validator.Validate()
// Output: map[string]string with field names as keys and error messages as values

errorMap := xvalidator.ConvertToMessages(validationErrors)
// Example output:
// {
//   "first_name": "FirstName is required",
//   "email": "Email must be a valid email address",
//   "age": "Age must be greater than or equal to 18"
// }
```

## Supported Validation Tags

| Tag | Description | Example | Error Message |
|-----|-------------|---------|---------------|
| `required` | Field must be present and non-zero | `validate:"required"` | "FirstName is required" |
| `min` | Minimum length (strings) or value (numbers) | `validate:"min=3"` | "FirstName is too short" |
| `max` | Maximum length (strings) or value (numbers) | `validate:"max=100"` | "FirstName is too long" |
| `gte` | Greater than or equal to | `validate:"gte=18"` | "Age must be greater than or equal to 18" |
| `lte` | Less than or equal to | `validate:"lte=120"` | "Age must be less than or equal to 120" |
| `email` | Valid email format | `validate:"email"` | "Email must be a valid email address" |
| `url` | Valid URL format | `validate:"url"` | "Website must be a valid URL" |
| `uuid` | Valid UUID format | `validate:"uuid"` | "TherapistID is invalid" |
| `datetime` | Valid datetime with format | `validate:"datetime=2006-01-02"` | "DOB is invalid" |
| `oneof` | Value must be one of the specified values | `validate:"oneof=0 1 2 3"` | "Grade is invalid" |
| `gtfield` | Greater than another field | `validate:"gtfield=StartDate"` | "EndDate must be greater than StartDate" |
| `omitempty` | Skip validation if field is empty | `validate:"omitempty,email"` | (validates only if provided) |

## Common Validation Patterns

### Required Fields

```go
type LoginInput struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}
```

### Optional Fields with Validation

```go
type UpdateProfileInput struct {
    Bio      *string `json:"bio,omitempty" validate:"omitempty,max=500"`
    Website  *string `json:"website,omitempty" validate:"omitempty,url"`
    Age      *int    `json:"age,omitempty" validate:"omitempty,gte=13,lte=120"`
}
```

### Enum Validation

```go
type CreateTaskInput struct {
    Status string `json:"status" validate:"required,oneof=pending in_progress completed"`
    Priority int  `json:"priority" validate:"required,oneof=1 2 3 4 5"`
}
```

### Date Validation

```go
type ScheduleInput struct {
    StartDate string `json:"start_date" validate:"required,datetime=2006-01-02"`
    EndDate   string `json:"end_date" validate:"required,datetime=2006-01-02"`
}
```

## Best Practices

1. **Validate Early**: Always validate input data at the handler level before processing
2. **Use Descriptive Field Names**: Field names in error messages come from struct field names, so use clear naming
3. **Combine with Custom Validation**: For complex business rules, perform xvalidator checks first, then add custom validation
4. **Handle Pointer Fields**: Use `omitempty` tag for optional pointer fields to avoid nil pointer issues
5. **Keep Validation Rules in Input Structs**: Separate input DTOs from domain models to keep validation logic clear

## Integration with Error Handling

Combine with your custom error package for consistent API responses:

```go
if validationErrors := xvalidator.Validator.Validate(req); len(validationErrors) > 0 {
    return errs.InvalidRequestData(xvalidator.ConvertToMessages(validationErrors))
}
```

This produces responses like:

```json
{
    "success": false,
    "message": "Invalid request data",
    "errors": {
        "email": "Email must be a valid email address",
        "age": "Age must be greater than or equal to 18"
    }
}
```

## Extending Error Messages

To add support for additional validation tags, update the `getErrorMessage` function:

```go
func getErrorMessage(err ErrorResponse) string {
    field := err.FailedField
    tag := err.Tag
    
    switch tag {
    case "required":
        return fmt.Sprintf("%s is required", field)
    case "min":
        return fmt.Sprintf("%s is too short", field)
    // Add your custom cases here
    case "alphanum":
        return fmt.Sprintf("%s must contain only letters and numbers", field)
    default:
        return fmt.Sprintf("%s is invalid", field)
    }
}
```

## Summary

The `xvalidator` package provides a clean, maintainable way to handle request validation in your Go applications. By leveraging struct tags and providing user-friendly error messages, it reduces boilerplate code and improves the developer experience while ensuring data integrity at the API boundary.
