# Makefile for Go application

.PHONY: build run test tidy clean dev install-air

# Default environment
ENV ?= dev

# Build the application
build:
	if not exist bin mkdir bin
	go build -v -o bin/app.exe ./cmd/app
	echo Build successful for environment: $(ENV)

# Run the application
run:
	ENV=$(ENV) go run ./cmd/app

# Run in development mode with live reloading
dev:
	air

# Run tests
test:
	go test -v ./...

# Tidy modules
tidy:
	go mod tidy

# Clean build artifacts
clean:
	rm -rf bin/