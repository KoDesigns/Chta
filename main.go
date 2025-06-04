package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/KoDesigns/chta/cmd"
	"github.com/KoDesigns/chta/internal/storage"
)

//go:embed examples/*.md
var embeddedCheatSheets embed.FS

// Version information (set by build flags)
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	// Set the embedded filesystem for storage package
	storage.EmbeddedFS = embeddedCheatSheets

	// Handle version flag before cobra takes over
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("chta version %s\n", Version)
		fmt.Printf("Built: %s\n", BuildTime)
		os.Exit(0)
	}

	cmd.Execute()
}
