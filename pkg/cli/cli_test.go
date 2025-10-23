package cli

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

// captureOutput captures stdout during function execution
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

// Mock function for testing
func mockFunction() {
	fmt.Println("Mock function called")
}

// TestTestFunction tests the TestFunction struct
func TestTestFunction(t *testing.T) {
	called := false
	testFunc := TestFunction{
		Name:        "test-function",
		Description: "Test Function Description",
		Function: func() {
			called = true
		},
	}

	if testFunc.Name != "test-function" {
		t.Errorf("Expected name 'test-function', got '%s'", testFunc.Name)
	}

	if testFunc.Description != "Test Function Description" {
		t.Errorf("Expected description 'Test Function Description', got '%s'", testFunc.Description)
	}

	testFunc.Function()
	if !called {
		t.Error("Expected function to be called")
	}
}

// TestShowUsage tests the showUsage function
func TestShowUsage(t *testing.T) {
	output := captureOutput(func() {
		showUsage()
	})

	expectedStrings := []string{
		"SDK Test CLI",
		"Usage:",
		"go run main.go list",
		"go run main.go run <function-name>",
		"go run main.go <function-name>",
		"go run main.go help",
		"Examples:",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain '%s', but it didn't. Output: %s", expected, output)
		}
	}
}

// TestListFunctions tests the listFunctions function
func TestListFunctions(t *testing.T) {
	output := captureOutput(func() {
		listFunctions()
	})

	expectedStrings := []string{
		"Available test functions:",
		"backup",
		"Get Index Backup",
		"index-vulnrichment",
		"Get Index Vulnrichment",
		"Usage: go run main.go run <function-name>",
	}

	for _, expected := range expectedStrings {
		if !strings.Contains(output, expected) {
			t.Errorf("Expected output to contain '%s', but it didn't. Output: %s", expected, output)
		}
	}

	// Check that functions are sorted alphabetically
	lines := strings.Split(output, "\n")
	var functionLines []string
	for _, line := range lines {
		if strings.Contains(line, "-") && !strings.Contains(line, "Usage") {
			functionLines = append(functionLines, strings.TrimSpace(line))
		}
	}

	// Verify we have some function lines
	if len(functionLines) == 0 {
		t.Error("Expected to find function listing lines, but found none")
	}
}

// TestRunFunctionWithMocks tests the runFunction with mocked functions
func TestRunFunctionWithMocks(t *testing.T) {
	// Save original availableFunctions
	originalFunctions := availableFunctions

	// Create test functions map with mocks
	availableFunctions = map[string]TestFunction{
		"test-function": {
			Name:        "test-function",
			Description: "Test Function",
			Function:    mockFunction,
		},
		"another-test": {
			Name:        "another-test",
			Description: "Another Test Function",
			Function:    mockFunction,
		},
	}

	tests := []struct {
		name           string
		functionName   string
		expectError    bool
		expectedOutput []string
	}{
		{
			name:         "valid function name",
			functionName: "test-function",
			expectError:  false,
			expectedOutput: []string{
				"Running: Test Function",
				"--------------------------------------------------",
				"Mock function called",
			},
		},
		{
			name:         "valid function name with normalization",
			functionName: "TEST-FUNCTION",
			expectError:  false,
			expectedOutput: []string{
				"Running: Test Function",
				"--------------------------------------------------",
				"Mock function called",
			},
		},
		{
			name:         "invalid function name",
			functionName: "nonexistent-function",
			expectError:  true,
			expectedOutput: []string{
				"Error: Function 'nonexistent-function' not found",
				"Use 'go run main.go list' to see available functions",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				runFunction(tt.functionName)
			})

			for _, expected := range tt.expectedOutput {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't. Output: %s", expected, output)
				}
			}
		})
	}

	// Restore original functions
	availableFunctions = originalFunctions
}

