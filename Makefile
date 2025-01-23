# Variables
APP_NAME := event-booking
DOCKER_IMAGE := $(APP_NAME)
DOCKER_CONTAINER := $(APP_NAME)-ctr
DOCKERFILE := Dockerfile
PORT := 8080

# Tools
GO := go
TEMPL := templ
AIR := air
DOCKER := docker

# Default target
.PHONY: all
all: docker-run

# Run the application locally with Air for live reload
.PHONY: run
run:
	@if [ -f .air.toml ]; then \
		$(AIR); \
	else \
		$(AIR) init && $(AIR); \
	fi

# Build the Docker image
.PHONY: docker-build
docker-build:
	@echo "Building the Docker image..."
	$(DOCKER) build -t $(DOCKER_IMAGE) -f $(DOCKERFILE) .

# Run the Docker container
.PHONY: docker-run
docker-run:
	@echo "Running the Docker container..."
	$(DOCKER) run -p $(PORT):8080 --rm -v $(PWD):/app -v /app/tmp --name $(DOCKER_CONTAINER) $(DOCKER_IMAGE)

# Stop the Docker container
.PHONY: docker-stop
docker-stop:
	@echo "Stopping the Docker container..."
	$(DOCKER) stop $(DOCKER_CONTAINER) || true

# Push the Docker image to a registry
.PHONY: docker-push
docker-push:
	@echo "Pushing the Docker image to the registry..."
	$(DOCKER) push $(DOCKER_IMAGE)

# Remove dangling Docker images
.PHONY: docker-clean
docker-clean:
	@echo "Cleaning up unused Docker images..."
	$(DOCKER) image prune -f

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	rm -rf ./bin

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Format code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	$(GO) fmt ./...

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	$(GO) test ./...

# Generate Templ components
.PHONY: generate
generate:
	@echo "Generating Templ components..."
	$(TEMPL) generate ./...

# Lint the code
.PHONY: lint
lint:
	@echo "Linting the code..."
	golangci-lint run

# Vendor dependencies
.PHONY: vendor
vendor:
	@echo "Vendoring dependencies..."
	$(GO) mod vendor

# Display help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  run          - Run the application locally with Air for live reload"
	@echo "  docker-build - Build the Docker image"
	@echo "  docker-run   - Run the Docker container"
	@echo "  docker-stop  - Stop the Docker container"
	@echo "  docker-push  - Push the Docker image to the registry"
	@echo "  docker-clean - Clean up unused Docker images"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install project dependencies"
	@echo "  fmt          - Format the code"
	@echo "  test         - Run tests"
	@echo "  generate     - Generate Templ components"
	@echo "  lint         - Lint the code"
	@echo "  vendor       - Vendor dependencies"
