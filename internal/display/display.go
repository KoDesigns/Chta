package display

import (
	"fmt"
	"strings"

	"github.com/KoDesigns/chta/internal/storage"
)

// ShowCheatSheet displays a cheat sheet with basic formatting
func ShowCheatSheet(name string) error {
	sheet, err := storage.GetCheatSheet(name)
	if err != nil {
		return fmt.Errorf("❌ %w", err)
	}

	// Simple terminal output for now
	fmt.Printf("📋 %s\n", sheet.Name)
	fmt.Println(strings.Repeat("─", 50))
	fmt.Println(sheet.Content)

	return nil
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