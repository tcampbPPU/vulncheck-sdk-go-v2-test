package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/dbugapp/dbug-go/dbug"
	"github.com/joho/godotenv"
	"github.com/tcampbPPU/sdk-go-v2-test/pkg/cli"
)

func init() {
	// Load environment variables from .env file
	loadEnv()

	// Log the start of the SDK examples
	dbug.Go("Running SDK Examples...")
}

func main() {
	// Run the CLI interface
	cli.Run()
}

func loadEnv() {
	wd, err := os.Getwd()
	if err != nil {
		log.Printf("Warning: Could not get working directory: %v", err)
	} else {
		for dir := wd; dir != filepath.Dir(dir); dir = filepath.Dir(dir) {
			envPath := filepath.Join(dir, ".env")
			if _, err := os.Stat(envPath); err == nil {
				err = godotenv.Load(envPath)
				if err != nil {
					log.Printf("Warning: Could not load .env file from %s: %v", envPath, err)
				}
				break
			}
		}
	}
}
