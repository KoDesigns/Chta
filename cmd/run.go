package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/KoDesigns/chta/internal/storage"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [tool-name]",
	Short: "Interactively run commands from a cheat sheet",
	Long: `Run commands from a cheat sheet interactively.

This command parses the cheat sheet, extracts all executable commands,
displays them as a numbered list, and lets you select which one to run.

Commands are displayed in pages of 10. Use n/next and p/prev to navigate.

Examples:
  chta run git                     # Show Git commands to run interactively
  chta run docker                  # Show Docker commands to run interactively
  chta run git --dry-run           # Show commands without executing
  chta run git --search commit     # Filter commands containing "commit"
  chta run git -s push             # Short form of search
  chta run git --interactive       # Real-time fuzzy search with arrow keys
  chta run git -i                  # Short form of interactive search`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		dryRun, _ := cmd.Flags().GetBool("dry-run")
		searchTerm, _ := cmd.Flags().GetString("search")
		interactive, _ := cmd.Flags().GetBool("interactive")

		if interactive {
			if err := runInteractiveSearch(toolName, dryRun); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		} else {
			if err := runInteractiveCommands(toolName, dryRun, searchTerm); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("dry-run", "d", false, "Show commands without executing them")
	runCmd.Flags().StringP("search", "s", "", "Filter commands by keyword")
	runCmd.Flags().BoolP("interactive", "i", false, "Interactive fuzzy search mode")
}

// runInteractiveCommands displays commands from a cheat sheet and lets user select one to run
func runInteractiveCommands(toolName string, dryRun bool, searchTerm string) error {
	// Get the cheat sheet
	sheet, err := storage.GetCheatSheet(toolName)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	// Extract commands from the markdown
	allCommands := extractCommands(sheet.Content)
	if len(allCommands) == 0 {
		return fmt.Errorf("❌ No executable commands found in %s cheat sheet", toolName)
	}

	// Apply search filter if provided
	commands := allCommands
	if searchTerm != "" {
		commands = filterCommandsBySearch(allCommands, searchTerm)
		if len(commands) == 0 {
			return fmt.Errorf("❌ No commands found matching '%s' in %s cheat sheet", searchTerm, toolName)
		}
	}

	// Display the commands with numbers (with pagination)
	titleCaser := cases.Title(language.Und)
	fmt.Printf("🐆 Interactive %s Commands\n", titleCaser.String(toolName))
	fmt.Println(strings.Repeat("─", 50))
	fmt.Println()

	if dryRun {
		// In dry-run mode, show all commands without pagination
		for i, cmd := range commands {
			fmt.Printf("%2d. %s\n", i+1, cmd.Description)
			fmt.Printf("    💻 %s\n", cmd.Command)
			fmt.Println()
		}
		fmt.Println("🔍 Dry run mode - commands shown but not executed")
		return nil
	}

	// Interactive mode with pagination
	const pageSize = 10
	totalPages := (len(commands) + pageSize - 1) / pageSize
	currentPage := 0

	for {
		// Display current page
		start := currentPage * pageSize
		end := start + pageSize
		if end > len(commands) {
			end = len(commands)
		}

		fmt.Printf("📄 Page %d of %d (%d-%d of %d commands)\n\n",
			currentPage+1, totalPages, start+1, end, len(commands))

		for i := start; i < end; i++ {
			fmt.Printf("%2d. %s\n", i+1, commands[i].Description)
			fmt.Printf("    💻 %s\n", commands[i].Command)
			fmt.Println()
		}

		// Show navigation options
		fmt.Println("Navigation:")
		if currentPage > 0 {
			fmt.Println("  p/prev - Previous page")
		}
		if currentPage < totalPages-1 {
			fmt.Println("  n/next - Next page")
		}
		fmt.Println("  1-" + strconv.Itoa(len(commands)) + " - Select command number")
		fmt.Println("  q/quit - Quit")
		fmt.Print("\nEnter selection: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read input: %w", err)
		}

		input = strings.TrimSpace(strings.ToLower(input))

		// Handle navigation
		switch input {
		case "q", "quit":
			fmt.Println("👋 Goodbye!")
			return nil
		case "n", "next":
			if currentPage < totalPages-1 {
				currentPage++
				fmt.Print("\033[2J\033[H") // Clear screen
				fmt.Printf("🐆 Interactive %s Commands\n", titleCaser.String(toolName))
				fmt.Println(strings.Repeat("─", 50))
				fmt.Println()
				continue
			} else {
				fmt.Println("❌ Already at last page")
				continue
			}
		case "p", "prev":
			if currentPage > 0 {
				currentPage--
				fmt.Print("\033[2J\033[H") // Clear screen
				fmt.Printf("🐆 Interactive %s Commands\n", titleCaser.String(toolName))
				fmt.Println(strings.Repeat("─", 50))
				fmt.Println()
				continue
			} else {
				fmt.Println("❌ Already at first page")
				continue
			}
		}

		// Try to parse as command number
		selection, err := strconv.Atoi(input)
		if err != nil || selection < 1 || selection > len(commands) {
			fmt.Printf("❌ Invalid selection. Enter 1-%d, n/p for navigation, or q to quit\n\n", len(commands))
			continue
		}

		selectedCmd := commands[selection-1]

		// Confirm execution for safety
		fmt.Printf("🔄 About to run: %s\n", selectedCmd.Command)
		fmt.Print("Continue? (y/N): ")

		confirm, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read confirmation: %w", err)
		}

		confirm = strings.ToLower(strings.TrimSpace(confirm))
		if confirm != "y" && confirm != "yes" {
			fmt.Println("❌ Command cancelled")
			return nil
		}

		// Execute the command
		fmt.Printf("🚀 Executing: %s\n", selectedCmd.Command)
		fmt.Println(strings.Repeat("─", 50))

		return executeCommand(selectedCmd.Command)
	}
}

