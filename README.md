# web-server-go

HTTP web server built with Go's standard library. 
Zero external dependencies.

## Stack

- **Language:** Go 1.26.5
- **Dependencies:** None — standard library only

## Routes

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Serves `static/index.html` |
| GET | `/form.html` | Serves the contact form |
| GET | `/hello` | Plain text response |
| POST | `/form` | Parses form body, returns JSON |
| GET | `/api/health` | Server status and uptime |

## Run

```bash
git clone https://github.com/adhyayana108/web-server-go.git
cd web-server-go
go run main.go
```

Server starts at `http://localhost:8080`.

```bash
# Custom port
PORT=9090 go run main.go

# Build binary
go build -o server .
./server
```

## Test

```bash
# Hello
curl http://localhost:8080/hello

# Submit form
curl -X POST http://localhost:8080/form \
  -d "name=Arjun&email=arjun@example.com&address=Lucknow"

# Health check
curl http://localhost:8080/api/health
```

## Project Structure

```
go-server/
├── go.mod
├── main.go
└── static/
    ├── index.html
    └── form.html
```

## Architecture

```
Request → loggingMiddleware → ServeMux → handler → Response
```

Every request passes through `loggingMiddleware` which logs method, path, status code, and duration. The `ServeMux` routes by path to the correct handler.

## What This Project Demonstrates

- Creating an HTTP server using Go's standard library
- Routing requests with `http.ServeMux`
- Handling GET and POST requests
- Serving static HTML files
- Processing form data and returning JSON responses
- Using middleware for request logging
- Using environment variables to configure the server port
- Setting server timeouts for safer request handling

