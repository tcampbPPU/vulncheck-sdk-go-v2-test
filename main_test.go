package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// TestMainIntegration tests the main function integration with CLI
func TestMainIntegration(t *testing.T) {
	// Build the binary first
	cmd := exec.Command("go", "build", "-o", "test-binary", ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build binary: %v", err)
	}
	defer os.Remove("test-binary") // Clean up

	tests := []struct {
		name         string
		args         []string
		expectedText []string
		shouldError  bool
	}{
		{
			name:         "help command",
			args:         []string{"help"},
			expectedText: []string{"SDK Test CLI", "Usage:"},
			shouldError:  false,
		},
		{
			name:         "list command",
			args:         []string{"list"},
			expectedText: []string{"Available test functions:", "backup", "index-vulnrichment"},
			shouldError:  false,
		},
		{
			name:         "invalid command",
			args:         []string{"nonexistent-function"},
			expectedText: []string{"Error: Function 'nonexistent-function' not found"},
			shouldError:  false, // The program handles this gracefully
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("./test-binary", tt.args...)
			output, err := cmd.CombinedOutput()

			if tt.shouldError && err == nil {
				t.Errorf("Expected command to fail, but it succeeded")
			}

			if !tt.shouldError && err != nil {
				// For graceful error handling (like invalid function names),
				// we don't expect the process to exit with an error
				if !strings.Contains(string(output), "Error: Function") {
					t.Errorf("Unexpected error: %v, output: %s", err, string(output))
				}
			}

			outputStr := string(output)
			for _, expected := range tt.expectedText {
				if !strings.Contains(outputStr, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't. Output: %s", expected, outputStr)
				}
			}
		})
	}
}

// TestBuildSuccess ensures the project builds successfully
func TestBuildSuccess(t *testing.T) {
	cmd := exec.Command("go", "build", ".")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Build failed: %v\nOutput: %s", err, string(output))
	}
}
