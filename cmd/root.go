package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/KoDesigns/chta/internal/display"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chta",
	Short: "🐆 Fast CLI cheat sheet tool",
	Long: `Chta (like Cheetah) is a lightning-fast CLI cheat sheet manager.

Easily create, manage, and share command cheat sheets for tools you're learning.
Perfect for keeping quick references to Git, Docker, FZF, and any other CLI tools.

Examples:
  chta                    # Show interactive cheat sheet browser
  chta git                # Show Git cheat sheet
  chta add docker         # Add a new Docker cheat sheet
  chta list               # List all available cheat sheets`,
	// Disable the default completion command
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	// Allow unknown arguments (cheat sheet names)
	Args: cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments, show welcome with available cheat sheets
		if len(args) == 0 {
			if err := display.ShowWelcome(); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		} else {
			// Show specific cheat sheet
			if err := display.ShowCheatSheet(args[0]); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chta.yaml)")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".chta")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
} 