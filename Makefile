.PHONY: build install clean test dev help completion uninstall install-local check-deps diagnose

# Version information
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME = $(shell date -u '+%Y-%m-%d_%H:%M:%S')
LDFLAGS = -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME)"

# Build the binary with version info
build:
	@echo "🔨 Building chta $(VERSION)..."
	go build $(LDFLAGS) -o chta main.go
	@echo "✅ Build complete!"

# Smart install - checks PATH and suggests best option
install: build check-path
	@if command -v chta >/dev/null 2>&1; then \
		echo "⚠️  chta is already installed. Use 'make uninstall' first or 'make install-force'"; \
		exit 1; \
	fi
	sudo mv chta /usr/local/bin/
	@echo "✅ chta installed globally to /usr/local/bin/"
	@echo "💡 Run 'chta --help' to get started"
	@echo "🎯 Try: chta git"

# Force install (overwrites existing)
install-force: build
	@echo "🔄 Force installing chta..."
	sudo mv chta /usr/local/bin/
	@echo "✅ chta installed/updated globally!"

# Install locally in user directory  
install-local: build
	@mkdir -p ~/bin
	mv chta ~/bin/
	@echo "✅ chta installed locally to ~/bin/"
	@echo ""
	@echo "🔍 Checking PATH configuration..."
	@if ! echo $$PATH | grep -q "$$HOME/bin"; then \
		echo "⚠️  ~/bin is not in your PATH"; \
		echo ""; \
		echo "💡 To fix this, run ONE of the following:"; \
		echo ""; \
		if [ "$$SHELL" = "/bin/zsh" ] || [ "$$SHELL" = "/usr/bin/zsh" ]; then \
			echo "📝 For Zsh (macOS default):"; \
			echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.zshrc"; \
			echo "   source ~/.zshrc"; \
		elif [ "$$SHELL" = "/bin/bash" ] || [ "$$SHELL" = "/usr/bin/bash" ]; then \
			echo "📝 For Bash:"; \
			echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.bashrc"; \
			echo "   source ~/.bashrc"; \
		elif command -v fish >/dev/null 2>&1; then \
			echo "📝 For Fish:"; \
			echo "   fish_add_path ~/bin"; \
		else \
			echo "📝 For your shell ($$SHELL):"; \
			echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.profile"; \
			echo "   source ~/.profile"; \
		fi; \
		echo ""; \
		echo "🔄 Then restart your terminal or run the source command"; \
		echo "✅ After that, you can use 'chta --help' from anywhere"; \
	else \
		echo "✅ ~/bin is already in your PATH - you're all set!"; \
		echo "💡 Try: chta --help"; \
	fi

# Auto-install: chooses best install method
auto-install: build
	@echo "🔍 Detecting best installation method..."
	@if [ -w /usr/local/bin ]; then \
		echo "🎯 Auto-installing globally (you have write access)..."; \
		mv chta /usr/local/bin/; \
		echo "✅ chta installed globally to /usr/local/bin/!"; \
	elif command -v sudo >/dev/null 2>&1; then \
		echo "🎯 Auto-installing globally (using sudo)..."; \
		echo "💡 You may be prompted for your password"; \
		if sudo mv chta /usr/local/bin/; then \
			echo "✅ chta installed globally to /usr/local/bin/!"; \
		else \
			echo "❌ Global install failed, trying user install..."; \
			$(MAKE) install-local; \
			exit 0; \
		fi; \
	else \
		echo "🎯 Auto-installing locally (no sudo available)..."; \
		mkdir -p ~/bin; \
		mv chta ~/bin/; \
		echo "✅ chta installed locally to ~/bin/"; \
		if ! echo $$PATH | grep -q "$$HOME/bin"; then \
			echo "⚠️  ~/bin not in PATH - see instructions above"; \
			$(MAKE) show-path-help; \
		fi; \
	fi
	@echo ""
	@echo "🎉 Installation complete!"
	@echo "💡 Try: chta --help"

# Helper target to show PATH setup instructions
show-path-help:
	@echo ""
	@echo "📝 PATH Setup Instructions:"
	@if [ "$$SHELL" = "/bin/zsh" ] || [ "$$SHELL" = "/usr/bin/zsh" ]; then \
		echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.zshrc && source ~/.zshrc"; \
	elif [ "$$SHELL" = "/bin/bash" ] || [ "$$SHELL" = "/usr/bin/bash" ]; then \
		echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.bashrc && source ~/.bashrc"; \
	else \
		echo "   echo 'export PATH=\"\$$HOME/bin:\$$PATH\"' >> ~/.profile && source ~/.profile"; \
	fi

# Generate shell completions
completion: build
	@echo "🔧 Setting up shell completions..."
	@mkdir -p ~/.local/share/bash-completion/completions
	@mkdir -p ~/.config/fish/completions  
	@mkdir -p ~/.zsh/completions
	./chta completion bash > ~/.local/share/bash-completion/completions/chta
	./chta completion fish > ~/.config/fish/completions/chta.fish
	./chta completion zsh > ~/.zsh/completions/_chta
	@echo "✅ Shell completions installed!"
	@echo "💡 Restart your shell or run: source ~/.bashrc"

# Uninstall chta
uninstall:
	@echo "🗑️  Uninstalling chta..."
	@if [ -f /usr/local/bin/chta ]; then \
		sudo rm /usr/local/bin/chta; \
		echo "✅ Removed /usr/local/bin/chta"; \
	fi
	@if [ -f ~/bin/chta ]; then \
		rm ~/bin/chta; \
		echo "✅ Removed ~/bin/chta"; \
	fi
	@if [ -f ~/.local/share/bash-completion/completions/chta ]; then \
		rm ~/.local/share/bash-completion/completions/chta; \
		echo "✅ Removed bash completions"; \
	fi
	@echo "👋 chta uninstalled"

