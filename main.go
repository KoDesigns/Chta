package main

import (
	"fmt"
	"os"

	"github.com/KoDesigns/chta/cmd"
)

// Version information (set by build flags)
var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	// Handle version flag before cobra takes over
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("chta version %s\n", Version)
		fmt.Printf("Built: %s\n", BuildTime)
		os.Exit(0)
	}

	cmd.Execute()
}
