package main

import (
	"fmt"
	"net/http"

	"github.com/asif2772/campaign-spend-tracker/pkg/config"
)

func main() {
	// Load application configuration
	cfg := config.Load()

	// Register the health check endpoint
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	// Log server startup
	fmt.Printf("Server starting on port %s\n", cfg.Port)

	// Start HTTP server
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
