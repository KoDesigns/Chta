# 🐆 Chta - Fast CLI Cheat Sheet Tool

[![Made with Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/KoDesigns/Chta)](https://github.com/KoDesigns/Chta/issues)
[![GitHub stars](https://img.shields.io/github/stars/KoDesigns/Chta)](https://github.com/KoDesigns/Chta/stargazers)

**Lightning-fast CLI cheat sheet manager with interactive execution and revolutionary navigation** ⚡

---

## ✨ Features

| Feature | Description |
|---------|-------------|
| 🚀 **Interactive TOC Navigation** | Revolutionary split-screen layout with table of contents |
| ⚡ **Direct Section Access** | Jump to any section with `chta git 3` |
| 🎮 **Interactive Execution** | Run commands directly from cheat sheets |
| 📱 **Smart Terminal UI** | Beautiful, responsive terminal interface |
| 🔍 **Fuzzy Search & Suggestions** | Smart search with typo tolerance |
| 🏗️ **Built-in Examples** | Comes with curated cheat sheets (Git, Docker, etc.) |
| 🛠️ **Custom Cheat Sheets** | Create and manage your own cheat sheets |
| 🌍 **Cross-platform** | Works on Linux, macOS, and Windows |
| 🎯 **Smart Installation** | Auto-detecting installation with PATH setup |

---

## 🚀 Command Syntax

### **Basic Usage**
```bash
# View cheat sheet with interactive TOC navigation
chta git

# Jump directly to section 3 (Branching)
chta git 3

# Jump to section 5 (Remote Operations)  
chta docker 5

# List all available cheat sheets
chta list
```

### **Revolutionary TOC Navigation**
When you open a cheat sheet, you get a **split-screen interface**:

```
Section Content                                    │ 📋 Table of Contents
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
## Branching                                      │   1. Basic Operations
                                                  │ ▶ 2. Branching
# List branches                                   │   3. Remote Operations
git branch                                        │   4. Viewing History
git branch -a  # Show all branches                │   5. Undoing Changes
                                                  │ ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
# Create branch                                   │ 🎮 Navigation:
git checkout -b <name>                            │   [1-5] Jump to section
                                                  │   [n]ext  [p]rev
# Switch branches                                 │   [h]elp  [q]uit
git switch <branch>                               │
```

**Navigation Controls:**
- **Numbers (1-9)**: Jump to specific section
- **n/next**: Go to next section  
- **p/prev**: Go to previous section
- **h/help**: Show help
- **q/quit**: Exit viewer

### **Interactive Execution**
```bash
# Run commands interactively with selection menu
chta run git

# Preview commands without execution (dry-run mode)
chta run git --dry-run

# Direct execution mode
chta exec docker
```

### **Management Commands**
```bash
# Initialize user cheat sheets directory
chta init

# List all available cheat sheets
chta list

# Show version and build info
chta --version
```

---

## 🎯 Quick Examples

### **⚡ Lightning Navigation**
```bash
# Traditional way (lots of scrolling)
chta git  # scroll... scroll... scroll...

# New way (instant access) 🚀
chta git 4     # Jump directly to "Remote Operations"
chta docker 2  # Jump to "Container Management"
chta git 7     # Jump to "Stashing"
```

### **🎮 Interactive Workflow**
```bash
# Open Git cheat sheet → Interactive TOC Navigation
chta git

# In the TUI:
# Press "3" → Jump to Branching section
# Press "n" → Go to next section  
# Press "5" → Jump to Remote Operations
# Press "q" → Exit
```

### **🏃‍♂️ Productivity Boost**
```bash
# Before: Search through long docs
man git | grep branch  # Hope for the best

# With Chta: Instant, structured access
chta git 3  # Directly to branching commands
# See all branch operations in organized format
```

---

## 🛠️ Installation

### **🎯 Smart Installation (Recommended)**
```bash
git clone https://github.com/KoDesigns/Chta.git
cd Chta
make auto-install
```

### **Alternative Methods**

| Method | Command | Use Case |
|--------|---------|----------|
| **Local Install** | `make install-local` | User-specific installation |
| **Global Install** | `make install` | System-wide access |
| **Force Install** | `make install-force` | Overwrite existing |

### **🔧 Shell Completion**
```bash
# Setup auto-completion for your shell
make completion
```

---

## 🎮 Interactive Navigation Tutorial

### **Step 1: Open any cheat sheet**
```bash
chta git
```

### **Step 2: Explore the split-screen layout**
- **Left side**: Table of Contents with numbered sections
- **Right side**: Section content with syntax highlighting
- **Current section**: Highlighted with ▶ arrow

### **Step 3: Navigate like a pro**
```bash
# In the interactive mode:
3        # Jump to section 3
n        # Next section
p        # Previous section  
h        # Help and tips
q        # Quit
```

### **Step 4: Direct access for speed**
```bash
# Skip the TUI - go directly to what you need
chta git 5      # Remote operations
chta docker 3   # Volume management
chta git 8      # Merge conflicts
```

---

## 📝 Creating Custom Cheat Sheets

### **Initialize your directory**
```bash
chta init
```

### **Create a new cheat sheet**
```bash
# Create markdown file
echo '# My Tool

## Installation
```bash
curl -sSL https://get.mytool.com | sh
```

## Basic Commands
```bash
mytool --help
mytool status
mytool deploy
```' > ~/.chta/sheets/mytool.md

# Use it immediately
chta mytool     # Interactive TOC view
chta mytool 2   # Jump to "Basic Commands"
```

### **Cheat Sheet Structure**
```markdown
# Tool Name

## Section 1 Name
Content here...

## Section 2 Name  
More content...

## Section N Name
Final content...
```

**💡 Pro Tip**: Use `##` (level 2) headers for main sections that appear in the TOC!

---

## 🔧 Make Commands Reference

### **📦 Build Commands**
```bash
make build          # Build the binary
make check-deps     # Check Go installation and dependencies
```

### **⚡ Installation Commands**
```bash
make auto-install   # 🎯 Smart install (recommended)
make install        # Install globally with sudo
make install-local  # Install to ~/bin
make install-force  # Force overwrite existing installation
make completion     # Setup shell auto-completion
make uninstall      # Remove chta completely
```

### **🧪 Development Commands**
```bash
make dev ARGS='git' # Run in development mode
make test           # Run tests
make test-race      # Run tests with race detection
make clean          # Remove build artifacts
```

### **🔍 Troubleshooting Commands**
```bash
make diagnose       # Installation diagnostics
make show-path-help # PATH setup guidance
```

---

## 🌍 Cross-Platform Support

### **📱 macOS**
```bash
# Installation
make auto-install

# PATH setup (if needed)
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

### **🐧 Linux**
```bash
# Installation
make auto-install

# Shell completion
make completion
```

### **🪟 Windows**
```bash
# Git Bash / PowerShell
git clone https://github.com/KoDesigns/Chta.git
cd Chta
go build -o chta.exe main.go

# Add to PATH manually or use Windows installer
```

---

## 💡 Pro Tips & Best Practices

### **🚀 Productivity Hacks**
```bash
# Create aliases for frequently used sections
alias gitbranch="chta git 3"
alias gitremote="chta git 4"  
alias dockerrun="chta docker 2"

# Quick reference without UI
chta git 5 | head -20  # First 20 lines of section 5
```

### **🎯 Speed Navigation**
- **Learn section numbers**: `chta git` → remember that branching = 3
- **Use direct access**: `chta git 3` is faster than scrolling
- **Bookmark workflows**: Know which sections you use most
- **Combine with shell history**: `!!` + section numbers

### **📚 Organization**
- **Group related commands** in the same cheat sheet
- **Use descriptive section names** (## Basic Operations vs ## Ops)
- **Keep sections focused** - one concept per section
- **Add examples** in code blocks for clarity

---

## 🐛 Troubleshooting

### **❌ Command not found after installation**
```bash
# Run diagnostics
make diagnose

# Check PATH setup
echo $PATH

# Manual PATH fix
export PATH="$HOME/bin:$PATH"  # Linux/macOS
```

### **❌ No cheat sheets available**
```bash
# Initialize user directory
chta init

# Check if cheat sheets are embedded
chta list

# Verify from any directory
cd /tmp && chta list
```

### **❌ Installation issues**
```bash
# Check dependencies
make check-deps

# Force reinstall
make clean && make install-force

# Get help
make show-path-help
```

---

## 🤝 Contributing

We welcome contributions! Here's how to get started:

1. **Fork the repository**
2. **Create a feature branch**: `git checkout -b feature/amazing-feature`
3. **Make your changes** following Go best practices
4. **Add tests** for new functionality
5. **Submit a pull request**

### **Development Setup**
```bash
git clone https://github.com/KoDesigns/Chta.git
cd Chta
go mod tidy
make dev ARGS='git'  # Test your changes
```

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 🌟 Star History

If you find Chta useful, please give it a star! ⭐

---

**Made with ❤️ for developers who value speed and productivity**

*Happy coding! 🚀* 