// runInteractiveSearch provides a real-time fuzzy search interface
func runInteractiveSearch(toolName string, dryRun bool) error {
	// Get the cheat sheet
	sheet, err := storage.GetCheatSheet(toolName)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	// Extract commands from the markdown
	allCommands := extractCommands(sheet.Content)
	if len(allCommands) == 0 {
		return fmt.Errorf("❌ No executable commands found in %s cheat sheet", toolName)
	}

	titleCaser := cases.Title(language.Und)
	fmt.Printf("🔍 Interactive Search for %s Commands\n", titleCaser.String(toolName))
	fmt.Println("Type to filter, ↑↓ to navigate, Enter to select, Esc to quit")
	fmt.Println(strings.Repeat("─", 60))

	// Interactive search state
	searchQuery := ""
	selectedIndex := 0
	filteredCommands := allCommands

	// Save original terminal state
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return fmt.Errorf("failed to set raw mode: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		// Clear screen and show interface
		fmt.Print("\033[2J\033[H") // Clear screen

		fmt.Printf("🔍 Interactive Search for %s Commands\n", titleCaser.String(toolName))
		fmt.Println("Type to filter, ↑↓ to navigate, Enter to select, Esc to quit")
		fmt.Println(strings.Repeat("─", 60))
		fmt.Printf("Search: %s_\n\n", searchQuery)

		// Filter commands based on search query
		filteredCommands = fuzzyFilterCommands(allCommands, searchQuery)

		if len(filteredCommands) == 0 {
			fmt.Printf("❌ No commands match '%s'\n", searchQuery)
		} else {
			// Ensure selected index is valid
			if selectedIndex >= len(filteredCommands) {
				selectedIndex = len(filteredCommands) - 1
			}
			if selectedIndex < 0 {
				selectedIndex = 0
			}

			// Display filtered commands with selection highlight
			const maxDisplay = 10
			displayCount := len(filteredCommands)
			if displayCount > maxDisplay {
				displayCount = maxDisplay
			}

			for i := 0; i < displayCount; i++ {
				cmd := filteredCommands[i]
				prefix := "  "
				if i == selectedIndex {
					prefix = "▶ " // Highlight selected
				}

				fmt.Printf("%s%d. %s\n", prefix, i+1, cmd.Description)
				fmt.Printf("    💻 %s\n", cmd.Command)
				if i < displayCount-1 {
					fmt.Println()
				}
			}

			if len(filteredCommands) > maxDisplay {
				fmt.Printf("\n... and %d more commands\n", len(filteredCommands)-maxDisplay)
			}
		}

		// Read single character
		buf := make([]byte, 1)
		n, err := os.Stdin.Read(buf)
		if err != nil || n == 0 {
			continue
		}

		char := buf[0]

		switch char {
		case 27: // ESC key
			fmt.Print("\033[2J\033[H") // Clear screen
			fmt.Println("👋 Search cancelled!")
			return nil

		case 13: // Enter key
			if len(filteredCommands) > 0 && selectedIndex < len(filteredCommands) {
				selectedCmd := filteredCommands[selectedIndex]

				// Restore terminal
				term.Restore(int(os.Stdin.Fd()), oldState)

				fmt.Print("\033[2J\033[H") // Clear screen
				fmt.Printf("🚀 Selected: %s\n", selectedCmd.Description)
				fmt.Printf("Command: %s\n", selectedCmd.Command)

				if dryRun {
					fmt.Println("🔍 Dry run mode - command not executed")
					return nil
				}

				// Confirm execution
				fmt.Print("Execute this command? (y/N): ")
				reader := bufio.NewReader(os.Stdin)
				confirm, err := reader.ReadString('\n')
				if err != nil {
					return fmt.Errorf("failed to read confirmation: %w", err)
				}

				confirm = strings.ToLower(strings.TrimSpace(confirm))
				if confirm == "y" || confirm == "yes" {
					return executeCommand(selectedCmd.Command)
				} else {
					fmt.Println("❌ Command cancelled")
					return nil
				}
			}

		case 65: // Up arrow (actually part of escape sequence, simplified)
			if selectedIndex > 0 {
				selectedIndex--
			}

		case 66: // Down arrow (simplified)
			if selectedIndex < len(filteredCommands)-1 {
				selectedIndex++
			}

		case 127, 8: // Backspace/Delete
			if len(searchQuery) > 0 {
				searchQuery = searchQuery[:len(searchQuery)-1]
				selectedIndex = 0 // Reset selection
			}

		default:
			// Regular character - add to search query
			if char >= 32 && char <= 126 { // Printable ASCII
				searchQuery += string(char)
				selectedIndex = 0 // Reset selection
			}
		}
	}
}

