package main

import (
    "fmt"
    "log"
    "net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
    // Log the incoming request method and URL
    log.Printf("Received request: %s %s", r.Method, r.URL.Path)

    // Analyze and extract HTTP headers
    headers := make(map[string][]string)
    for key, value := range r.Header {
        headers[key] = value
    }

    // Log the headers
    log.Printf("Request headers: %v", headers)

    // Respond with the analyzed headers
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"headers": %v}`, headers)
}

func main() {
    // Define a handler for the main URL
    http.HandleFunc("/", headersHandler)

    // Start the HTTP server
    port := ":8080" // Port to listen on
    log.Printf("Server is listening on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}