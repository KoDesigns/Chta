# Chta 🐆

> Fast like a cheetah, helpful like a cheat sheet

Chta is a lightning-fast CLI cheat sheet manager for developers who are always learning new tools. Keep quick references to Git, Docker, FZF, and any other CLI tools at your fingertips.

## Features

- **Self-documenting**: The primary cheat sheet documents Chta itself
- **Built-in examples**: Comes with Git, Docker, and FZF cheat sheets
- **Fast access**: Quick lookup with `chta <tool-name>`
- **Markdown format**: All cheat sheets use standard Markdown
- **Extensible**: Add your own cheat sheets easily
- **Open source**: MIT licensed and community-driven

## Installation

### From Source

```bash
# Clone the repository
git clone https://github.com/KoDesigns/chta.git
cd chta

# Build and install
go build -o chta main.go
sudo mv chta /usr/local/bin/
```

### Quick Test

```bash
# In the project directory
go run main.go
```

## Usage

### Basic Commands

```bash
# Show welcome message and available cheat sheets
chta

# View a specific cheat sheet
chta git
chta docker
chta chta          # Self-reference!

# List all available cheat sheets
chta list

# Get help
chta --help
```

### Available Cheat Sheets

- `chta chta` - How to use Chta itself
- `chta git` - Essential Git commands
- More coming soon!

## Project Structure

```
chta/
├── cmd/                    # CLI command definitions
│   ├── root.go            # Root command and configuration
│   └── list.go            # List command
├── internal/              # Private application code
│   ├── storage/           # Cheat sheet file management
│   └── display/           # Output formatting
├── examples/              # Built-in cheat sheets
│   ├── chta.md           # Self-reference
│   └── git.md            # Git commands
├── main.go               # Entry point
├── go.mod                # Go module definition
└── .cursorrules          # Development guidelines
```

## Development

### Prerequisites

- Go 1.21 or later
- Basic understanding of CLI development

### Getting Started

```bash
# Clone and enter directory
git clone https://github.com/KoDesigns/chta.git
cd chta

# Install dependencies
go mod tidy

# Run locally
go run main.go

# Build binary
go build -o chta main.go
```

### Adding New Cheat Sheets

1. Create a new `.md` file in the `examples/` directory
2. Follow the existing format with clear sections and examples
3. Test with `go run main.go <your-sheet-name>`

### Code Guidelines

This project follows Go best practices defined in `.cursorrules`:

- Use `gofmt` and `goimports` for formatting
- Write tests for core functionality
- Keep functions small and focused
- Handle errors explicitly
- Use meaningful variable names

## Roadmap

- [ ] User storage directory support
- [ ] Remote cheat sheet fetching
- [ ] Interactive fuzzy search
- [ ] Syntax highlighting with Glamour
- [ ] Configuration management
- [ ] Add/edit/remove commands
- [ ] Community cheat sheet sharing

## Contributing

We welcome contributions! Here's how to get started:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following our coding standards
4. Add tests for your changes
5. Commit with conventional commits (`feat:`, `fix:`, `docs:`)
6. Push and create a Pull Request

### Contribution Ideas

- Add new cheat sheets for popular CLI tools
- Improve the display formatting
- Add search functionality
- Enhance error handling
- Write documentation

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for quick CLI reference
- Built with ❤️ using [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper)
- Named after the cheetah for speed and "cheat" for cheat sheets

---

**Made with ❤️ for developers who are always learning new tools**

[⭐ Star this project](https://github.com/KoDesigns/chta) if you find it useful! 