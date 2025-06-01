package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KoDesigns/chta/internal/storage"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize user cheat sheets directory",
	Long: `Initialize your personal cheat sheets directory at ~/.chta/sheets/

This creates the directory structure where you can add your own markdown cheat sheets.
After running this, you can add .md files to ~/.chta/sheets/ and they'll be auto-detected.

Examples:
  chta init                           # Create ~/.chta/sheets/ directory
  echo "# My Tool" > ~/.chta/sheets/mytool.md  # Add your own cheat sheet`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := initUserDirectory(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

// initUserDirectory creates the user cheat sheets directory and shows helpful info
func initUserDirectory() error {
	// Create the directory
	if err := storage.CreateUserSheetsDir(); err != nil {
		return fmt.Errorf("failed to create user directory: %w", err)
	}

	// Get the directory path to show user
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	sheetsDir := filepath.Join(homeDir, ".chta", "sheets")

	fmt.Printf("🎉 User cheat sheets directory created!\n\n")
	fmt.Printf("📁 Directory: %s\n\n", sheetsDir)
	fmt.Printf("📝 How to add your own cheat sheets:\n")
	fmt.Printf("   1. Create a .md file in the directory above\n")
	fmt.Printf("   2. Add commands in markdown code blocks\n")
	fmt.Printf("   3. Use with: chta <filename-without-extension>\n\n")
	fmt.Printf("📋 Example:\n")
	fmt.Printf("   echo '# My Tool\\n\\n```bash\\nmytool --help\\n```' > %s/mytool.md\n", sheetsDir)
	fmt.Printf("   chta mytool          # View your cheat sheet\n")
	fmt.Printf("   chta run mytool      # Run commands interactively\n\n")
	fmt.Printf("💡 Your cheat sheets will override built-in ones with the same name\n")

	return nil
}
