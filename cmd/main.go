package main

import (
    "fmt"
    "log"
    "net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
    // Analyze and extract HTTP headers
    headers := make(map[string][]string)
    for key, value := range r.Header {
        headers[key] = value
    }

    // Respond with the analyzed headers
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, `{"headers": %v}`, headers)
}

func main() {
    // Define a handler for the /headers route
    http.HandleFunc("/headers", headersHandler)

    // Start the HTTP server
    port := ":8080" // Port to listen on
    fmt.Printf("Server is listening on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
