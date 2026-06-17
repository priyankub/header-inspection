package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{"headers": r.Header}); err != nil {
		log.Printf("Error encoding JSON: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headersHandler)

	port := ":8080"
	srv := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,   // Max time to read request from the client
		WriteTimeout: 10 * time.Second,  // Max time to write response to the client
		IdleTimeout:  120 * time.Second, // Max time for connections using TCP Keep-Alive
	}

	log.Printf("Server is listening on port %s...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
