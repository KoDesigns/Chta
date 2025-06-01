package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/KoDesigns/chta/internal/storage"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run [tool-name]",
	Short: "Interactively run commands from a cheat sheet",
	Long: `Run commands from a cheat sheet interactively.

This command parses the cheat sheet, extracts all executable commands,
displays them as a numbered list, and lets you select which one to run.

Examples:
  chta run git               # Show Git commands to run interactively
  chta run docker            # Show Docker commands to run interactively
  chta run git --dry-run     # Show commands without executing`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		toolName := args[0]
		dryRun, _ := cmd.Flags().GetBool("dry-run")

		if err := runInteractiveCommands(toolName, dryRun); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolP("dry-run", "d", false, "Show commands without executing them")
}

// runInteractiveCommands displays commands from a cheat sheet and lets user select one to run
func runInteractiveCommands(toolName string, dryRun bool) error {
	// Get the cheat sheet
	sheet, err := storage.GetCheatSheet(toolName)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	// Extract commands from the markdown
	commands := extractCommands(sheet.Content)
	if len(commands) == 0 {
		return fmt.Errorf("❌ No executable commands found in %s cheat sheet", toolName)
	}

	// Display the commands with numbers
	fmt.Printf("🐆 Interactive %s Commands\n", strings.Title(toolName))
	fmt.Println(strings.Repeat("─", 50))
	fmt.Println()

	for i, cmd := range commands {
		fmt.Printf("%2d. %s\n", i+1, cmd.Description)
		fmt.Printf("    💻 %s\n", cmd.Command)
		fmt.Println()
	}

	if dryRun {
		fmt.Println("🔍 Dry run mode - commands shown but not executed")
		return nil
	}

	// Get user selection
	fmt.Print("Enter command number to run (or 'q' to quit): ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	input = strings.TrimSpace(input)
	if input == "q" || input == "quit" {
		fmt.Println("👋 Goodbye!")
		return nil
	}

	// Parse selection
	selection, err := strconv.Atoi(input)
	if err != nil || selection < 1 || selection > len(commands) {
		return fmt.Errorf("❌ Invalid selection. Please enter a number between 1 and %d", len(commands))
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

	// Look for common command patterns
	commonCommands := []string{"git", "docker", "chta", "npm", "yarn", "go", "curl", "wget", "ssh", "scp", "ls", "cd", "mkdir", "rm", "cp", "mv"}

	for _, cmd := range commonCommands {
		if strings.HasPrefix(line, cmd+" ") || line == cmd {
			return true
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

// executeCommand runs a shell command
func executeCommand(command string) error {
	// Use shell to execute the command so pipes and other shell features work
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}
