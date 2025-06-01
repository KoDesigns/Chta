# Chta 🐆

> Fast CLI cheat sheet manager with interactive command execution and beautiful markdown rendering

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
chta git                  # View Git commands with beautiful formatting
chta run git              # Execute Git commands interactively
chta run git -i           # Real-time fuzzy search with arrow keys
chta list                 # See all available cheat sheets
```

## Features

- **📖 Beautiful Display** - `chta git` shows cheat sheets with syntax highlighting
- **🚀 Interactive Execution** - `chta run git` lets you pick and run commands directly  
- **🔍 Fuzzy Search** - `chta run git -i` provides real-time interactive search
- **📝 Markdown Support** - All cheat sheets are simple `.md` files with glamour rendering
- **⚡ Fast** - Instant lookup and execution
- **🔧 Extensible** - Add your own cheat sheets easily
- **🌍 Cross-platform** - Works on Linux, macOS, and Windows

## Core Commands

```bash
chta <tool>                      # View cheat sheet with beautiful formatting
chta run <tool>                  # Interactive command execution with pagination
chta run <tool> --interactive    # Real-time fuzzy search mode
chta run <tool> -i               # Short form of interactive search
chta run <tool> --search commit  # Filter commands by keyword
chta run <tool> --dry-run        # Preview commands safely
chta list                        # List available cheat sheets
chta init                        # Initialize user cheat sheets directory
```

## Interactive Modes

### 📄 **Paginated Mode** (default)
```bash
chta run git                     # Browse commands with pagination
# Navigate: n/next, p/prev, 1-N to select, q to quit
```

### 🔍 **Interactive Search Mode** 
```bash
chta run git --interactive       # Real-time fuzzy search
# Type to filter, ↑↓ to navigate, Enter to select, Esc to quit
```

### 🎯 **Keyword Search Mode**
```bash
chta run git --search commit     # Filter commands containing "commit"
chta run git -s push             # Filter commands containing "push"
```

## Built-in Cheat Sheets

- **`chta git`** - Essential Git commands
- **`chta docker`** - Docker commands and examples  
- **`chta chta`** - How to use Chta itself

## Adding Your Own

**Step 1: Initialize your user directory**
```bash
chta init                 # Creates ~/.chta/sheets/ directory
```

**Step 2: Add your cheat sheets**
```bash
# Create a new cheat sheet file
echo "# My Tool

## Basic Commands
\`\`\`bash
mytool --help
mytool init
mytool deploy --prod
\`\`\`" > ~/.chta/sheets/mytool.md

# Use it immediately
chta mytool               # View your cheat sheet
chta run mytool           # Run commands interactively
```

**Alternative: Add to examples/ directory**
```bash
# For built-in cheat sheets (in project directory)
echo "# My Tool

\`\`\`bash
mytool --help
\`\`\`" > examples/mytool.md
```

**Auto-detection:**
- User cheat sheets: `~/.chta/sheets/*.md` (recommended)
- Built-in cheat sheets: `examples/*.md` 
- User cheat sheets override built-ins with the same name

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
**Platforms:** Linux, macOS, Windows

## License

Apache 2.0

---

**Made for everyone who values speed and simplicity** 🚀 