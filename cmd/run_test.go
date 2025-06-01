package cmd

import (
	"testing"
)

func TestIsExecutableCommand(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
		desc     string
	}{
		// Should detect as commands
		{"git status", true, "basic git command"},
		{"docker run", true, "basic docker command"},
		{"kubectl get pods", true, "kubernetes command"},
		{"terraform plan", true, "terraform command"},
		{"make build", true, "make command"},
		{"npm install", true, "npm command"},
		{"yarn start", true, "yarn command"},
		{"go build", true, "go command"},
		{"python script.py", true, "python command"},
		{"node app.js", true, "node command"},

		// Should NOT detect as commands
		{"", false, "empty string"},
		{"# This is a comment", false, "comment"},
		{"// This is a comment", false, "C-style comment"},
		{"/* Block comment */", false, "block comment"},
		{"<url>", false, "placeholder"},
		{"git clone <repository>", false, "command with placeholder"},
		{"VARIABLE=value", false, "variable assignment"},
		{"echo 'hello'", false, "echo command (excluded)"},
		{"export PATH=/usr/bin", false, "export command (excluded)"},
		{"$HOME/bin/command", false, "variable reference"},
		{"@deprecated", false, "annotation"},

		// Edge cases
		{"git", true, "single command word"},
		{"ls -la", true, "command with flags"},
		{"cd /home/user", true, "cd command"},
		{"   git status   ", true, "command with whitespace"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := isExecutableCommand(tc.input)
			if result != tc.expected {
				t.Errorf("isExecutableCommand(%q) = %v, expected %v", tc.input, result, tc.expected)
			}
		})
	}
}

func TestGetCommandDescription(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
		desc     string
	}{
		{"git status", "Git status", "git command with subcommand"},
		{"docker run", "Docker run", "docker command with subcommand"},
		{"kubectl get pods", "kubectl command", "non-git/docker command"},
		{"git status # Check repository status", "Check repository status", "command with inline comment"},
		{"make build # Build the project", "Build the project", "command with comment"},
		{"unknown", "unknown command", "unknown command"},
		{"", "Command", "empty command"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := getCommandDescription(tc.input)
			if result != tc.expected {
				t.Errorf("getCommandDescription(%q) = %q, expected %q", tc.input, result, tc.expected)
			}
		})
	}
}

func TestExtractCommands(t *testing.T) {
	testMarkdown := `# Test Cheat Sheet

Some intro text.

## Basic Commands

` + "```bash" + `
git status
git add .
git commit -m "message"
` + "```" + `

## Advanced Commands

` + "```bash" + `
# This is a comment
git push origin main
docker run -it ubuntu  # Run ubuntu container
<placeholder-command>
export PATH=/usr/bin
kubectl get pods
` + "```"

	commands := extractCommands(testMarkdown)

	expectedCommands := []string{
		"git status",
		"git add .",
		"git commit -m \"message\"",
		"git push origin main",
		"docker run -it ubuntu  # Run ubuntu container",
		"kubectl get pods",
	}

	if len(commands) != len(expectedCommands) {
		t.Errorf("Expected %d commands, got %d", len(expectedCommands), len(commands))
		t.Logf("Got commands: %+v", commands)
		return
	}

	for i, expected := range expectedCommands {
		if commands[i].Command != expected {
			t.Errorf("Command %d: expected %q, got %q", i, expected, commands[i].Command)
		}
	}
}

func TestFuzzyFilterCommands(t *testing.T) {
	commands := []Command{
		{Command: "git add .", Description: "Add all files"},
		{Command: "git commit -m", Description: "Create a commit"},
		{Command: "git push origin", Description: "Push to remote"},
		{Command: "docker build", Description: "Build Docker image"},
		{Command: "docker run -d", Description: "Run container in background"},
	}

	testCases := []struct {
		query    string
		expected int
		desc     string
	}{
		{"", 5, "empty query returns all commands"},
		{"git", 3, "filter by command prefix"},
		{"commit", 1, "filter by description"},
		{"docker", 2, "filter by docker commands"},
		{"xyz", 0, "no matches"},
		{"ADD", 1, "case-insensitive matching"},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := fuzzyFilterCommands(commands, tc.query)
			if len(result) != tc.expected {
				t.Errorf("Expected %d results for query '%s', got %d", tc.expected, tc.query, len(result))
			}
		})
	}
}
