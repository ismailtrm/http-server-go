# HTTP Server in Go

An experimental HTTP/1.1 server implementation built from scratch in Go, using only the standard `net` package for learning purposes.

## About

This is a learning project focused on understanding HTTP protocol fundamentals and TCP socket programming in Go. The goal is to implement a basic HTTP server without relying on Go's `net/http` package or other high-level abstractions.

**Note:** This project was initially based on the [CodeCrafters HTTP Server challenge](https://codecrafters.io/challenges/http-server), but has been extended and personalized for independent learning.

## Features

- TCP connection handling
- HTTP/1.1 request parsing
- Custom string manipulation (no external dependencies)
- GET and POST method support
- Basic routing
- Proper connection lifecycle management with defer

## Project Structure

```
.
├── app/
│   └── main.go       # Main HTTP server implementation
├── go.mod            # Go module definition
├── run.sh            # Build and run script
└── README.md         # This file
```

## Getting Started

### Prerequisites

- Go 1.25 or higher

### Running the Server

```bash
# Using the run script
./run.sh

# Or directly with go run
go run app/main.go
```

The server will start listening on `0.0.0.0:4221`

### Testing

```bash
# Test root endpoint
curl http://localhost:4221/

# Test echo endpoint (if implemented)
curl http://localhost:4221/echo/hello

# Test with verbose output
curl -v http://localhost:4221/
```

## Implementation Details

### Custom String Type

The project implements a custom `str` type with a `Split()` method to demonstrate Go's type system and method receivers:

```go
type str string

func (s str) Split(sep byte) []string {
    // Custom split implementation
}
```

### Request Parsing

HTTP requests are parsed into structured types:

```go
type HTTP struct {
    method         string
    request_target string
    protocol       string
}
```

### Routing

The server uses a switch statement for clean method-based routing:

- `GET /` - Returns 200 OK
- `POST /` - Returns 200 OK
- Other methods - Returns 405 Method Not Allowed
- Unknown paths - Returns 404 Not Found

## Limitations

This is an educational project with intentional limitations:

- Single-threaded connection handling (no concurrency)
- Basic string parsing (no robust HTTP header parsing)
- No HTTP body handling
- No HTTPS/TLS support
- No production-ready error handling
- Limited HTTP method support

**Not intended for production use.**

## License

MIT

## Author

[ismailtrm](https://github.com/ismailtrm)
