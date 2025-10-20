package main

import (
	"fmt"

	"github.com/ducthinh993/renovate-gomodtidyall-test/api"
	"github.com/ducthinh993/renovate-gomodtidyall-test/sdk"
)

func main() {
	// Initialize SDK client
	client := sdk.NewClient("web-app")

	// Process a message using SDK (which uses shared)
	response := client.ProcessMessage("Hello from web app!")
	fmt.Printf("SDK Response: %+v\n", response)

	// Setup API routes to test the API dependency
	_ = api.SetupRoutes()
	fmt.Println("API router configured successfully")

	fmt.Printf("Web application started with %s\n", client.GetName())
}