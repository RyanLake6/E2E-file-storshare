package main

import (
	"E2E-file-storshare/cmd"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	
	cmd.Execute()
}
