# SkillSpark Quick Start Guide

This guide will help you get started with the SkillSpark project quickly.

## Table of Contents

- [Prerequisites](#prerequisites)
- [VS Code Setup](#vs-code-setup)
- [Installation](#installation)
- [Running the Application](#running-the-application)
  - [Option 1: Using Docker (Recommended)](#option-1-using-docker-recommended)
  - [Option 2: Running Services Independently](#option-2-running-services-independently)
- [Verification](#verification)
- [Available Commands](#available-commands)
- [Troubleshooting](#troubleshooting)

---

## Prerequisites

Before you begin, ensure you have the following installed on your system:

### Required Software

1. **Docker & Docker Compose**

   - Download from: https://www.docker.com/products/docker-desktop
   - Version: Docker 20.10+ and Docker Compose 2.0+
   - Verify installation:
     ```bash
     docker --version
     docker compose version
     ```

2. **Node.js**

   - Version: 18.x or higher
   - Download from: https://nodejs.org/
   - Verify installation:
     ```bash
     node --version
     npm --version
     ```

3. **Bun** (JavaScript runtime & package manager)

   - Version: 1.0+
   - Install:
     ```bash
     curl -fsSL https://bun.sh/install | bash
     ```
   - Or visit: https://bun.sh/docs/installation
   - Verify installation:
     ```bash
     bun --version
     ```

4. **Go (Golang)**

   - Version: 1.24.0 or higher
   - Download from: https://go.dev/dl/
   - Verify installation:
     ```bash
     go version
     ```

5. **Supabase CLI** (for database management)

   - Install via npm:
     ```bash
     npm install -g supabase
     ```
   - Or visit: https://supabase.com/docs/guides/cli
   - Verify installation:
     ```bash
     supabase --version
     ```

6. **Make** (build automation)
   - **Linux**: Usually pre-installed
   - **macOS**: Install Xcode Command Line Tools
     ```bash
     xcode-select --install
     ```
   - **Windows**: Install via Chocolatey
     ```bash
     choco install make
     ```

---

## VS Code Setup

### Recommended Extensions

Install these essential VS Code extensions for the best development experience:

1. **Go** (`golang.go`)

   - Official Go language support
   - Provides IntelliSense, debugging, and testing

2. **ESLint** (`dbaeumer.vscode-eslint`)

   - JavaScript/TypeScript linting
   - Auto-fix on save

3. **Prettier - Code formatter** (`esbenp.prettier-vscode`)

   - Consistent code formatting
   - Works with JavaScript, TypeScript, JSON, CSS, etc.

4. **Tailwind CSS IntelliSense** (`bradlc.vscode-tailwindcss`)

   - Autocomplete for Tailwind classes
   - Syntax highlighting

5. **Docker** (`ms-azuretools.vscode-docker`)
   - Manage Docker containers and images
   - View logs and inspect containers

---

## Installation

1. **Clone the repository** (if you haven't already):

   ```bash
   git clone <repository-url>
   cd skillspark
   ```

2. **Set up environment variables for backend**:

   ```bash
   cd backend
   cp env.sample .env
   ```

   Edit `.env` and fill in the required values:

3. **Install Git hooks** (optional but recommended):
   ```bash
   make setup-hooks
   ```

---

## Running the Application

### Option 1: Using Docker (Recommended)

This is the easiest way to run both frontend and backend together.

#### Start All Services

```bash
make up
```

This command will:

- Build both frontend and backend Docker images
- Start the containers with hot reload enabled
- Set up networking between services

**Access your application:**

- Frontend: http://localhost
- Backend API: http://localhost:8080
- Backend API Docs: http://localhost:8080/docs

#### View Logs

```bash
# All services
make logs

# Backend only
make logs-backend

# Frontend only
make logs-frontend
```

#### Stop All Services

```bash
# Stop and remove containers
make down

# Stop without removing (preserves state)
make stop
```

#### Restart Services

```bash
make restart
```

---

### Option 2: Running Services Independently

For development, you might want to run services separately for more control.

#### Running Backend Independently

1. **Navigate to backend directory**:

   ```bash
   cd backend
   ```

2. **Install Go dependencies**:

   ```bash
   make deps
   ```

3. **Start local Supabase** (PostgreSQL database):

   ```bash
   make db-start
   ```

   This will start a local Supabase instance on Docker. Note the database URL:

   - Dashboard: http://localhost:54323
   - Database: `postgresql://postgres:postgres@127.0.0.1:54322/postgres`

4. **Apply database migrations**:

   ```bash
   make db-reset
   ```

5. **Run the backend server**:

   ```bash
   make dev
   ```

   The backend will be available at http://localhost:8080

**Backend Commands:**

```bash
make help              # Show all available commands
make test              # Run all tests
make lint              # Run linter
make format            # Format code
make db-status         # Check database status
```

#### Running Frontend Independently

1. **Navigate to frontend directory**:

   ```bash
   cd frontend/web
   ```

2. **Install dependencies**:

   ```bash
   bun install
   ```

3. **Start the development server**:

   ```bash
   make dev
   # or directly:
   bun run dev
   ```

   The frontend will be available at http://localhost:5173

**Frontend Commands:**

```bash
make help              # Show all available commands
make build             # Build for production
make lint              # Run ESLint
make format            # Format code with Prettier
make type-check        # Run TypeScript type checking
```

#### Running with Docker (Individual Services)

**Backend only:**

```bash
# From project root
make up-backend
```

**Frontend only:**

```bash
# From project root
make up-frontend
```

---

## Verification

### Check if Everything is Running

1. **Check Docker containers** (if using Docker):

   ```bash
   make ps
   ```

   You should see both `backend` and `web-frontend` containers running.

2. **Test the backend**:

   ```bash
   curl http://localhost:8080
   ```

3. **Test the frontend**:

   - Open http://localhost (Docker) or http://localhost:5173 (local dev)
   - You should see the SkillSpark application

4. **Check API documentation**:
   - Visit http://localhost:8080/docs
   - You should see the interactive API documentation

### Run Tests

**Backend tests:**

```bash
cd backend
make test
```

**Frontend tests:**

```bash
cd frontend/web
bun test
```

---

## Available Commands

### Root Level Commands (Docker)

```bash
make help              # Show all available commands
make up                # Start all services with hot reload
make down              # Stop and remove all containers
make restart           # Restart all services
make logs              # View logs from all services
make build             # Build all services
make up-backend        # Start only backend
make up-frontend       # Start only frontend
make clean             # Remove containers and volumes
```

### Backend Commands

```bash
cd backend
make help              # Show all available backend commands
make dev               # Run development server
make test              # Run all tests
make lint              # Run linter
make format            # Format code
make db-start          # Start local Supabase
make db-stop           # Stop local Supabase
make db-reset          # Reset database and apply migrations
make db-new NAME=...   # Create new migration
make api-gen           # Generate OpenAPI spec
```

### Frontend Commands

```bash
cd frontend/web
make help              # Show all available frontend commands
make dev               # Start development server
make build             # Build for production
make preview           # Preview production build
make lint              # Run ESLint
make format            # Format code with Prettier
make type-check        # Run TypeScript type checking
make check             # Run all checks (type, lint, format)
make fix               # Fix all issues
```

---

## Troubleshooting

1. Check the logs: `make logs`
2. Read the documentation in `docs/` folder
3. Check existing GitHub issues
4. Ask the team on your communication channel

---

## Next Steps

- Read the [Contributing Guide](CONTRIBUTING.md)
- Check out the [Documentation](docs/)
- Explore the [API Documentation](http://localhost:8080/docs) when the server is running
- Set up your IDE with the recommended extensions

Reach out to your TLs if you have any issues <3
