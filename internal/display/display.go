package display

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/KoDesigns/chta/internal/storage"
	"github.com/charmbracelet/glamour"
	"golang.org/x/term"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Config holds display configuration
type Config struct {
	Width     int
	StylePath string
	DarkMode  bool
}

// Section represents a section in the cheat sheet
type Section struct {
	Number  int
	Title   string
	Content string
	Level   int
}

// DefaultConfig returns default display configuration
func DefaultConfig() Config {
	return Config{
		Width:    80,
		DarkMode: true, // Default to dark mode for terminal
	}
}

// ParseSections extracts sections from markdown content
func ParseSections(content string) []Section {
	var sections []Section
	lines := strings.Split(content, "\n")

	currentSection := Section{Number: 0, Title: "Introduction", Level: 1}
	sectionNumber := 0
	var currentContent []string

	// Only treat ## (level 2) headers as main sections
	headerRegex := regexp.MustCompile(`^(#{2})\s+(.+)$`)

	for _, line := range lines {
		if matches := headerRegex.FindStringSubmatch(line); matches != nil {
			// Save previous section if it has content
			if len(currentContent) > 0 || currentSection.Number == 0 {
				currentSection.Content = strings.Join(currentContent, "\n")
				sections = append(sections, currentSection)
			}

			// Start new section
			sectionNumber++
			title := strings.TrimSpace(matches[2])

			currentSection = Section{
				Number: sectionNumber,
				Title:  title,
				Level:  2, // All main sections are level 2
			}
			currentContent = []string{line} // Include the header in content
		} else {
			currentContent = append(currentContent, line)
		}
	}

	// Add final section
	if len(currentContent) > 0 {
		currentSection.Content = strings.Join(currentContent, "\n")
		sections = append(sections, currentSection)
	}

	return sections
}

// RenderTOC creates a table of contents sidebar
func RenderTOC(sections []Section, selectedSection int) string {
	var toc strings.Builder

	toc.WriteString("📋 Table of Contents\n")
	toc.WriteString(strings.Repeat("━", 25) + "\n")

	for _, section := range sections {
		prefix := "  "
		if section.Number == selectedSection {
			prefix = "▶ " // Highlight selected section
		}

		// Indent based on header level
		indent := strings.Repeat("  ", section.Level-1)

		toc.WriteString(fmt.Sprintf("%s%s%d. %s\n", prefix, indent, section.Number, section.Title))
	}

	toc.WriteString(strings.Repeat("━", 25) + "\n")
	toc.WriteString("🎮 Navigation:\n")
	toc.WriteString("  [1-9] Jump to section\n")
	toc.WriteString("  [n]ext  [p]rev\n")
	toc.WriteString("  [h]elp  [q]uit\n")

	return toc.String()
}

// RenderSectionContent renders a specific section with enhanced formatting
func RenderSectionContent(section Section, config Config) (string, error) {
	// Add section header
	header := fmt.Sprintf("# Section %d: %s\n\n", section.Number, section.Title)
	content := header + section.Content

	return RenderMarkdown(content, config)
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

// ShowCheatSheetWithSection displays a cheat sheet with TOC navigation
func ShowCheatSheetWithSection(name string, sectionNumber int) error {
	sheet, err := storage.GetCheatSheet(name)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	sections := ParseSections(sheet.Content)
	if len(sections) == 0 {
		return fmt.Errorf("❌ No sections found in %s cheat sheet", name)
	}

	// Validate section number
	if sectionNumber < 1 || sectionNumber > len(sections) {
		sectionNumber = 1
	}

	return ShowInteractiveTOC(name, sections, sectionNumber)
}

// ShowInteractiveTOC displays an interactive TOC-based cheat sheet viewer
func ShowInteractiveTOC(name string, sections []Section, currentSection int) error {
	titleCaser := cases.Title(language.Und)
	config := DefaultConfig()

	// Auto-detect terminal width and split it
	termWidth := getTerminalWidth()
	if termWidth > 120 {
		config.Width = termWidth - 35 // Leave space for TOC
	}

	for {
		// Clear screen
		fmt.Print("\033[2J\033[H")

		// Header
		fmt.Printf("🐆 %s Cheat Sheet - Interactive Navigation\n", titleCaser.String(name))
		fmt.Println(strings.Repeat("═", 70))
		fmt.Println()

		// Split screen layout
		toc := RenderTOC(sections, currentSection)

		// Get current section
		var currentSectionData Section
		for _, section := range sections {
			if section.Number == currentSection {
				currentSectionData = section
				break
			}
		}

		// Render section content
		sectionContent, err := RenderSectionContent(currentSectionData, config)
		if err != nil {
			sectionContent = currentSectionData.Content // Fallback to plain text
		}

		// Display split layout
		tocLines := strings.Split(toc, "\n")
		contentLines := strings.Split(sectionContent, "\n")

		maxLines := len(tocLines)
		if len(contentLines) > maxLines {
			maxLines = len(contentLines)
		}

		for i := 0; i < maxLines; i++ {
			// TOC column (30 chars wide)
			tocLine := ""
			if i < len(tocLines) {
				tocLine = tocLines[i]
			}

			// Pad or truncate TOC line to 30 chars
			if len(tocLine) > 30 {
				tocLine = tocLine[:27] + "..."
			}
			tocLine = fmt.Sprintf("%-30s", tocLine)

			// Content column
			contentLine := ""
			if i < len(contentLines) {
				contentLine = contentLines[i]
			}

			fmt.Printf("%s │ %s\n", tocLine, contentLine)
		}

		fmt.Println()
		fmt.Printf("📍 Section %d/%d | Press number (1-%d), n/p, h for help, q to quit\n",
			currentSection, len(sections), len(sections))
		fmt.Print("⚡ Command: ")

		// Read user input
		var input string
		fmt.Scanln(&input)
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "q", "quit", "exit":
			fmt.Println("👋 Goodbye!")
			return nil

		case "h", "help", "?":
			showTOCHelp(name, len(sections))
			continue

		case "n", "next":
			if currentSection < len(sections) {
				currentSection++
			} else {
				fmt.Println("❌ Already at last section")
				fmt.Print("Press Enter to continue...")
				fmt.Scanln()
			}

		case "p", "prev", "previous":
			if currentSection > 1 {
				currentSection--
			} else {
				fmt.Println("❌ Already at first section")
				fmt.Print("Press Enter to continue...")
				fmt.Scanln()
			}

		default:
			// Try to parse as section number
			if num, err := strconv.Atoi(input); err == nil && num >= 1 && num <= len(sections) {
				currentSection = num
			} else {
				fmt.Printf("❌ Invalid input '%s'. Use 1-%d, n/p, h, or q\n", input, len(sections))
				fmt.Print("Press Enter to continue...")
				fmt.Scanln()
			}
		}
	}
}

