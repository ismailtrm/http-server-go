# HTTP Server in Go

A simple HTTP/1.1 server implementation written in Go from scratch.

## Features

- ✅ TCP connection handling
- ✅ HTTP/1.1 request parsing
- ✅ Custom string manipulation (no external dependencies)
- ✅ GET and POST method support
- ✅ Basic routing
- ✅ Proper connection lifecycle management with defer

## Project Structure

```
.
├── app/
│   └── main.go       # Main HTTP server implementation
├── go.mod            # Go module definition
└── README.md         # This file
```

## Getting Started

### Prerequisites

- Go 1.25 or higher

### Running the Server

```bash
go run app/main.go
```

The server will start listening on `0.0.0.0:4221`

### Testing

```bash
# Test root endpoint
curl http://localhost:4221/

# Test echo endpoint
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

## Learning Objectives

This project demonstrates:

1. TCP socket programming in Go
2. HTTP protocol fundamentals
3. Custom type creation and method receivers
4. Proper resource management with `defer`
5. String manipulation without external libraries
6. Error handling patterns in Go

## License

MIT

## Author

[ismailtrm](https://github.com/ismailtrm)