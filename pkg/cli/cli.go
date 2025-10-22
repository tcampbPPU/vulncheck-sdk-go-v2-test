package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/tcampbPPU/sdk-go-v2-test/pkg/backup"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/browse"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/cpe"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/index"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/pdns"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/purl"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/rule"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/tag"
)

// TestFunction represents a test function with metadata
type TestFunction struct {
	Name        string
	Description string
	Function    func()
}

// availableFunctions maps function names to their implementations
var availableFunctions = map[string]TestFunction{
	"index-initial-access": {
		Name:        "index-initial-access",
		Description: "Get Index Initial Access",
		Function:    index.GetIndexInitialAccess,
	},
	"index-vulnrichment": {
		Name:        "index-vulnrichment",
		Description: "Get Index Vulnrichment",
		Function:    index.GetIndexVulnrichment,
	},
	"index-cve-filter": {
		Name:        "index-cve-filter",
		Description: "Get Index with CVE Filter",
		Function:    index.GetIndexWithCveFilter,
	},
	"index-botnet-filter": {
		Name:        "index-botnet-filter",
		Description: "Get Index with Botnet Filter",
		Function:    index.GetIndexWithBotnetFilter,
	},
	"index-ip-intel": {
		Name:        "index-ip-intel",
		Description: "Get Index IP Intel",
		Function:    index.GetIndexIpIntel,
	},
	"index-canary": {
		Name:        "index-canary",
		Description: "Get Index Canaries",
		Function:    index.GetIndexCanaries,
	},
	"browse-indexes": {
		Name:        "browse-indexes",
		Description: "Browse Indexes",
		Function:    browse.BrowseIndexes,
	},
	"browse-backups": {
		Name:        "browse-backups",
		Description: "Browse Backups",
		Function:    browse.BrowseBackups,
	},
	"backup": {
		Name:        "backup",
		Description: "Get Index Backup",
		Function:    backup.GetIndexBackup,
	},
	"purl": {
		Name:        "purl",
		Description: "Get PURL",
		Function:    purl.GetPurl,
	},
	"rule": {
		Name:        "rule",
		Description: "Get Rule",
		Function:    rule.GetRule,
	},
	"tag": {
		Name:        "tag",
		Description: "Get Tag",
		Function:    tag.GetTag,
	},
	"pdns": {
		Name:        "pdns",
		Description: "Get PDNS",
		Function:    pdns.GetPdns,
	},
	"cpe": {
		Name:        "cpe",
		Description: "Get CPE",
		Function:    cpe.GetCpe,
	},
}

// Run handles the command-line interface logic
func Run() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "list", "ls", "ll":
		listFunctions()
	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please specify a function name to run")
			fmt.Println("Usage: go run main.go run <function-name>")
			fmt.Println("Use 'go run main.go list' to see available functions")
			return
		}
		runFunction(os.Args[2])
	case "help", "-h", "--help":
		showUsage()
	default:
		// Try to run the command directly as a function name
		runFunction(command)
	}
}

// showUsage displays the CLI usage information
func showUsage() {
	fmt.Println("SDK Test CLI")
	fmt.Println("Usage:")
	fmt.Println("  go run main.go list                    # List all available test functions")
	fmt.Println("  go run main.go run <function-name>     # Run a specific test function")
	fmt.Println("  go run main.go <function-name>         # Run a specific test function (shorthand)")
	fmt.Println("  go run main.go help                    # Show this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go list")
	fmt.Println("  go run main.go run index-vulnrichment")
	fmt.Println("  go run main.go index-vulnrichment")
}

// listFunctions displays all available test functions
func listFunctions() {
	fmt.Println("Available test functions:")
	fmt.Println()

	// Sort function names for consistent output
	var names []string
	for name := range availableFunctions {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		testFunc := availableFunctions[name]
		fmt.Printf("  %-25s - %s\n", testFunc.Name, testFunc.Description)
	}

	fmt.Println()
	fmt.Println("Usage: go run main.go run <function-name>")
	fmt.Println("   Or: go run main.go <function-name>")
}

// runFunction executes a specific test function by name
func runFunction(functionName string) {
	// Normalize function name (remove spaces, convert to lowercase)
	normalizedName := strings.ToLower(strings.ReplaceAll(functionName, " ", "-"))

	testFunc, exists := availableFunctions[normalizedName]
	if !exists {
		fmt.Printf("Error: Function '%s' not found\n", functionName)
		fmt.Println("Use 'go run main.go list' to see available functions")
		return
	}

	fmt.Printf("Running: %s\n", testFunc.Description)
	fmt.Println(strings.Repeat("-", 50))
	testFunc.Function()
}
