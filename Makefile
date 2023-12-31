# Default target
default_target: help

# Project files with the .go extension
project_files = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Run go project for development
.PHONY: run
run:
	go run $(project_files)

# Install development dependencies
.PHONY: install_dev
install_dev:
	go mod download

# Build project
.PHONY: build
build:
	CGO_ENABLED=0 go build -o bin/fake_api -v
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/fake_api.exe -v

# Help
.PHONY: help
help:
	@echo "Required: Go-^1.21"
	@echo ""
	@echo "Install: go mod download"
	@echo ""
	@echo "Makefile Targets:"
	@echo "... run (Run go project for development)"
	@echo "... install_dev (Install development dependencies)"
	@echo "... build (Build project)"
