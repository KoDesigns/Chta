package storage

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// CheatSheet represents a cheat sheet with its metadata
type CheatSheet struct {
	Name    string
	Path    string
	Content string
}

// GetCheatSheet reads a cheat sheet by name from examples directory or user storage
func GetCheatSheet(name string) (*CheatSheet, error) {
	// First check examples directory (built-in cheat sheets)
	examplesPath := filepath.Join("examples", name+".md")
	if content, err := os.ReadFile(examplesPath); err == nil {
		return &CheatSheet{
			Name:    name,
			Path:    examplesPath,
			Content: string(content),
		}, nil
	}

	// TODO: Check user storage directory
	// For now, return error if not found in examples
	return nil, fmt.Errorf("cheat sheet '%s' not found", name)
}

// ListCheatSheets returns a list of all available cheat sheets
func ListCheatSheets() ([]string, error) {
	var sheets []string

	// Check examples directory
	err := filepath.WalkDir("examples", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if strings.HasSuffix(path, ".md") {
			name := strings.TrimSuffix(filepath.Base(path), ".md")
			sheets = append(sheets, name)
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to list cheat sheets: %w", err)
	}

	// TODO: Also check user storage directory

	return sheets, nil
} 