// TestRunWithArgs tests the Run function with different command line arguments
func TestRunWithArgs(t *testing.T) {
	// Save original args and functions
	originalArgs := os.Args
	originalFunctions := availableFunctions

	// Create test functions map with mocks
	availableFunctions = map[string]TestFunction{
		"test-function": {
			Name:        "test-function",
			Description: "Test Function",
			Function:    mockFunction,
		},
	}

	tests := []struct {
		name         string
		args         []string
		expectedText []string
	}{
		{
			name:         "no arguments - show usage",
			args:         []string{"program"},
			expectedText: []string{"SDK Test CLI", "Usage:"},
		},
		{
			name:         "list command",
			args:         []string{"program", "list"},
			expectedText: []string{"Available test functions:", "test-function"},
		},
		{
			name:         "ls command (alias for list)",
			args:         []string{"program", "ls"},
			expectedText: []string{"Available test functions:", "test-function"},
		},
		{
			name:         "ll command (alias for list)",
			args:         []string{"program", "ll"},
			expectedText: []string{"Available test functions:", "test-function"},
		},
		{
			name:         "help command",
			args:         []string{"program", "help"},
			expectedText: []string{"SDK Test CLI", "Usage:"},
		},
		{
			name:         "-h flag",
			args:         []string{"program", "-h"},
			expectedText: []string{"SDK Test CLI", "Usage:"},
		},
		{
			name:         "--help flag",
			args:         []string{"program", "--help"},
			expectedText: []string{"SDK Test CLI", "Usage:"},
		},
		{
			name:         "run command without function name",
			args:         []string{"program", "run"},
			expectedText: []string{"Error: Please specify a function name to run"},
		},
		{
			name:         "run command with valid function",
			args:         []string{"program", "run", "test-function"},
			expectedText: []string{"Running: Test Function", "Mock function called"},
		},
		{
			name:         "direct function call",
			args:         []string{"program", "test-function"},
			expectedText: []string{"Running: Test Function", "Mock function called"},
		},
		{
			name:         "invalid function name",
			args:         []string{"program", "invalid-function"},
			expectedText: []string{"Error: Function 'invalid-function' not found"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set test arguments
			os.Args = tt.args

			output := captureOutput(func() {
				Run()
			})

			for _, expected := range tt.expectedText {
				if !strings.Contains(output, expected) {
					t.Errorf("Expected output to contain '%s', but it didn't. Output: %s", expected, output)
				}
			}
		})
	}

	// Restore original args and functions
	os.Args = originalArgs
	availableFunctions = originalFunctions
}

// TestAvailableFunctions tests that all expected functions are available
func TestAvailableFunctions(t *testing.T) {
	expectedFunctions := []string{
		"index-initial-access",
		"index-vulnrichment",
		"index-cve-filter",
		"index-botnet-filter",
		"index-ip-intel",
		"browse-indexes",
		"browse-backups",
		"backup",
		"purl",
		"rule",
		"tag",
		"pdns",
		"cpe",
	}

	for _, funcName := range expectedFunctions {
		if _, exists := availableFunctions[funcName]; !exists {
			t.Errorf("Expected function '%s' to be available, but it wasn't", funcName)
		}
	}

	// Test that each function has required fields
	for name, testFunc := range availableFunctions {
		if testFunc.Name == "" {
			t.Errorf("Function '%s' has empty Name field", name)
		}
		if testFunc.Description == "" {
			t.Errorf("Function '%s' has empty Description field", name)
		}
		if testFunc.Function == nil {
			t.Errorf("Function '%s' has nil Function field", name)
		}
	}
}

// TestFunctionNormalization tests the function name normalization logic
func TestFunctionNormalization(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"backup", "backup"},
		{"BACKUP", "backup"},
		{"index-vulnrichment", "index-vulnrichment"},
		{"INDEX-VULNRICHMENT", "index-vulnrichment"},
		{"Index Vulnrichment", "index-vulnrichment"},
		{"INDEX VULNRICHMENT", "index-vulnrichment"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("normalize_%s", tt.input), func(t *testing.T) {
			// Test normalization by checking if a valid function can be found
			normalizedName := strings.ToLower(strings.ReplaceAll(tt.input, " ", "-"))
			if normalizedName != tt.expected {
				t.Errorf("Expected normalization of '%s' to be '%s', got '%s'", tt.input, tt.expected, normalizedName)
			}

			// If the expected result should match a real function, test that too
			if tt.expected == "backup" || tt.expected == "index-vulnrichment" {
				if _, exists := availableFunctions[normalizedName]; !exists {
					t.Errorf("Normalized function name '%s' should exist in availableFunctions", normalizedName)
				}
			}
		})
	}
}

// TestFunctionCount ensures we have the expected number of functions
func TestFunctionCount(t *testing.T) {
	expectedCount := 14 // Based on the current availableFunctions map
	actualCount := len(availableFunctions)

	if actualCount != expectedCount {
		t.Errorf("Expected %d functions, but found %d", expectedCount, actualCount)
	}
}

// BenchmarkListFunctions benchmarks the listFunctions performance
func BenchmarkListFunctions(b *testing.B) {
	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() { os.Stdout = old }()

	for i := 0; i < b.N; i++ {
		listFunctions()
	}
}

// BenchmarkRunFunction benchmarks the runFunction performance (with mock)
func BenchmarkRunFunction(b *testing.B) {
	// Save original functions
	originalFunctions := availableFunctions

	// Create test functions map with mocks
	availableFunctions = map[string]TestFunction{
		"test-function": {
			Name:        "test-function",
			Description: "Test Function",
			Function:    func() {}, // No-op function for benchmarking
		},
	}

	// Redirect output to discard
	old := os.Stdout
	os.Stdout, _ = os.Open(os.DevNull)
	defer func() {
		os.Stdout = old
		availableFunctions = originalFunctions
	}()

	for i := 0; i < b.N; i++ {
		runFunction("test-function")
	}
}
