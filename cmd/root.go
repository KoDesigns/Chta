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
	// Allow unknown arguments (cheat sheet names)
	Args: cobra.ArbitraryArgs,
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// Auto-complete cheat sheet names
		if len(args) == 0 {
			sheets, err := getAvailableSheets()
			if err != nil {
				return nil, cobra.ShellCompDirectiveError
			}
			return sheets, cobra.ShellCompDirectiveNoFileComp
		}
		return nil, cobra.ShellCompDirectiveNoFileComp
	},
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

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:
  $ source <(chta completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ chta completion bash > /etc/bash_completion.d/chta
  # macOS:
  $ chta completion bash > /usr/local/etc/bash_completion.d/chta

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ chta completion zsh > "${fpath[1]}/_chta"

  # You will need to start a new shell for this setup to take effect.

Fish:
  $ chta completion fish | source

  # To load completions for each session, execute once:
  $ chta completion fish > ~/.config/fish/completions/chta.fish

PowerShell:
  PS> chta completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> chta completion powershell > chta.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

// getAvailableSheets returns list of available cheat sheets for completion
func getAvailableSheets() ([]string, error) {
	// Import storage here to avoid import cycle
	sheets, err := display.GetAvailableSheets()
	if err != nil {
		return nil, err
	}
	return sheets, nil
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

	// Add completion command
	rootCmd.AddCommand(completionCmd)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.chta.yaml)")
	rootCmd.Flags().BoolP("version", "v", false, "Show version information")

	// Enable shell completion for subcommands
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "help [command]",
		Short:  "Help about any command",
		Hidden: true,
	})
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
