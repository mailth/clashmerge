.PHONY: build build-frontend build-backend run clean install-deps

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
	GOOS=linux GOARCH=amd64 go build -o clashmerge .

# Build both frontend and backend
build: build-frontend build-backend
	mkdir -p output/web
	cp clashmerge output/
	cp -r web/out output/web

build-image: build
	VERSION= $(shell cat VERSION)
	docker build -t mailth/clashmerge:$(VERSION) .
	docker push mailth/clashmerge:$(VERSION)
	@echo "Image built and pushed successfully: mailth/clashmerge:$(VERSION)"

tag:
	VERSION= $(shell cat VERSION)
	git tag $(VERSION)
	git push origin $(VERSION)

update-version:
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is not set. Usage: make update-version VERSION=v1.0.0"; \
		exit 1; \
	fi
	echo $(VERSION) > VERSION
	git add VERSION
	git commit -m "update version to $(VERSION)"
	git push

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
