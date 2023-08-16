# Default target
default_target: help

project_files = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# Run go project for development
.PHONY: run
run:
	go run $(project_files)

# Install development dependencies
.PHONY: install_dev
install_dev:
	go mod download

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
