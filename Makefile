.PHONY: help docker-up docker-down docker-stop docker-restart docker-logs \
        docker-build docker-clean docker-prune \
        up down stop restart logs build clean \
        up-backend up-frontend logs-backend logs-frontend \
        build-backend build-frontend shell-backend shell-frontend setup-hooks

# Default target - show help
.DEFAULT_GOAL := help

# ------------------------
# OS Detection
# ------------------------
OS := $(shell uname 2>/dev/null || echo Windows_NT)
ifeq ($(OS),Windows_NT)
    DOCKER_COMPOSE = docker compose
    WATCH_CMD = docker watch
else
    DOCKER_COMPOSE = docker compose
    WATCH_FLAG = --watch
endif

# ------------------------
# Colors
# ------------------------
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[1;33m
BLUE := \033[0;34m
CYAN := \033[0;36m
NC := \033[0m
BOLD := \033[1m

# ------------------------
# Help
# ------------------------
help:
	@echo "$(BOLD)SkillSpark - Docker Commands$(NC)"
	@echo ""
	@echo "$(BLUE)Main Commands:$(NC)"
	@echo "  make up                - Start all services with hot reload"
	@echo "  make down              - Stop and remove all containers"
	@echo "  make restart           - Restart all services"
	@echo "  make logs              - View logs from all services"
	@echo "  make build             - Build all services"
	@echo ""
	@echo "$(BLUE)Individual Services:$(NC)"
	@echo "  make up-backend        - Start only backend with hot reload"
	@echo "  make up-frontend       - Start only frontend with hot reload"
	@echo "  make logs-backend      - View backend logs"
	@echo "  make logs-frontend     - View frontend logs"
	@echo "  make build-backend     - Build only backend"
	@echo "  make build-frontend    - Build only frontend"
	@echo ""
	@echo "$(BLUE)Container Management:$(NC)"
	@echo "  make stop              - Stop all services (without removing)"
	@echo "  make shell-backend     - Open shell in backend container"
	@echo "  make shell-frontend    - Open shell in frontend container"
	@echo "  make ps                - Show running containers"
	@echo ""
	@echo "$(BLUE)Cleanup:$(NC)"
	@echo "  make clean             - Remove containers and volumes"
	@echo "  make prune             - Remove all unused Docker resources"
	@echo ""
ifeq ($(OS),Windows_NT)
	@echo "$(YELLOW)Note: On Windows, hot reload uses 'docker watch' in a separate terminal$(NC)"
endif

# ------------------------
# Main Commands
# ------------------------
up:
	@echo "$(BOLD)Starting all services...$(NC)"
ifeq ($(OS),Windows_NT)
	@echo "$(YELLOW)On Windows: Starting containers. Run 'docker watch' in another terminal for hot reload.$(NC)"
	@$(DOCKER_COMPOSE) up --build
else
	@$(DOCKER_COMPOSE) up --build $(WATCH_FLAG)
endif

down:
	@echo "$(BOLD)Stopping all services...$(NC)"
	@$(DOCKER_COMPOSE) down
	@echo "$(GREEN)All services stopped$(NC)"

stop:
	@echo "$(BOLD)Stopping all services (containers preserved)...$(NC)"
	@$(DOCKER_COMPOSE) stop
	@echo "$(GREEN)All services stopped$(NC)"

restart:
	@echo "$(BOLD)Restarting all services...$(NC)"
	@$(DOCKER_COMPOSE) restart
	@echo "$(GREEN)All services restarted$(NC)"

logs:
	@echo "$(BOLD)Viewing logs (Ctrl+C to exit)...$(NC)"
	@$(DOCKER_COMPOSE) logs -f

build:
	@echo "$(BOLD)Building all services...$(NC)"
	@$(DOCKER_COMPOSE) build
	@echo "$(GREEN)Build complete$(NC)"

ps:
	@echo "$(BOLD)Running containers:$(NC)"
	@$(DOCKER_COMPOSE) ps

# ------------------------
# Backend Commands
# ------------------------
up-backend:
	@echo "$(BOLD)Starting backend service...$(NC)"
ifeq ($(OS),Windows_NT)
	@echo "$(YELLOW)On Windows: Starting backend. Run 'docker watch' in another terminal for hot reload.$(NC)"
	@$(DOCKER_COMPOSE) up --build backend
else
	@$(DOCKER_COMPOSE) up --build $(WATCH_FLAG) backend
endif

logs-backend:
	@echo "$(BOLD)Viewing backend logs (Ctrl+C to exit)...$(NC)"
	@$(DOCKER_COMPOSE) logs -f backend

build-backend:
	@echo "$(BOLD)Building backend...$(NC)"
	@$(DOCKER_COMPOSE) build backend
	@echo "$(GREEN)Backend build complete$(NC)"

shell-backend:
	@echo "$(BOLD)Opening shell in backend container...$(NC)"
	@$(DOCKER_COMPOSE) exec backend sh

restart-backend:
	@echo "$(BOLD)Restarting backend...$(NC)"
	@$(DOCKER_COMPOSE) restart backend
	@echo "$(GREEN)Backend restarted$(NC)"

# ------------------------
# Frontend Commands
# ------------------------
up-frontend:
	@echo "$(BOLD)Starting frontend service...$(NC)"
ifeq ($(OS),Windows_NT)
	@echo "$(YELLOW)On Windows: Starting frontend. Run 'docker watch' in another terminal for hot reload.$(NC)"
	@$(DOCKER_COMPOSE) up --build frontend
else
	@$(DOCKER_COMPOSE) up --build $(WATCH_FLAG) frontend
endif

logs-frontend:
	@echo "$(BOLD)Viewing frontend logs (Ctrl+C to exit)...$(NC)"
	@$(DOCKER_COMPOSE) logs -f frontend

build-frontend:
	@echo "$(BOLD)Building frontend...$(NC)"
	@$(DOCKER_COMPOSE) build frontend
	@echo "$(GREEN)Frontend build complete$(NC)"

shell-frontend:
	@echo "$(BOLD)Opening shell in frontend container...$(NC)"
	@$(DOCKER_COMPOSE) exec frontend sh

restart-frontend:
	@echo "$(BOLD)Restarting frontend...$(NC)"
	@$(DOCKER_COMPOSE) restart frontend
	@echo "$(GREEN)Frontend restarted$(NC)"

# ------------------------
# Cleanup Commands
# ------------------------
clean:
	@echo "$(BOLD)Removing all containers and volumes...$(NC)"
	@$(DOCKER_COMPOSE) down -v
	@echo "$(GREEN)Cleanup complete$(NC)"

prune:
	@echo "$(BOLD)Removing all unused Docker resources...$(NC)"
	@echo "$(YELLOW)This will remove:$(NC)"
	@echo "  - All stopped containers"
	@echo "  - All networks not used by containers"
	@echo "  - All dangling images"
	@echo "  - All build cache"
	@echo ""
	@docker system prune -af --volumes
	@echo "$(GREEN)Docker pruned$(NC)"

# ------------------------
# Development Helpers
# ------------------------
rebuild:
	@echo "$(BOLD)Rebuilding all services (no cache)...$(NC)"
	@$(DOCKER_COMPOSE) build --no-cache
	@echo "$(GREEN)Rebuild complete$(NC)"

rebuild-backend:
	@echo "$(BOLD)Rebuilding backend (no cache)...$(NC)"
	@$(DOCKER_COMPOSE) build --no-cache backend
	@echo "$(GREEN)Backend rebuild complete$(NC)"

rebuild-frontend:
	@echo "$(BOLD)Rebuilding frontend (no cache)...$(NC)"
	@$(DOCKER_COMPOSE) build --no-cache frontend
	@echo "$(GREEN)Frontend rebuild complete$(NC)"

# ------------------------
# Watch Helper (for Windows)
# ------------------------
watch:
ifeq ($(OS),Windows_NT)
	@echo "$(BOLD)Starting Docker watch for hot reload...$(NC)"
	@echo "$(YELLOW)Make sure containers are running with 'make up' first$(NC)"
	@docker watch
else
	@echo "$(YELLOW)Watch is built into 'make up' on Unix systems$(NC)"
	@echo "Use 'make up' to start with hot reload enabled"
endif

setup-hooks:
	@echo "Installing git hooks..."
	@ln -sf ../../scripts/hooks/pre-commit .git/hooks/pre-commit
	@echo "✅ Hooks installed!"