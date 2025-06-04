# 🐆 Chta
> **Fast CLI cheat sheet manager** - View, search, and execute commands interactively

[![Go](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-Apache%202.0-green.svg)](LICENSE)
[![Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20macOS%20%7C%20Windows-lightgrey.svg)]()

---

## 🚀 Quick Start

```bash
# Install
git clone https://github.com/KoDesigns/chta.git && cd chta
make install

# Try it out
chta git                    # 📖 View Git cheat sheet
chta run git --interactive  # 🔍 Interactive command execution
```

---

## 📋 Command Syntax

### Basic Usage
```bash
chta <tool>                 # View cheat sheet with syntax highlighting
chta list                   # List all available cheat sheets
chta init                   # Create user cheat sheets directory
```

### Interactive Execution
```bash
chta run <tool>                     # Interactive command selection
chta run <tool> --interactive       # Real-time fuzzy search
chta run <tool> -i                  # Short form of interactive
chta run <tool> --search <keyword>  # Filter by keyword
chta run <tool> -s <keyword>        # Short form of search
chta run <tool> --dry-run           # Preview without execution
```

### Examples
```bash
# View cheat sheets
chta git                    # Display Git commands with beautiful formatting
chta docker                 # Show Docker commands and examples
chta chta                   # Learn how to use Chta itself

# Interactive execution modes
chta run git                # Navigate with n/p, select with 1-9, quit with q
chta run git -i             # Type to search, ↑↓ to navigate, Enter to select
chta run git -s commit      # Show only commands containing "commit"
chta run git --dry-run      # See what commands would run without executing
```

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 📖 **Beautiful Display** | Syntax-highlighted markdown rendering |
| 🚀 **Interactive Execution** | Pick and run commands directly |
| 🔍 **Fuzzy Search** | Real-time filtering with arrow key navigation |
| 📝 **Markdown Support** | Simple `.md` files with code block extraction |
| ⚡ **Lightning Fast** | Instant lookup and execution |
| 🔧 **Extensible** | Add custom cheat sheets easily |
| 🌍 **Cross-platform** | Linux, macOS, and Windows support |

---

## 📚 Built-in Cheat Sheets

```bash
chta git        # Git version control commands
chta docker     # Docker containerization commands  
chta chta       # Chta usage and examples
```

---

## 🛠️ Create Custom Cheat Sheets

### 1. Initialize user directory
```bash
chta init                           # Creates ~/.chta/sheets/
```

### 2. Create a cheat sheet
```bash
# Create file: ~/.chta/sheets/kubernetes.md
cat > ~/.chta/sheets/kubernetes.md << 'EOF'
# Kubernetes Commands

## Pod Management
```bash
kubectl get pods
kubectl describe pod <pod-name>
kubectl logs <pod-name>
kubectl exec -it <pod-name> -- /bin/bash
```

## Service Management
```bash
kubectl get services
kubectl expose deployment <deployment> --port=80 --target-port=8080
```
EOF
```

### 3. Use immediately
```bash
chta kubernetes             # View your cheat sheet
chta run kubernetes -i      # Interactive execution
```

### Cheat Sheet Locations
- **User sheets**: `~/.chta/sheets/*.md` (recommended)
- **Built-in sheets**: `examples/*.md` (project directory)
- User sheets override built-ins with same name

---

## 💻 Installation

### Option 1: Makefile (Recommended)
```bash
git clone https://github.com/KoDesigns/chta.git
cd chta
make install        # System-wide (/usr/local/bin)
# or
make install-user   # User directory (~/bin)
```

### Option 2: Manual Build  
```bash
go build -o chta main.go
sudo mv chta /usr/local/bin/
```

### Option 3: Development
```bash
go run main.go git  # Test locally
```

---

## 🎯 Interactive Modes

### Default Mode
```
$ chta run git
Git Commands - Interactive Mode
[n]ext [p]rev [1-9] to select [q]uit

1. git status
2. git add .
3. git commit -m "message"
...

Select command (1-9): 
```

### Fuzzy Search Mode
```
$ chta run git -i
🔍 Search: com_
> git commit -m "message"
  git commit --amend
  git commit --no-verify
↑↓ navigate | Enter select | Esc quit
```

---

## 🔧 Requirements

- **Go**: 1.21 or higher
- **OS**: Linux, macOS, Windows
- **Terminal**: Any modern terminal with color support

---

## 📄 License

[Apache 2.0](LICENSE)

---

<div align="center">

**Made with ❤️ for developers who value speed and simplicity**

[⭐ Star on GitHub](https://github.com/KoDesigns/chta) | [🐛 Report Issues](https://github.com/KoDesigns/chta/issues) | [💡 Request Features](https://github.com/KoDesigns/chta/issues/new)

</div> 