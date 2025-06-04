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

// EmbeddedFS holds the embedded cheat sheets (set by main package)
var EmbeddedFS fs.FS

// getUserSheetsDir returns the user's cheat sheets directory
func getUserSheetsDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".chta", "sheets"), nil
}

// GetCheatSheet reads a cheat sheet by name from user storage or embedded examples
func GetCheatSheet(name string) (*CheatSheet, error) {
	// First check user storage directory
	if userDir, err := getUserSheetsDir(); err == nil {
		userPath := filepath.Join(userDir, name+".md")
		if content, err := os.ReadFile(userPath); err == nil {
			return &CheatSheet{
				Name:    name,
				Path:    userPath,
				Content: string(content),
			}, nil
		}
	}

	// Then check embedded examples
	if EmbeddedFS != nil {
		embeddedPath := filepath.Join("examples", name+".md")
		if content, err := fs.ReadFile(EmbeddedFS, embeddedPath); err == nil {
			return &CheatSheet{
				Name:    name,
				Path:    embeddedPath,
				Content: string(content),
			}, nil
		}
	}

	// Fallback: check local examples directory (for development)
	examplesPath := filepath.Join("examples", name+".md")
	if _, err := os.Stat("examples"); err == nil {
		if content, err := os.ReadFile(examplesPath); err == nil {
			return &CheatSheet{
				Name:    name,
				Path:    examplesPath,
				Content: string(content),
			}, nil
		}
	}

	// If exact match not found, try fuzzy matching
	if suggestions := findSimilarCheatSheets(name); len(suggestions) > 0 {
		if len(suggestions) == 1 {
			return nil, fmt.Errorf("cheat sheet '%s' not found. Did you mean '%s'?", name, suggestions[0])
		}
		return nil, fmt.Errorf("cheat sheet '%s' not found. Did you mean: %s?", name, strings.Join(suggestions, ", "))
	}

	return nil, fmt.Errorf("cheat sheet '%s' not found", name)
}

// findSimilarCheatSheets returns cheat sheet names similar to the input using fuzzy matching
func findSimilarCheatSheets(name string) []string {
	allSheets, err := ListCheatSheets()
	if err != nil {
		return nil
	}

	var suggestions []string
	name = strings.ToLower(name)

	for _, sheet := range allSheets {
		sheetLower := strings.ToLower(sheet)

		// Exact substring match (highest priority)
		if strings.Contains(sheetLower, name) || strings.Contains(name, sheetLower) {
			suggestions = append(suggestions, sheet)
			continue
		}

		// Prefix match
		if strings.HasPrefix(sheetLower, name) || strings.HasPrefix(name, sheetLower) {
			suggestions = append(suggestions, sheet)
			continue
		}

		// Levenshtein distance for typos (distance <= 2)
		if levenshteinDistance(name, sheetLower) <= 2 {
			suggestions = append(suggestions, sheet)
		}
	}

	// Limit to 3 suggestions to avoid overwhelming user
	if len(suggestions) > 3 {
		suggestions = suggestions[:3]
	}

	return suggestions
}

// levenshteinDistance calculates the edit distance between two strings
func levenshteinDistance(a, b string) int {
	if len(a) == 0 {
		return len(b)
	}
	if len(b) == 0 {
		return len(a)
	}

	matrix := make([][]int, len(a)+1)
	for i := range matrix {
		matrix[i] = make([]int, len(b)+1)
		matrix[i][0] = i
	}
	for j := range matrix[0] {
		matrix[0][j] = j
	}

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = min(
					matrix[i-1][j]+1,   // deletion
					matrix[i][j-1]+1,   // insertion
					matrix[i-1][j-1]+1, // substitution
				)
			}
		}
	}

	return matrix[len(a)][len(b)]
}

// min returns the minimum of three integers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// ListCheatSheets returns a list of all available cheat sheets
func ListCheatSheets() ([]string, error) {
	var sheets []string
	seenSheets := make(map[string]bool)

	// First check user storage directory (user sheets override built-ins)
	if userDir, err := getUserSheetsDir(); err == nil {
		if _, err := os.Stat(userDir); err == nil {
			err := filepath.WalkDir(userDir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if d.IsDir() {
					return nil
				}

				if strings.HasSuffix(path, ".md") {
					name := strings.TrimSuffix(filepath.Base(path), ".md")
					if !seenSheets[name] {
						sheets = append(sheets, name)
						seenSheets[name] = true
					}
				}

				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("failed to list user cheat sheets: %w", err)
			}
		}
	}

	// Then check embedded examples
	if EmbeddedFS != nil {
		err := fs.WalkDir(EmbeddedFS, "examples", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, ".md") {
				name := strings.TrimSuffix(filepath.Base(path), ".md")
				if !seenSheets[name] {
					sheets = append(sheets, name)
					seenSheets[name] = true
				}
			}

			return nil
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list embedded cheat sheets: %w", err)
		}
	}

	// Fallback: check local examples directory (for development)
	if _, err := os.Stat("examples"); err == nil {
		err := filepath.WalkDir("examples", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if d.IsDir() {
				return nil
			}

			if strings.HasSuffix(path, ".md") {
				name := strings.TrimSuffix(filepath.Base(path), ".md")
				if !seenSheets[name] {
					sheets = append(sheets, name)
					seenSheets[name] = true
				}
			}

			return nil
		})

		if err != nil {
			return nil, fmt.Errorf("failed to list cheat sheets: %w", err)
		}
	}
	// If examples directory doesn't exist, that's fine - just continue with user sheets

	return sheets, nil
}

// CreateUserSheetsDir creates the user sheets directory if it doesn't exist
func CreateUserSheetsDir() error {
	userDir, err := getUserSheetsDir()
	if err != nil {
		return err
	}
	return os.MkdirAll(userDir, 0755)
}
