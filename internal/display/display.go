package display

import (
	"fmt"
	"os"
	"strings"

	"github.com/KoDesigns/chta/internal/storage"
	"github.com/charmbracelet/glamour"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Config holds display configuration
type Config struct {
	Width     int
	StylePath string
	DarkMode  bool
}

// DefaultConfig returns default display configuration
func DefaultConfig() Config {
	return Config{
		Width:    80,
		DarkMode: true, // Default to dark mode for terminal
	}
}

// RenderMarkdown renders markdown content with syntax highlighting
func RenderMarkdown(content string, config Config) (string, error) {
	// Configure glamour renderer
	var style string
	if config.DarkMode {
		style = "dark"
	} else {
		style = "light"
	}

	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(config.Width),
		glamour.WithStylePath(style),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create renderer: %w", err)
	}

	return r.Render(content)
}

// PrintCheatSheet renders and prints a cheat sheet with beautiful formatting
func PrintCheatSheet(name, content string) error {
	config := DefaultConfig()

	// Auto-detect terminal width
	if width := getTerminalWidth(); width > 0 {
		config.Width = width
	}

	rendered, err := RenderMarkdown(content, config)
	if err != nil {
		// Fallback to plain text if rendering fails
		fmt.Printf("🐆 %s Cheat Sheet\n", cases.Title(language.Und).String(name))
		fmt.Println(strings.Repeat("─", 50))
		fmt.Println(content)
		return nil
	}

	fmt.Print(rendered)
	return nil
}

// getTerminalWidth returns the terminal width or 0 if unknown
func getTerminalWidth() int {
	// Try to get terminal width from environment or terminal
	if width, ok := os.LookupEnv("COLUMNS"); ok {
		if w := parseInt(width); w > 0 {
			return w
		}
	}

	// Default to 80 if we can't determine
	return 80
}

// parseInt safely parses an integer, returning 0 on error
func parseInt(s string) int {
	result := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return result
}

// RenderCommandList renders a numbered list of commands with highlighting
func RenderCommandList(commands []Command, title string) (string, error) {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# 🐆 %s\n\n", title))

	for i, cmd := range commands {
		sb.WriteString(fmt.Sprintf("## %d. %s\n\n", i+1, cmd.Description))
		sb.WriteString(fmt.Sprintf("```bash\n%s\n```\n\n", cmd.Command))
	}

	config := DefaultConfig()
	return RenderMarkdown(sb.String(), config)
}

// Command represents an executable command (duplicated here to avoid import cycles)
type Command struct {
	Command     string
	Description string
}

// ShowCheatSheet displays a cheat sheet with beautiful formatting
func ShowCheatSheet(name string) error {
	sheet, err := storage.GetCheatSheet(name)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	// Use glamour rendering for beautiful output
	return PrintCheatSheet(sheet.Name, sheet.Content)
}

// ListCheatSheets displays all available cheat sheets
func ListCheatSheets() error {
	sheets, err := storage.ListCheatSheets()
	if err != nil {
		return fmt.Errorf("❌ Failed to list cheat sheets: %w", err)
	}

	if len(sheets) == 0 {
		fmt.Println("📋 No cheat sheets found")
		return nil
	}

	fmt.Println("📋 Available cheat sheets:")
	for _, sheet := range sheets {
		fmt.Printf("  • %s\n", sheet)
	}

	fmt.Printf("\n💡 Use 'chta <name>' to view a cheat sheet\n")
	return nil
}

// ShowWelcome displays the welcome message with available cheat sheets
func ShowWelcome() error {
	fmt.Println("🐆 Welcome to Chta - Fast CLI Cheat Sheet Tool")
	fmt.Println()

	sheets, err := storage.ListCheatSheets()
	if err == nil && len(sheets) > 0 {
		fmt.Println("Available cheat sheets:")
		for _, sheet := range sheets {
			fmt.Printf("  chta %s\n", sheet)
		}
		fmt.Println()
	}

	fmt.Println("Run 'chta --help' for more commands")
	fmt.Println("Run 'chta chta' to see how to use Chta itself!")

	return nil
}
