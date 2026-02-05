.PHONY: help dev build preview install clean type-check lint lint-fix format format-check \
        deps tidy check fix ci analyze size

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
PACKAGE_MANAGER := bun
SRC_DIR := src

# ------------------------
# Help
# ------------------------
help:
	@echo "$(BOLD)SkillSpark Frontend - Available Commands$(NC)"
	@echo ""
	@echo "$(BLUE)Development:$(NC)"
	@echo "  make dev               - Start Vite development server (http://localhost:5173)"
	@echo "  make build             - Build for production with Vite"
	@echo "  make preview           - Preview production build locally"
	@echo ""
	@echo "$(BLUE)Code Quality:$(NC)"
	@echo "  make type-check        - Run TypeScript type checking"
	@echo "  make lint              - Run ESLint"
	@echo "  make lint-fix          - Fix ESLint issues"
	@echo "  make format            - Format code with Prettier"
	@echo "  make format-check      - Check code formatting"
	@echo "  make check             - Run all checks (type-check, lint, format-check)"
	@echo "  make fix               - Fix all issues (lint-fix, format)"
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
	@echo "$(BLUE)Analysis:$(NC)"
	@echo "  make analyze           - Analyze bundle size with Vite"
	@echo "  make size              - Check bundle size"
	@echo ""
	@echo "$(YELLOW)Using: Vite + React + TypeScript + Bun$(NC)"

# ------------------------
# Development
# ------------------------
dev:
	@echo "$(BOLD)Starting Vite development server...$(NC)"
	@echo "$(BLUE)Visit: http://localhost:5173$(NC)"
	@$(PACKAGE_MANAGER) run dev

build:
	@echo "$(BOLD)Building for production with Vite...$(NC)"
	@$(PACKAGE_MANAGER) run build
	@echo "$(GREEN)Build complete! Output in dist/$(NC)"

preview:
	@echo "$(BOLD)Starting Vite preview server...$(NC)"
	@$(PACKAGE_MANAGER) run preview

# ------------------------
# Code Quality
# ------------------------
type-check:
	@echo "$(BOLD)Running TypeScript type checking...$(NC)"
	@if $(PACKAGE_MANAGER) run type-check 2>/dev/null; then \
		echo "$(GREEN)Type checking complete$(NC)"; \
	else \
		echo "$(YELLOW)type-check script not found, running tsc directly...$(NC)"; \
		$(PACKAGE_MANAGER) x tsc --noEmit && echo "$(GREEN)Type checking complete$(NC)"; \
	fi

lint:
	@echo "$(BOLD)Running ESLint...$(NC)"
	@if $(PACKAGE_MANAGER) run lint 2>/dev/null; then \
		echo "$(GREEN)Linting complete$(NC)"; \
	else \
		echo "$(YELLOW)lint script not found, running eslint directly...$(NC)"; \
		$(PACKAGE_MANAGER) x eslint $(SRC_DIR) --ext ts,tsx --report-unused-disable-directives --max-warnings 0; \
	fi

lint-fix:
	@echo "$(BOLD)Fixing ESLint issues...$(NC)"
	@if $(PACKAGE_MANAGER) run lint:fix 2>/dev/null; then \
		echo "$(GREEN)Linting issues fixed$(NC)"; \
	else \
		echo "$(YELLOW)lint:fix script not found, running eslint directly...$(NC)"; \
		$(PACKAGE_MANAGER) x eslint $(SRC_DIR) --ext ts,tsx --fix && echo "$(GREEN)Linting issues fixed$(NC)"; \
	fi

format:
	@echo "$(BOLD)Formatting code with Prettier...$(NC)"
	@if $(PACKAGE_MANAGER) run format 2>/dev/null; then \
		echo "$(GREEN)Code formatted$(NC)"; \
	else \
		echo "$(YELLOW)format script not found, running prettier directly...$(NC)"; \
		$(PACKAGE_MANAGER) x prettier --write "$(SRC_DIR)/**/*.{ts,tsx,js,jsx,json,css,scss,md}" && echo "$(GREEN)Code formatted$(NC)"; \
	fi

format-check:
	@echo "$(BOLD)Checking code formatting...$(NC)"
	@if $(PACKAGE_MANAGER) run format:check 2>/dev/null; then \
		echo "$(GREEN)Format check complete$(NC)"; \
	else \
		echo "$(YELLOW)format:check script not found, running prettier directly...$(NC)"; \
		$(PACKAGE_MANAGER) x prettier --check "$(SRC_DIR)/**/*.{ts,tsx,js,jsx,json,css,scss,md}"; \
	fi

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
	@echo "$(BOLD)Installing dependencies with Bun...$(NC)"
	@$(PACKAGE_MANAGER) install
	@echo "$(GREEN)Dependencies installed$(NC)"

deps:
	@echo "$(BOLD)Updating dependencies...$(NC)"
	@$(PACKAGE_MANAGER) update
	@echo "$(GREEN)Dependencies updated$(NC)"

tidy:
	@echo "$(BOLD)Cleaning and reinstalling dependencies...$(NC)"
	@rm -rf $(NODE_MODULES) bun.lockb
	@$(PACKAGE_MANAGER) install
	@echo "$(GREEN)Dependencies reinstalled$(NC)"

# ------------------------
# Cleanup
# ------------------------
clean:
	@echo "$(BOLD)Cleaning Vite build artifacts and cache...$(NC)"
	@rm -rf dist .vite coverage
	@echo "$(GREEN)Build artifacts cleaned$(NC)"

clean-all: clean
	@echo "$(BOLD)Removing node_modules...$(NC)"
	@rm -rf $(NODE_MODULES) bun.lockb
	@echo "$(GREEN)All artifacts and dependencies removed$(NC)"

# ------------------------
# CI/CD
# ------------------------
ci: install type-check lint format-check build
	@echo "$(GREEN)CI pipeline complete$(NC)"

# ------------------------
# Analysis
# ------------------------
analyze:
	@echo "$(BOLD)Analyzing bundle size with Vite...$(NC)"
	@$(PACKAGE_MANAGER) run build -- --mode=analyze

size:
	@echo "$(BOLD)Checking bundle size...$(NC)"
	@du -sh dist/* 2>/dev/null || echo "$(RED)No dist folder found. Run 'make build' first.$(NC)"