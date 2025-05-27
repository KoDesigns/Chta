package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/KoDesigns/chta/internal/display"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available cheat sheets",
	Long: `List all available cheat sheets from both built-in examples 
and user-created cheat sheets.

Examples:
  chta list               # Show all available cheat sheets`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := display.ListCheatSheets(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
} 