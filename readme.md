# Go HTTP Header Inspection

This is a simple Go application that runs an HTTP server to inspect incoming requests and analyze their HTTP headers. It provides a basic API endpoint to view the headers of requests made to the server.

## Installation

### Option 1: Build from Source

1. Clone the repository:

   ```bash
   git clone https://github.com/priyankub/header-inspection.git
   ```

2. Navigate to the project directory:

    ```bash
    cd header-inspection
    ```

3. Build the Go application:

    ```bash
    go build
    ```

4. Run the application:

    ```bash
    ./header-inspection
    ```

5. Open your web browser or use a tool like cURL to make requests to the server:

    ```bash
    curl http://localhost:8080
    ```

    This will display the HTTP headers of the request.

### Option 2: Use Docker

You can install the application using Docker. Simply pull the Docker image from the GitHub Container Registry:

```bash
   docker pull ghcr.io/priyankub/header-inspection:latest 
```

#### Usage

1. Run the Docker container:

    ```bash
    docker run -p 8080:8080 ghcr.io/priyankub/header-inspection:latest
    ```

    This will start the application and expose it on port 8080.

2. Open your web browser or use a tool like cURL to make requests to the server:

    ```bash
    curl http://localhost:8080
    ```

    This will display the HTTP headers of the request.

## API Endpoints

GET /: Retrieves the HTTP headers of the incoming request.

## Dependencies

This project has no external dependencies beyond the standard library included with Go.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests for any improvements or new features you'd like to see.

## License

MIT License
