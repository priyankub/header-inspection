package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Define a simple, clean HTML template with a table
const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>HTTP Header Inspection</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Helvetica, Arial, sans-serif; background-color: #f4f4f9; color: #333; padding: 2rem; margin: 0; }
        .container { max-width: 800px; margin: 0 auto; background: #fff; padding: 2rem; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); }
        h1 { text-align: center; color: #2c3e50; }
        table { border-collapse: collapse; width: 100%; margin-top: 1.5rem; }
        th, td { padding: 12px 16px; text-align: left; border-bottom: 1px solid #e0e0e0; word-break: break-all; }
        th { background-color: #2980b9; color: white; font-weight: 600; }
        tr:hover { background-color: #f1f5f8; }
        td strong { color: #2c3e50; }
    </style>
</head>
<body>
    <div class="container">
        <h1>Incoming Request Headers</h1>
        <table>
            <thead>
                <tr>
                    <th>Header Name</th>
                    <th>Value(s)</th>
                </tr>
            </thead>
            <tbody>
                {{range $key, $values := .}}
                <tr>
                    <td><strong>{{$key}}</strong></td>
                    <td>
                        {{range $values}}
                            <div>{{.}}</div>
                        {{end}}
                    </td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
</body>
</html>
`

func headersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	// Parse the HTML template
	tmpl, err := template.New("headers").Parse(htmlTemplate)
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Execute the template, passing the HTTP headers as the data
	if err := tmpl.Execute(w, r.Header); err != nil {
		log.Printf("Error executing template: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", headersHandler)

	port := ":8080"
	srv := &http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server is listening on port %s...\n", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}