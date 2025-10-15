.PHONY: build build-frontend build-backend run clean install-deps version

# Default target
all: build

# Install frontend dependencies
install-deps:
	cd web && pnpm install

# Build frontend
build-frontend:
	cd web && pnpm run build

# Build backend
build-backend:
	@echo "Building backend..."; \
	GOOS=linux GOARCH=amd64 go build -o clashmerge .

# Build both frontend and backend
build: build-frontend build-backend
	mkdir -p output/web
	cp clashmerge output/
	cp -r web/out output/web



# Prevent make from treating version arguments as targets
%:
	@:

# Run in development mode
run:
	LOG_LEVEL=debug \
	DATA_DIR="./data" \
	CONFIG_DIR="./config" \
	go run .

# Clean build artifacts
clean:
	rm -f clashmerge
	rm -rf web/dist
	rm -rf web/node_modules

# Test
test:
	go test -v ./...
