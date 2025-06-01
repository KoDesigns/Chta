package cmd

import (
	"fmt"
	"os"

	"github.com/KoDesigns/chta/internal/display"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chta",
	Short: "🐆 Fast CLI cheat sheet tool",
	Long: `Chta (like Cheetah) is a fast CLI cheat sheet manager with interactive command execution.

Quick reference and command execution - perfect for everyone.

Examples:
  chta git                # View Git cheat sheet
  chta run git            # Execute Git commands interactively  
  chta run git --dry-run  # Preview commands safely
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
