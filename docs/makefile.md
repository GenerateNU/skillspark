# SkillSpark Makefile Quick Guide

## What is a Makefile?

A Makefile automates common development tasks. Run `make <command>` to execute predefined workflows. Use `make help` in any directory to see available commands.

## Project Structure

```
skillspark/
├── Makefile           # Docker orchestration
├── backend/Makefile   # Go development & testing
└── frontend/web/Makefile  # React/TypeScript development
```

---

## Essential Commands by Location

### Root Directory (Docker Management)

```bash
make help           # Show all commands
make up             # Start all services
make down           # Stop all services
make logs           # View all logs
make restart        # Restart services

# Individual services
make up-backend     # Start backend only
make logs-frontend  # View frontend logs only
make shell-backend  # Open backend container shell
```

### Backend Directory (Go Development)

```bash
# Testing
make test           # Run all tests
make test-unit      # Run unit tests (fast)
make test-coverage  # Generate coverage report

# Code quality
make lint           # Check code
make lint-fix       # Auto-fix issues
make format         # Format code

# Database
make db-start       # Start local Supabase
make db-reset       # Reset & apply migrations
make db-new NAME=x  # Create new migration
make db-push        # Push to production (⚠️ careful!) only do with permission

# Development
make dev            # Start dev server
make api-gen        # Generate API docs
```

### Frontend Directory (React/TypeScript)

```bash
# Development
make dev            # Start dev server (localhost:5173)
make build          # Production build
make preview        # Preview production build

# Code quality
make check          # Run all checks
make fix            # Auto-fix all issues
make lint           # ESLint
make type-check     # TypeScript checking
make format         # Prettier formatting

# Dependencies
make install        # Install dependencies
make deps           # Update dependencies
```

---

## Common Workflows

### Daily Development

```bash
# Morning
make up                          # Start everything

# During work
cd backend && make test          # Test backend changes
cd frontend/web && make check    # Verify frontend

# End of day
make down                        # Clean shutdown
```

### Before Committing

```bash
# Backend
cd backend
make lint && make test

# Frontend
cd frontend/web
make check
```

### Creating a Database Migration

```bash
cd backend
make db-new NAME=add_feature     # Creates migration file
# Edit the migration file
make db-reset                    # Apply locally
make db-push                     # Push to production
```

---

## Quick Tips

- **Always run `make help`** in any directory to see what's available
- **Combine commands**: `make lint && make test`
- **Debug issues**: `make logs` or `make logs-backend`
- **Clean restart**: `make down && make up`