// Command represents an executable command with description
type Command struct {
	Command     string
	Description string
}

// extractCommands parses markdown content and extracts executable commands
func extractCommands(content string) []Command {
	var commands []Command
	lines := strings.Split(content, "\n")

	var inCodeBlock bool
	var currentBlockCommands []string
	var lastHeading string

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Track headings for context
		if strings.HasPrefix(line, "#") {
			lastHeading = strings.TrimSpace(strings.TrimLeft(line, "#"))
			continue
		}

		// Track code blocks
		if strings.HasPrefix(line, "```") {
			if inCodeBlock {
				// End of code block - process collected commands
				for _, cmd := range currentBlockCommands {
					if isExecutableCommand(cmd) {
						description := fmt.Sprintf("%s: %s", lastHeading, getCommandDescription(cmd))
						commands = append(commands, Command{
							Command:     cmd,
							Description: description,
						})
					}
				}
				currentBlockCommands = nil
			} else {
				// Start of code block
				currentBlockCommands = nil
			}
			inCodeBlock = !inCodeBlock
			continue
		}

		// Collect commands in code blocks
		if inCodeBlock && line != "" && !strings.HasPrefix(line, "#") {
			currentBlockCommands = append(currentBlockCommands, line)
		}
	}

	return commands
}

// isExecutableCommand determines if a line is an executable command
func isExecutableCommand(line string) bool {
	line = strings.TrimSpace(line)

	// Skip empty lines and comments
	if line == "" || strings.HasPrefix(line, "#") {
		return false
	}

	// Skip lines that are clearly examples or placeholders
	if strings.Contains(line, "<") && strings.Contains(line, ">") {
		return false
	}

	// Skip lines that look like output or descriptions
	if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") {
		return false
	}

	// Skip variable assignments (improved detection)
	if strings.Contains(line, "=") && !strings.Contains(line, " == ") && !strings.Contains(line, " != ") {
		// Allow command line args like --flag=value
		if !strings.Contains(line, "--") && !strings.Contains(line, "-") {
			return false
		}
	}

	// Skip lines that look like file paths or URLs
	if strings.HasPrefix(line, "/") && strings.Count(line, "/") > 2 {
		return false
	}
	if strings.HasPrefix(line, "http://") || strings.HasPrefix(line, "https://") {
		return false
	}

	// Look for command patterns
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return false
	}

	firstWord := parts[0]

	// Skip obvious non-commands (expanded list)
	nonCommands := []string{
		"echo", "export", "set", "source", ".", "alias",
		"printf", "read", "eval", "exec", "which", "type",
		"unset", "shift", "return", "exit", "break", "continue",
	}
	for _, nonCmd := range nonCommands {
		if firstWord == nonCmd {
			return false
		}
	}

	// Skip shell built-ins that are typically not useful in cheat sheets
	if firstWord == "cd" && len(parts) == 1 {
		return false // cd without arguments
	}

	// Accept well-known command patterns
	knownCommands := []string{
		"git", "docker", "kubectl", "terraform", "ansible", "vagrant",
		"npm", "yarn", "go", "python", "pip", "node", "java", "mvn",
		"make", "cmake", "cargo", "rustc", "gcc", "clang",
		"curl", "wget", "ssh", "scp", "rsync", "tar", "zip", "unzip",
		"aws", "gcloud", "az", "heroku", "firebase",
		"systemctl", "service", "crontab", "ps", "top", "htop",
		"find", "grep", "awk", "sed", "sort", "head", "tail",
	}

	for _, knownCmd := range knownCommands {
		if firstWord == knownCmd {
			return true
		}
	}

	// Accept if it looks like a command (contains executable characters)
	// Most commands are lowercase and may contain hyphens, underscores
	if len(firstWord) > 0 {
		// Basic heuristic: if it looks like a command name
		hasLetters := false
		for _, r := range firstWord {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				hasLetters = true
				break
			}
		}

		// If it has letters and doesn't start with special chars that indicate non-commands
		if hasLetters && !strings.HasPrefix(firstWord, "$") && !strings.HasPrefix(firstWord, "@") {
			// Additional check: if it looks like a typical command
			// Commands usually don't contain spaces in the first word
			if !strings.Contains(firstWord, " ") {
				return true
			}
		}
	}

	return false
}

