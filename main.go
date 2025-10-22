package main

import (
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/cli"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/utils"
)

func init() {
	// Load environment variables from .env file
	utils.LoadEnv()
}

func main() {
	// Run the CLI interface
	cli.Run()
}
