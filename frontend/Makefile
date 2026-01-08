.PHONY: help dev build preview install clean type-check lint lint-fix format format-check \
        deps tidy

# Default target - show help
.DEFAULT_GOAL := help

# ------------------------
# Colors
# ------------------------
RED := \033[0;31m
GREEN := \033[0;32m
YELLOW := \033[1;33m
BLUE := \033[0;34m
NC := \033[0m
BOLD := \033[1m

# ------------------------
# Configuration
# ------------------------
NODE_MODULES := node_modules
PACKAGE_MANAGER := npm

# Detect if pnpm or yarn is available
ifeq ($(shell command -v pnpm 2>/dev/null),)
    ifeq ($(shell command -v yarn 2>/dev/null),)
        PACKAGE_MANAGER := npm
    else
        PACKAGE_MANAGER := yarn
    endif
else
    PACKAGE_MANAGER := pnpm
endif

# ------------------------
# Help
# ------------------------
help:
	@echo "$(BOLD)SkillSpark Frontend - Available Commands$(NC)"
	@echo ""
	@echo "$(BLUE)Development:$(NC)"
	@echo "  make dev               - Start development server"
	@echo "  make build             - Build for production"
	@echo "  make preview           - Preview production build locally"
	@echo ""
	@echo "$(BLUE)Code Quality:$(NC)"
	@echo "  make type-check        - Run TypeScript type checking"
	@echo "  make lint              - Run ESLint"
	@echo "  make lint-fix          - Fix ESLint issues"
	@echo "  make format            - Format code with Prettier"
	@echo "  make format-check      - Check code formatting"
	@echo ""
	@echo "$(BLUE)Dependencies:$(NC)"
	@echo "  make install           - Install dependencies"
	@echo "  make deps              - Update dependencies"
	@echo "  make tidy              - Clean and reinstall dependencies"
	@echo ""
	@echo "$(BLUE)Cleanup:$(NC)"
	@echo "  make clean             - Remove build artifacts and cache"
	@echo "  make clean-all         - Remove build artifacts and node_modules"
	@echo ""
	@echo "$(YELLOW)Using package manager: $(PACKAGE_MANAGER)$(NC)"

# ------------------------
# Development
# ------------------------
dev:
	@echo "$(BOLD)Starting development server...$(NC)"
	@$(PACKAGE_MANAGER) run dev

build:
	@echo "$(BOLD)Building for production...$(NC)"
	@$(PACKAGE_MANAGER) run build
	@echo "$(GREEN)Build complete! Output in dist/$(NC)"

preview:
	@echo "$(BOLD)Starting preview server...$(NC)"
	@$(PACKAGE_MANAGER) run preview

# ------------------------
# Code Quality
# ------------------------
type-check:
	@echo "$(BOLD)Running TypeScript type checking...$(NC)"
	@$(PACKAGE_MANAGER) run type-check || $(PACKAGE_MANAGER) exec tsc --noEmit
	@echo "$(GREEN)Type checking complete$(NC)"

lint:
	@echo "$(BOLD)Running ESLint...$(NC)"
	@$(PACKAGE_MANAGER) run lint || $(PACKAGE_MANAGER) exec eslint . --ext ts,tsx --report-unused-disable-directives --max-warnings 0

lint-fix:
	@echo "$(BOLD)Fixing ESLint issues...$(NC)"
	@$(PACKAGE_MANAGER) run lint:fix || $(PACKAGE_MANAGER) exec eslint . --ext ts,tsx --fix
	@echo "$(GREEN)Linting issues fixed$(NC)"

format:
	@echo "$(BOLD)Formatting code with Prettier...$(NC)"
	@$(PACKAGE_MANAGER) run format || $(PACKAGE_MANAGER) exec prettier --write "src/**/*.{ts,tsx,js,jsx,json,css,scss,md}"
	@echo "$(GREEN)Code formatted$(NC)"

format-check:
	@echo "$(BOLD)Checking code formatting...$(NC)"
	@$(PACKAGE_MANAGER) run format:check || $(PACKAGE_MANAGER) exec prettier --check "src/**/*.{ts,tsx,js,jsx,json,css,scss,md}"

# Check all (type-check, lint, format-check)
check: type-check lint format-check
	@echo "$(GREEN)All checks passed!$(NC)"

# Fix all (lint-fix, format)
fix: lint-fix format
	@echo "$(GREEN)All fixes applied!$(NC)"

# ------------------------
# Dependencies
# ------------------------
install:
	@echo "$(BOLD)Installing dependencies...$(NC)"
	@$(PACKAGE_MANAGER) install
	@echo "$(GREEN)Dependencies installed$(NC)"

deps:
	@echo "$(BOLD)Updating dependencies...$(NC)"
ifeq ($(PACKAGE_MANAGER),npm)
	@npm update
else ifeq ($(PACKAGE_MANAGER),yarn)
	@yarn upgrade
else
	@pnpm update
endif
	@echo "$(GREEN)Dependencies updated$(NC)"

tidy:
	@echo "$(BOLD)Cleaning and reinstalling dependencies...$(NC)"
	@rm -rf $(NODE_MODULES) package-lock.json yarn.lock pnpm-lock.yaml
	@$(PACKAGE_MANAGER) install
	@echo "$(GREEN)Dependencies reinstalled$(NC)"

# ------------------------
# Cleanup
# ------------------------
clean:
	@echo "$(BOLD)Cleaning build artifacts and cache...$(NC)"
	@rm -rf dist .vite coverage
	@echo "$(GREEN)Build artifacts cleaned$(NC)"

clean-all: clean
	@echo "$(BOLD)Removing node_modules...$(NC)"
	@rm -rf $(NODE_MODULES) package-lock.json yarn.lock pnpm-lock.yaml
	@echo "$(GREEN)All artifacts and dependencies removed$(NC)"

# ------------------------
# CI/CD
# ------------------------
ci: install type-check lint format-check test build
	@echo "$(GREEN)CI pipeline complete$(NC)"

# ------------------------
# Analysis
# ------------------------
analyze:
	@echo "$(BOLD)Analyzing bundle size...$(NC)"
	@$(PACKAGE_MANAGER) run build -- --mode=analyze

size:
	@echo "$(BOLD)Checking bundle size...$(NC)"
	@du -sh dist/* 2>/dev/null || echo "No dist folder found. Run 'make build' first."