.PHONY: all build install dev run-backend run-frontend clean db-up db-down db-restart db-logs db-shell build-linux

# 默认目标
all: build

# ==============================================================================
# Database (Docker)
# ==============================================================================

db-up:
	@echo "Starting database..."
	cd backend && docker compose up -d

db-down:
	@echo "Stopping database..."
	cd backend && docker compose down

db-restart: db-down db-up

db-logs:
	@echo "Showing database logs..."
	cd backend && docker compose logs -f

db-shell:
	@echo "Entering database shell..."
	cd backend && docker compose exec postgres psql -U cook -d cook_db

# ==============================================================================
# Development
# ==============================================================================

# 安装所有依赖
install: install-backend install-frontend

install-backend:
	@echo "Installing backend dependencies..."
	cd backend && go mod tidy

install-frontend:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install

# 启动开发环境 (建议开启两个终端分别运行，或者使用 make -j2 dev-all)
dev:
	@echo "建议开启两个终端分别运行:"
	@echo "  make run-backend"
	@echo "  make run-frontend"
	@echo "或者使用并行模式: make -j2 dev-all"

dev-all: run-backend run-frontend

run-backend:
	@echo "Starting Go backend..."
	@echo "Ensuring database is up..."
	$(MAKE) db-up
	cd backend && go run main.go

run-frontend:
	@echo "Starting React frontend..."
	cd frontend && npm run dev

# ==============================================================================
# Build
# ==============================================================================

build: build-backend build-frontend

build-backend:
	@echo "Building backend..."
	cd backend && go build -o server main.go

# Cross-compile for Linux (for server deployment)
build-linux:
	@echo "Building backend for Linux..."
	cd backend && GOOS=linux GOARCH=amd64 go build -o server-linux main.go

build-frontend:
	@echo "Building frontend..."
	cd frontend && npm run build

# ==============================================================================
# Clean
# ==============================================================================

clean:
	rm -f backend/server backend/server-linux
	rm -rf frontend/dist
