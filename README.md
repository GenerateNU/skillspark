# SkillSpark 🎯

SkillSpark is a unified platform for students in Thailand to discover and engage with diverse activities and experiences. We're building a product that gives people the chance to explore and easily access all the learning opportunities this country has to offer—a platform to spark their curiosity and develop new skills!

## Tech Stack

### Backend

- **Framework:** [Fiber](https://gofiber.io/) (Express-inspired web framework for Go)
- **Language:** Go 1.22
- **Linting:** golangci-lint v1.64.8
- **Testing:** Go's built-in testing framework with coverage reporting

### Frontend

- **Framework:** React 18
- **Build Tool:** Vite
- **Language:** TypeScript
- **Package Manager:** Bun
- **Linting:** ESLint

### Database

- **Database:** PostgreSQL (hosted on [Supabase](https://supabase.com/))
- **CLI:** Supabase CLI for migrations and local development

### DevOps

- **CI/CD:** GitHub Actions
- **Containerization:** Docker
- **Orchestration:** Docker Compose with hot reload
- **Deployment:** Digital Ocean (coming soon)

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go** 1.22 or higher - [Download](https://golang.org/dl/)
- **Bun** (latest) - [Download](https://bun.sh/)
- **Docker** and Docker Compose - [Download](https://www.docker.com/products/docker-desktop)
- **golangci-lint** v1.64.8 - [Installation](https://golangci-lint.run/welcome/install/)
- **Supabase CLI** - [Installation](https://supabase.com/docs/guides/cli)
- **Make** - Usually pre-installed on macOS/Linux, [Windows installation](https://gnuwin32.sourceforge.net/packages/make.htm)

### macOS Installation

```bash
# Install Go
brew install go

# Install Bun
curl -fsSL https://bun.sh/install | bash

# Install Docker
brew install --cask docker

# Install golangci-lint
brew install golangci-lint

# Install Supabase CLI
brew install supabase/tap/supabase
```

### Windows Installation

```bash
# Install Bun (PowerShell)
powershell -c "irm bun.sh/install.ps1 | iex"

# Install other tools via their respective installers
```

## Getting Started

### Quick Start (Recommended - Using Docker)

```bash
# Clone the repository
git clone <your-repo-url>
cd skillspark

# Start all services with hot reload
make up
```

This will start:

- **Frontend** at http://localhost (Vite dev server with hot reload)
- **Backend** at http://localhost:8080 (Fiber server with hot reload)

**Other useful commands:**

```bash
make down      # Stop all services
make logs      # View logs from all services
make restart   # Restart all services
make help      # See all available commands
```

### Manual Setup (Without Docker)

#### 1. Environment Setup

**Backend Environment Variables**
Create a `.env` file in the `backend/` directory based off the template

**Frontend Environment Variables**
Create a `.env` file in the `frontend/` directory based off the template

#### 2. Database Setup (Supabase)

```bash
cd backend

# Start local Supabase (requires Docker)
make db-start

# Create a new migration (if needed)
make db-new NAME=initial_schema

# Reset database and apply migrations
make db-reset
```

**Local Supabase URLs:**

- Dashboard: http://localhost:54323
- Database: postgresql://postgres:postgres@127.0.0.1:54322/postgres

**Link to Remote Supabase (Production):**

```bash
# Find your project ref in Supabase dashboard URL
make db-link REF=your-project-ref

# Push migrations to remote
make db-push
```

#### 3. Backend Setup

```bash
cd backend

# Install dependencies
make download

# Run tests
make test

# Start development server
make dev
```

Backend will start at http://localhost:8080

#### 4. Frontend Setup

```bash
cd frontend

# Install dependencies
make install

# Start development server
make dev
```

Frontend will start at http://localhost:5173

## Development Workflow

### Using Makefiles

This project uses Makefiles for streamlined development. Run `make help` in any directory to see available commands.

#### Root Directory Commands

```bash
make help              # Show all available commands
make up                # Start all services with Docker
make down              # Stop all services
make logs              # View logs from all services
make up-backend        # Start only backend
make up-frontend       # Start only frontend
make build             # Build all services
make clean             # Remove containers and volumes
```

#### Backend Commands (cd backend)

```bash
# Development
make dev               # Start development server
make build             # Build the application

# Testing
make test              # Run all tests
make test-coverage     # Run tests with coverage report
make test-one TEST=TestName  # Run specific test

# Code Quality
make lint              # Run golangci-lint
make lint-fix          # Fix linting issues
make format            # Format code with gofmt

# Database
make db-start          # Start local Supabase
make db-stop           # Stop local Supabase
make db-reset          # Reset DB and apply migrations
make db-new NAME=...   # Create new migration
make db-status         # Show DB status

# API Documentation
make api-gen           # Generate OpenAPI spec
make api-preview       # Preview API docs at /docs

# Dependencies
make tidy              # Run go mod tidy
make deps              # Update all dependencies
```

#### Frontend Commands (cd frontend)

```bash
# Development
make dev               # Start Vite dev server
make build             # Build for production
make preview           # Preview production build

# Code Quality
make type-check        # Run TypeScript checking
make lint              # Run ESLint
make lint-fix          # Fix ESLint issues
make format            # Format with Prettier
make format-check      # Check formatting
make check             # Run all checks
make fix               # Fix all issues

# Dependencies
make install           # Install dependencies with Bun
make deps              # Update dependencies
make tidy              # Clean and reinstall

# Analysis
make analyze           # Analyze bundle size
make size              # Check bundle size
```

### Pre-commit Hooks

This project uses Husky for pre-commit hooks that automatically:

- Format Go code with `gofmt`
- Run `golangci-lint` on Go files
- Run ESLint on frontend files  
- Tidy Go modules

**Setup hooks:**

```bash
# From root directory
make setup-hooks
```

Or install manually:

```bash
cd frontend
bun install  # Installs husky automatically
```

### Hot Reload in Docker

Docker Compose is configured with hot reload for both services:

- **Backend**: Uses Air for Go hot reload
- **Frontend**: Vite's built-in HMR

Changes to your code will automatically trigger rebuilds!

**Note for Windows users:** Hot reload uses `docker watch`. If `make up` doesn't enable it automatically, run `docker watch` in a separate terminal.

## Testing

### Backend Tests

```bash
cd backend

# Run all tests
make test

# Run with coverage (70% minimum recommended)
make test-coverage

# Run specific test
make test-one TEST=TestActivityHandler

# Run only unit tests
make test-unit

# Run database tests
make test-db

# Clean test cache
make test-clean
```

### Frontend Tests

```bash
cd frontend

# Run tests (if configured)
bun test
```

## Code Quality

### Running All Checks

**Backend:**

```bash
cd backend
make lint && make format-check && make test
```

**Frontend:**

```bash
cd frontend
make check  # Runs type-check, lint, and format-check
```

### Auto-fixing Issues

**Backend:**

```bash
cd backend
make lint-fix && make format
```

**Frontend:**

```bash
cd frontend
make fix  # Runs lint-fix and format
```

## Repository Structure

```
skillspark/
├── backend/              # Go Fiber backend service
│   ├── cmd/
│   │   ├── main.go      # Application entry point
│   │   └── genapi/      # OpenAPI spec generator
│   ├── internal/        # Private application code
│   │   ├── handlers/   # HTTP request handlers
│   │   ├── models/     # Data models
│   │   ├── services/   # Business logic
│   │   ├── storage/    # Database layer
│   │   └── supabase/   # Supabase migrations
│   ├── api/            # OpenAPI specifications
│   ├── Makefile        # Backend commands
│   └── go.mod          # Go dependencies
├── frontend/            # React + Vite frontend
│   ├── src/
│   │   ├── components/ # React components
│   │   ├── pages/      # Page components
│   │   ├── hooks/      # Custom React hooks
│   │   ├── utils/      # Utility functions
│   │   └── api/        # API client
│   ├── public/         # Static assets
│   ├── Makefile        # Frontend commands
│   ├── vite.config.ts  # Vite configuration
│   └── package.json    # Dependencies and scripts
├── docs/               # Documentation
├── .github/
│   └── workflows/      # CI/CD pipelines
├── .husky/             # Git hooks
├── scripts/            # Utility scripts
├── docker-compose.yml  # Service orchestration
├── .golangci.yml       # Linting configuration
├── Makefile            # Root-level commands
└── README.md
```

## CI/CD Pipeline

Our GitHub Actions workflows automatically run on every pull request:

### Backend Checks (`.github/workflows/backend.yaml`)

- ✅ Path filtering (only runs when backend files change)
- ✅ Linting with golangci-lint
- ✅ Unit and integration tests
- ✅ Coverage reporting (70% minimum recommended)
- ✅ Test result summaries posted to PRs
- ✅ Coverage reports uploaded as artifacts

### Triggers

- Pull requests to `main` branch
- Only runs when respective backend or frontend files are modified

### Coverage Enforcement

- Minimum coverage threshold: **70%**
- Coverage reports are generated and posted to PRs
- Tests must pass for PR to be mergeable

## API Documentation

The backend API is built with Go Fiber and includes OpenAPI documentation.

### Accessing API Docs

```bash
# Generate OpenAPI spec
cd backend
make api-gen

# Start the server and view docs
make dev

# Open browser to:
http://localhost:8080/docs
```

### Example Endpoints

```
GET    /api/v1/activities       # List all activities
GET    /api/v1/activities/:id   # Get activity by ID
POST   /api/v1/activities       # Create new activity
PUT    /api/v1/activities/:id   # Update activity
DELETE /api/v1/activities/:id   # Delete activity
```

## Docker

### Development with Docker Compose

```bash
# Start all services
make up

# Start individual services
make up-backend
make up-frontend

# View logs
make logs
make logs-backend
make logs-frontend

# Restart services
make restart

# Stop services
make down

# Clean everything (including volumes)
make clean
```

### Production Builds

```bash
# Build images
make build

# Or build individually
make build-backend
make build-frontend

# Rebuild without cache
make rebuild
```

### Container Management

```bash
# Open shell in container
make shell-backend
make shell-frontend

# View running containers
make ps

# Clean Docker resources
make prune
```

## Deployment

🚧 **Deployment to Digital Ocean is coming soon!**

## Contributing

1. **Create a feature branch from `main`**

```bash
   git checkout -b feature/your-feature-name
```

2. **Make your changes**
   - Follow the style guides in `docs/`
   - Write tests for new features
   - Ensure code passes all checks

3. **Run checks locally**

```bash
   # Backend
   cd backend && make lint && make test
   
   # Frontend  
   cd frontend && make check
```

4. **Commit your changes**
   - Pre-commit hooks will run automatically
   - Fix any issues before committing

5. **Push and create a pull request**

```bash
   git push origin feature/your-feature-name
```

6. **Wait for CI checks and code review**

### Branch Naming Convention

- `feature/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation updates
- `refactor/` - Code refactoring
- `test/` - Test additions or modifications

## Troubleshooting

### Docker Issues

**Port already in use:**

```bash
# Check what's using the port
lsof -i :8080  # macOS/Linux
netstat -ano | findstr :8080  # Windows

# Kill the process or change port in docker-compose.yml
```

**Docker not starting:**

```bash
# Make sure Docker Desktop is running
docker info

# Restart Docker if needed
```

**Hot reload not working:**

```bash
# Windows users: run docker watch in separate terminal
docker watch

# Or restart containers
make restart
```

### Backend Issues

**Database connection failed:**

```bash
# Check Supabase is running
make db-status

# Restart Supabase
make db-stop && make db-start

# Check .env file has correct credentials
```

**Tests failing:**

```bash
# Clean test cache
make test-clean

# Run tests again
make test
```

**golangci-lint version mismatch:**

```bash
# Install specific version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8
```

### Frontend Issues

**Module not found:**

```bash
cd frontend
make tidy  # Clean and reinstall
```

**Build fails:**

```bash
# Clear Bun cache
rm -rf node_modules bun.lockb
make install
```

**API calls failing:**

```bash
# Check backend is running
curl http://localhost:8080/health

# Verify VITE_API_URL in .env
```

**Bun issues:**

```bash
# Update Bun
bun upgrade

# Clear Bun cache
bun pm cache rm
```

### Common Make Command Issues

**Make not found (Windows):**

- Install Make from [GnuWin32](http://gnuwin32.sourceforge.net/packages/make.htm)
- Or use Git Bash which includes Make

**Permission denied:**

```bash
# macOS/Linux: Make scripts executable
chmod +x scripts/*.sh
```

## Useful Resources

- [Go Fiber Documentation](https://docs.gofiber.io/)
- [React Documentation](https://react.dev/)
- [Vite Documentation](https://vitejs.dev/)
- [Bun Documentation](https://bun.sh/docs)
- [Supabase Documentation](https://supabase.com/docs)
- [Docker Compose Documentation](https://docs.docker.com/compose/)

## Team

Built with ❤️ by the SkillSpark team

## License

[Add license information]