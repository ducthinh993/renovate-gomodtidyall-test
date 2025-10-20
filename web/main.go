package main

import (
	"fmt"
	"log"

	"github.com/ducthinh993/renovate-gomodtidyall-test/api"
	"github.com/ducthinh993/renovate-gomodtidyall-test/sdk"
)

func main() {
	// Initialize SDK client
	client := sdk.NewClient("web-app")

	// Process a message using SDK (which uses shared)
	response := client.ProcessMessage("Hello from web app!")
	fmt.Printf("SDK Response: %+v\n", response)

	// Setup API routes
	router := api.SetupRoutes()
	fmt.Println("API router configured successfully")

	// In a real app, you would start the HTTP server here
	// log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Printf("Web application started with %s\n", client.GetName())
}