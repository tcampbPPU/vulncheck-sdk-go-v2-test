package main

import (
	"github.com/dbugapp/dbug-go/dbug"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/cli"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/utils"
)

func init() {
	// Load environment variables from .env file
	utils.LoadEnv()

	// Log the start of the SDK examples
	dbug.Go("Running SDK Examples...")
}

func main() {
	// Run the CLI interface
	cli.Run()
}
