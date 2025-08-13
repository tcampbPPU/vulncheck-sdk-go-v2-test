package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {
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
