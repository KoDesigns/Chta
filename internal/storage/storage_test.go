package storage

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetUserSheetsDir(t *testing.T) {
	userDir, err := getUserSheetsDir()
	if err != nil {
		t.Fatalf("getUserSheetsDir() failed: %v", err)
	}

	if userDir == "" {
		t.Error("getUserSheetsDir() returned empty string")
	}

	// Should end with .chta/sheets
	expectedSuffix := filepath.Join(".chta", "sheets")
	if !filepath.IsAbs(userDir) {
		t.Error("getUserSheetsDir() should return absolute path")
	}

	if !strings.HasSuffix(userDir, expectedSuffix) {
		t.Errorf("getUserSheetsDir() = %q, should end with %q", userDir, expectedSuffix)
	}
}

func TestListCheatSheetsEmptyDirectories(t *testing.T) {
	// Test with no examples directory and no user directory
	originalWd, _ := os.Getwd()

	// Create temporary directory and change to it
	tempDir := t.TempDir()
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalWd)

	sheets, err := ListCheatSheets()
	if err != nil {
		t.Errorf("ListCheatSheets() failed: %v", err)
	}

	if len(sheets) != 0 {
		t.Errorf("Expected 0 sheets in empty directories, got %d", len(sheets))
	}
}

func TestListCheatSheetsWithExamples(t *testing.T) {
	originalWd, _ := os.Getwd()

	// Create temporary directory with examples
	tempDir := t.TempDir()
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalWd)

	// Create examples directory with test files
	err = os.MkdirAll("examples", 0755)
	if err != nil {
		t.Fatalf("Failed to create examples directory: %v", err)
	}

	testFiles := []string{"git.md", "docker.md", "test.md"}
	for _, file := range testFiles {
		path := filepath.Join("examples", file)
		err = os.WriteFile(path, []byte("# Test\n\n```bash\ntest command\n```"), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file %s: %v", file, err)
		}
	}

	sheets, err := ListCheatSheets()
	if err != nil {
		t.Errorf("ListCheatSheets() failed: %v", err)
	}

	expectedSheets := []string{"git", "docker", "test"}
	if len(sheets) != len(expectedSheets) {
		t.Errorf("Expected %d sheets, got %d", len(expectedSheets), len(sheets))
	}

	// Check that all expected sheets are present (order may vary)
	found := make(map[string]bool)
	for _, sheet := range sheets {
		found[sheet] = true
	}

	for _, expected := range expectedSheets {
		if !found[expected] {
			t.Errorf("Expected sheet %q not found in results", expected)
		}
	}
}

func TestGetCheatSheetNotFound(t *testing.T) {
	originalWd, _ := os.Getwd()

	// Create temporary empty directory
	tempDir := t.TempDir()
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalWd)

	_, err = GetCheatSheet("nonexistent")
	if err == nil {
		t.Error("GetCheatSheet() should fail for nonexistent sheet")
	}

	expectedError := "cheat sheet 'nonexistent' not found"
	if err.Error() != expectedError {
		t.Errorf("Expected error %q, got %q", expectedError, err.Error())
	}
}

func TestCreateUserSheetsDir(t *testing.T) {
	// This test will create actual directories, so be careful
	// We'll use the real user directory but clean up after
	userDir, err := getUserSheetsDir()
	if err != nil {
		t.Fatalf("getUserSheetsDir() failed: %v", err)
	}

	// Remove if exists (cleanup from previous tests)
	os.RemoveAll(userDir)

	err = CreateUserSheetsDir()
	if err != nil {
		t.Errorf("CreateUserSheetsDir() failed: %v", err)
	}

	// Check that directory was created
	if _, err := os.Stat(userDir); os.IsNotExist(err) {
		t.Error("CreateUserSheetsDir() did not create the directory")
	}

	// Cleanup
	os.RemoveAll(filepath.Dir(userDir)) // Remove .chta directory
}

func TestFindSimilarCheatSheets(t *testing.T) {
	// Mock the ListCheatSheets function for testing
	originalWd, _ := os.Getwd()

	// Create temporary directory with examples
	tempDir := t.TempDir()
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	defer os.Chdir(originalWd)

	// Create examples directory with test files
	err = os.MkdirAll("examples", 0755)
	if err != nil {
		t.Fatalf("Failed to create examples directory: %v", err)
	}

	testFiles := []string{"git.md", "docker.md", "kubernetes.md", "terraform.md"}
	for _, file := range testFiles {
		path := filepath.Join("examples", file)
		err = os.WriteFile(path, []byte("# Test\n\n```bash\ntest command\n```"), 0644)
		if err != nil {
			t.Fatalf("Failed to write test file %s: %v", file, err)
		}
	}

	testCases := []struct {
		input    string
		expected []string
		desc     string
	}{
		{"gi", []string{"git"}, "prefix match"},
		{"dock", []string{"docker"}, "partial match"},
		{"gti", []string{"git"}, "typo (levenshtein distance 1)"},
		{"kuber", []string{"kubernetes"}, "partial match"},
		{"xyz", []string{}, "no matches"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := findSimilarCheatSheets(tc.input)

			if len(result) != len(tc.expected) {
				t.Errorf("Expected %d suggestions, got %d", len(tc.expected), len(result))
				return
			}

			for i, expected := range tc.expected {
				if i >= len(result) || result[i] != expected {
					t.Errorf("Expected suggestion %d to be %q, got %q", i, expected, result[i])
				}
			}
		})
	}
}

func TestLevenshteinDistance(t *testing.T) {
	testCases := []struct {
		a, b     string
		expected int
		desc     string
	}{
		{"", "", 0, "empty strings"},
		{"git", "git", 0, "identical strings"},
		{"git", "gti", 2, "transposition (actually 2 operations)"},
		{"docker", "doker", 1, "one deletion"},
		{"git", "gitt", 1, "one insertion"},
		{"abc", "xyz", 3, "complete replacement"},
		{"kubernetes", "kuber", 5, "multiple operations"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := levenshteinDistance(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("levenshteinDistance(%q, %q) = %d, expected %d", tc.a, tc.b, result, tc.expected)
			}
		})
	}
}
