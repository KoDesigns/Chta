# Chta 🐆

> Fast CLI cheat sheet manager with interactive command execution

**Chta** (like Cheetah) helps anyone quickly reference and execute commands from markdown cheat sheets. Perfect for everyone.

## Quick Start

```bash
# Clone and build
git clone https://github.com/KoDesigns/chta.git
cd chta

# Option 1: Use Makefile (recommended)
make install              # Install globally
# or
make install-user         # Install in ~/bin

# Option 2: Manual build
go build -o chta main.go
sudo mv chta /usr/local/bin/

# Use built-in cheat sheets
chta git                  # View Git commands
chta run git              # Execute Git commands interactively
chta list                 # See all available cheat sheets
```

## Features

- **📖 View** - `chta git` shows the full Git cheat sheet
- **🚀 Execute** - `chta run git` lets you pick and run commands directly  
- **📝 Markdown** - All cheat sheets are simple `.md` files
- **⚡ Fast** - Instant lookup and execution
- **🔧 Extensible** - Add your own cheat sheets easily

## Core Commands

```bash
chta <tool>              # View cheat sheet
chta run <tool>          # Interactive command execution
chta run <tool> --dry-run # Preview commands safely
chta list                # List available cheat sheets
```

## Built-in Cheat Sheets

- **`chta git`** - Essential Git commands
- **`chta docker`** - Docker commands and examples  
- **`chta chta`** - How to use Chta itself

## Adding Your Own

Create `.md` files in the `examples/` directory:

```markdown
# My Tool Cheat Sheet

## Basic Commands
```bash
mytool init
mytool build --prod
mytool deploy
```
```

Run with: `chta mytool` or `chta run mytool`

## How It Works

1. **Markdown parsing** - Extracts commands from code blocks
2. **Smart filtering** - Only executable commands, skips placeholders  
3. **Interactive selection** - Pick commands by number
4. **Safe execution** - Confirmation before running

## Development

```bash
go run main.go git          # Test locally
go build -o chta main.go    # Build binary
```

**Requirements:** Go 1.21+

## License

Apache 2.0

---

**Made for everyone who values speed and simplicity** 🚀 