// getCommandDescription creates a short description for a command
func getCommandDescription(command string) string {
	// Try to extract description from comments
	if idx := strings.Index(command, "#"); idx != -1 {
		desc := strings.TrimSpace(command[idx+1:])
		if desc != "" {
			return desc
		}
	}

	// Generate description based on command
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return "Command"
	}

	switch parts[0] {
	case "git":
		if len(parts) > 1 {
			return fmt.Sprintf("Git %s", parts[1])
		}
		return "Git command"
	case "docker":
		if len(parts) > 1 {
			return fmt.Sprintf("Docker %s", parts[1])
		}
		return "Docker command"
	default:
		return fmt.Sprintf("%s command", parts[0])
	}
}

// executeCommand runs a shell command cross-platform
func executeCommand(command string) error {
	var cmd *exec.Cmd

	// Use appropriate shell based on operating system
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/C", command)
	default: // linux, darwin (macOS), etc.
		cmd = exec.Command("/bin/sh", "-c", command)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

// filterCommandsBySearch filters commands based on a search term
func filterCommandsBySearch(commands []Command, searchTerm string) []Command {
	var filteredCommands []Command
	for _, cmd := range commands {
		if strings.Contains(cmd.Command, searchTerm) || strings.Contains(cmd.Description, searchTerm) {
			filteredCommands = append(filteredCommands, cmd)
		}
	}
	return filteredCommands
}

// fuzzyFilterCommands filters commands using fuzzy matching logic
func fuzzyFilterCommands(commands []Command, query string) []Command {
	if query == "" {
		return commands
	}

	var filteredCommands []Command
	query = strings.ToLower(query)

	for _, cmd := range commands {
		// Check if query matches command or description (case-insensitive)
		cmdLower := strings.ToLower(cmd.Command)
		descLower := strings.ToLower(cmd.Description)

		if strings.Contains(cmdLower, query) || strings.Contains(descLower, query) {
			filteredCommands = append(filteredCommands, cmd)
		}
	}

	return filteredCommands
}
