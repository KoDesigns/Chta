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

### 🎯 Smart Install (Recommended)
```bash
git clone https://github.com/KoDesigns/chta.git && cd chta
make auto-install       # Automatically chooses best install method
```

### 📦 Installation Options

#### Option 1: System-wide (requires sudo)
```bash
make install            # Install to /usr/local/bin (Linux/macOS)
make install-force      # Force overwrite existing installation
```

#### Option 2: User directory (no sudo needed)
```bash
make install-local      # Install to ~/bin
```

**⚠️ PATH Setup for User Install:**
If you see "command not found" after user install, add `~/bin` to your PATH:

```bash
# For bash users (Linux/macOS)
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc
source ~/.bashrc

# For zsh users (macOS default)
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc

# For fish users
fish_add_path ~/bin
```

#### Option 3: Manual Build
```bash
go build -o chta main.go
# Then move to desired location:
sudo mv chta /usr/local/bin/        # Global
# or
mv chta ~/bin/                      # User (ensure ~/bin is in PATH)
```

### 🐚 Shell Completion (Optional)
```bash
make completion         # Setup auto-completion for all shells
# or manual setup:
chta completion bash > ~/.local/share/bash-completion/completions/chta
chta completion zsh > ~/.zsh/completions/_chta
chta completion fish > ~/.config/fish/completions/chta.fish
```

### 🗑️ Uninstall
```bash
make uninstall          # Removes chta and completions completely
```

### 🔧 Troubleshooting

**Can't find chta command after install:**
```bash
# Check if chta is installed
which chta
ls -la ~/bin/chta       # For user install
ls -la /usr/local/bin/chta  # For system install

# Check PATH
echo $PATH | grep -o ~/bin      # Should show ~/bin for user install
echo $PATH | grep -o /usr/local/bin  # Should show /usr/local/bin for system install

# Fix PATH (choose your shell)
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.bashrc && source ~/.bashrc    # Bash
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc && source ~/.zshrc      # Zsh
```

**Permission denied:**
```bash
# Make sure binary is executable
chmod +x ~/bin/chta
# or
chmod +x /usr/local/bin/chta
```

### 🪟 Windows Support

**Option 1: Git Bash/WSL (Recommended)**
```bash
git clone https://github.com/KoDesigns/chta.git
cd chta
go build -o chta.exe main.go
# Move to a directory in PATH, e.g.:
mv chta.exe /c/Windows/System32/   # (requires admin)
# or add current directory to PATH
```

**Option 2: PowerShell**
```powershell
git clone https://github.com/KoDesigns/chta.git
cd chta
go build -o chta.exe main.go
# Add to PATH or move to Windows directory
```

**Windows PATH Setup:**
1. Search "Environment Variables" in Start Menu
2. Click "Environment Variables" 
3. Under "User Variables", select "Path" → Edit
4. Add the directory containing `chta.exe`
5. Restart terminal

### 🔍 Verify Installation
```bash
chta --version          # Check version
chta --help            # Show help
chta git               # Test with built-in cheat sheet
make help              # Show all available make commands
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