# Check dependencies and environment
check-deps:
	@echo "🔍 Checking dependencies..."
	@if ! command -v go >/dev/null 2>&1; then \
		echo "❌ Go is not installed. Please install Go 1.21+"; \
		exit 1; \
	fi
	@GO_VERSION=$$(go version | grep -o 'go[0-9]\+\.[0-9]\+' | sed 's/go//'); \
	if [ "$$(echo "$$GO_VERSION < 1.21" | bc -l 2>/dev/null || echo 0)" = "1" ]; then \
		echo "⚠️  Go version $$GO_VERSION detected. Go 1.21+ recommended"; \
	else \
		echo "✅ Go version $$GO_VERSION"; \
	fi

# Diagnose installation issues
diagnose:
	@echo "🔍 Chta Installation Diagnostics"
	@echo "================================"
	@echo ""
	@echo "🖥️  System Info:"
	@echo "   OS: $$(uname -s)"
	@echo "   Shell: $$SHELL"
	@echo "   User: $$USER"
	@echo ""
	@echo "📁 Installation Locations:"
	@if [ -f ~/bin/chta ]; then \
		echo "   ✅ User install: ~/bin/chta"; \
		ls -la ~/bin/chta; \
	else \
		echo "   ❌ User install: ~/bin/chta (not found)"; \
	fi
	@if [ -f /usr/local/bin/chta ]; then \
		echo "   ✅ Global install: /usr/local/bin/chta"; \
		ls -la /usr/local/bin/chta; \
	else \
		echo "   ❌ Global install: /usr/local/bin/chta (not found)"; \
	fi
	@echo ""
	@echo "🛤️  PATH Analysis:"
	@echo "   PATH: $$PATH"
	@echo ""
	@if echo $$PATH | grep -q "$$HOME/bin"; then \
		echo "   ✅ ~/bin is in PATH"; \
	else \
		echo "   ❌ ~/bin is NOT in PATH"; \
	fi
	@if echo $$PATH | grep -q "/usr/local/bin"; then \
		echo "   ✅ /usr/local/bin is in PATH"; \
	else \
		echo "   ❌ /usr/local/bin is NOT in PATH"; \
	fi
	@echo ""
	@echo "🔍 Command Detection:"
	@if command -v chta >/dev/null 2>&1; then \
		echo "   ✅ chta found: $$(which chta)"; \
		echo "   📊 Version: $$(chta --version 2>/dev/null || echo 'Version check failed')"; \
	else \
		echo "   ❌ chta command not found"; \
	fi
	@echo ""
	@echo "💡 Recommendations:"
	@if ! command -v chta >/dev/null 2>&1; then \
		echo "   🔧 chta is not accessible. Try:"; \
		if [ -f ~/bin/chta ]; then \
			echo "      → Add ~/bin to PATH (see 'make show-path-help')"; \
		elif [ -f /usr/local/bin/chta ]; then \
			echo "      → /usr/local/bin should be in PATH by default"; \
			echo "      → Try restarting your terminal"; \
		else \
			echo "      → Install chta first with 'make auto-install'"; \
		fi; \
	else \
		echo "   ✅ chta is properly installed and accessible!"; \
	fi

# Check if /usr/local/bin is in PATH
check-path:
	@if ! echo $$PATH | grep -q "/usr/local/bin"; then \
		echo "⚠️  /usr/local/bin is not in your PATH"; \
		echo "💡 You may need to add it to your shell config"; \
	fi

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -f chta
	@echo "✅ Clean complete"

# Run tests with coverage
test:
	@echo "🧪 Running tests..."
	go test -v -cover ./...

# Run tests with race detection
test-race:
	@echo "🏁 Running tests with race detection..."
	go test -v -race ./...

# Development mode - run without building
dev:
	@echo "🔧 Running in development mode..."
	go run $(LDFLAGS) main.go $(ARGS)

# Show comprehensive help
help:
	@echo "🐆 Chta - Fast CLI Cheat Sheet Tool"
	@echo ""
	@echo "📦 Build Commands:"
	@echo "  make build          Build the binary"
	@echo "  make check-deps     Check Go installation and dependencies"
	@echo ""
	@echo "⚡ Installation Commands:"
	@echo "  make auto-install   🎯 Smart install (recommended)"
	@echo "  make install        Install globally with sudo"
	@echo "  make install-local  Install to ~/bin"
	@echo "  make install-force  Force overwrite existing installation"
	@echo "  make completion     Setup shell auto-completion"
	@echo "  make uninstall      Remove chta completely"
	@echo ""
	@echo "🔧 Troubleshooting Commands:"
	@echo "  make diagnose       Diagnose installation issues"
	@echo "  make show-path-help Show PATH setup instructions"
	@echo ""
	@echo "🧪 Development Commands:"
	@echo "  make dev ARGS='git' Run in development mode"
	@echo "  make test           Run tests"
	@echo "  make test-race      Run tests with race detection"
	@echo "  make clean          Remove build artifacts"
	@echo ""
	@echo "💡 Quick Start:"
	@echo "  make auto-install   # Install chta"
	@echo "  make diagnose       # If you have issues"
	@echo "  chta git            # Try it out!"

# Development with specific args (e.g., make dev-git)
dev-%:
	@echo "🔧 Running: chta $*"
	go run $(LDFLAGS) main.go $*

# Default target shows help
all: help 