// showTOCHelp displays help for the TOC navigation
func showTOCHelp(name string, sectionCount int) {
	fmt.Print("\033[2J\033[H") // Clear screen
	fmt.Printf("❓ %s Cheat Sheet - Navigation Help\n", cases.Title(language.Und).String(name))
	fmt.Println(strings.Repeat("═", 50))
	fmt.Println()
	fmt.Println("🎮 Navigation Commands:")
	fmt.Printf("  1-%d         Jump directly to section number\n", sectionCount)
	fmt.Println("  n, next     Go to next section")
	fmt.Println("  p, prev     Go to previous section")
	fmt.Println("  h, help, ?  Show this help")
	fmt.Println("  q, quit     Exit cheat sheet viewer")
	fmt.Println()
	fmt.Println("💡 Pro Tips:")
	fmt.Printf("  • Use 'chta %s 3' to open directly at section 3\n", name)
	fmt.Println("  • Numbers in the left TOC show available sections")
	fmt.Println("  • Current section is highlighted with ▶")
	fmt.Println("  • Content width adapts to your terminal size")
	fmt.Println()
	fmt.Println("🚀 Quick Access Examples:")
	fmt.Printf("  chta %s 1    # Open at introduction\n", name)
	fmt.Printf("  chta %s 2    # Jump to section 2\n", name)
	fmt.Printf("  chta %s      # Start from beginning\n", name)
	fmt.Println()
	fmt.Print("Press Enter to continue...")
	fmt.Scanln()
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

// getTerminalWidth returns the terminal width, defaulting to 80 if unable to detect
func getTerminalWidth() int {
	if width, _, err := term.GetSize(int(os.Stdout.Fd())); err == nil && width > 0 {
		return width
	}
	return 80 // Default fallback
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
	// Show welcome with list of available cheat sheets
	sheets, err := storage.ListCheatSheets()
	if err != nil {
		return fmt.Errorf("failed to list cheat sheets: %w", err)
	}

	fmt.Println("🐆 Welcome to Chta - Fast CLI Cheat Sheet Tool")
	fmt.Println()

	if len(sheets) == 0 {
		fmt.Println("📋 No cheat sheets found")
		fmt.Println("💡 Try: chta init  # to create user directory")
		return nil
	}

	fmt.Println("📋 Available cheat sheets:")
	for _, sheet := range sheets {
		fmt.Printf("  • %s\n", sheet)
	}

	fmt.Println()
	fmt.Println("💡 Usage:")
	fmt.Println("  chta <name>           # View cheat sheet")
	fmt.Println("  chta run <name>       # Interactive execution")
	fmt.Println("  chta run <name> -i    # Fuzzy search mode")
	fmt.Println("  chta list             # List all sheets")
	fmt.Println("  chta init             # Setup user directory")

	return nil
}

// GetAvailableSheets returns list of available cheat sheets for shell completion
func GetAvailableSheets() ([]string, error) {
	return storage.ListCheatSheets()
}
