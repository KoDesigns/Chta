# Chta - CLI Cheat Sheet Tool 🐆

*Fast like a cheetah, helpful like a cheat sheet*

## Basic Usage

```bash
# Show this cheat sheet
chta chta

# Show interactive browser (coming soon)
chta

# Show specific cheat sheet
chta git
chta docker
chta fzf
```

## Managing Cheat Sheets

```bash
# List all available cheat sheets
chta list

# Add a new cheat sheet
chta add <tool-name>

# Edit existing cheat sheet
chta edit <tool-name>

# Remove a cheat sheet
chta remove <tool-name>
```

## Remote Cheat Sheets

```bash
# Fetch cheat sheet from URL
chta fetch <url> <tool-name>

# Publish your cheat sheet (coming soon)
chta publish <tool-name>

# Search community cheat sheets (coming soon)
chta search <keyword>
```

## Configuration

```bash
# Show current configuration
chta config

# Set default editor
chta config set editor vim

# Set storage location
chta config set storage ~/.local/share/chta
```

## Tips & Tricks

- **Self-documenting**: This cheat sheet documents Chta itself!
- **Markdown format**: All cheat sheets use standard Markdown
- **Built-in examples**: Start with `git`, `docker`, `fzf` examples
- **Fast search**: Use fuzzy finding to quickly locate commands
- **Shareable**: Export and share your cheat sheets with team

## File Locations

```bash
# Cheat sheets stored in:
~/.local/share/chta/sheets/

# Configuration file:
~/.chta.yaml

# Built-in examples:
$(chta --install-dir)/examples/
```

## Pro Tips

1. **Start small**: Begin with commands you use daily
2. **Use examples**: Include real command examples with explanations  
3. **Group logically**: Organize commands by functionality
4. **Keep it current**: Update cheat sheets as you learn new tricks
5. **Share the love**: Contribute useful cheat sheets back to the community

---

*Made with ❤️ for developers who are always learning new tools* 