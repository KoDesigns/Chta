.PHONY: build install clean test dev

# Build the binary
build:
	go build -o chta main.go

# Install globally (requires sudo)
install: build
	sudo mv chta /usr/local/bin/

# Install locally in user's bin directory
install-user: build
	mkdir -p ~/bin
	mv chta ~/bin/
	@echo "Add ~/bin to your PATH if not already there"

# Clean build artifacts
clean:
	rm -f chta

# Run tests
test:
	go test ./...

# Development mode - run without building
dev:
	go run main.go

# Show help
help:
	@echo "Chta - Fast CLI Cheat Sheet Tool"
	@echo ""
	@echo "Usage:"
	@echo "  make build        Build the binary"
	@echo "  make install      Install globally (requires sudo)"
	@echo "  make install-user Install in ~/bin"
	@echo "  make clean        Remove build artifacts"
	@echo "  make test         Run tests"
	@echo "  make dev          Run in development mode"
	@echo "  make help         Show this help"

# Default target
all: build 