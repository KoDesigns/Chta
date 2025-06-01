# Chta - Fast CLI Cheat Sheets 🐆

*Quick reference and command execution for everyone*

## View Cheat Sheets

```bash
chta git                    # View Git cheat sheet
chta docker                 # View Docker cheat sheet  
chta chta                   # View this cheat sheet
chta list                   # List all available sheets
```

## Execute Commands Interactively

```bash
chta run git                # Pick and run Git commands
chta run docker             # Pick and run Docker commands
chta run git --dry-run      # Preview commands safely
```

## Adding Your Own

```bash
# Create new cheat sheet
echo "# My Tool\n\n\`\`\`bash\nmytool --help\n\`\`\`" > examples/mytool.md

# Use it
chta mytool
chta run mytool
```

## How Interactive Run Works

1. **Extract** - Finds executable commands in markdown code blocks
2. **Display** - Shows numbered list of commands with descriptions  
3. **Select** - User picks a number (1, 2, 3...)
4. **Confirm** - Asks "Continue? (y/N)" for safety
5. **Execute** - Runs the command in your shell

## Tips

- **Start with built-ins**: `git`, `docker`, `chta`
- **Use dry-run first**: `--dry-run` to preview safely
- **Markdown format**: Standard code blocks work perfectly
- **Context matters**: Commands get descriptions from section headings

---

*Fast like a cheetah, helpful like a cheat sheet* 🚀 