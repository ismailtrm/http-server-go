# HTTP Server in Go

An experimental HTTP/1.1 server implementation built from scratch in Go using only the `net` package. This project focuses on understanding low-level HTTP protocol mechanics without relying on Go's `net/http` package.

## About

This is a learning-oriented project that implements core HTTP server functionality manually. The goal is to deeply understand TCP socket programming, HTTP message format, and connection lifecycle management.

**Note:** Initially based on [CodeCrafters HTTP Server challenge](https://codecrafters.io/challenges/http-server), extended with additional features and learning experiments.

## What's Implemented

### 1. TCP Socket Management
- Raw TCP listener on port 4221 (using `net.Listen`)
- Manual connection acceptance and handling
- Proper resource cleanup with `defer`

### 2. Custom String Parsing
- Custom `str` type with `Split()` method (eliminates `strings` package dependency)
- Multi-character separator support (handles `\r\n`, `\r\n\r\n`)
- Demonstrates Go's type system and method receivers

### 3. HTTP Request Parsing
```go
type HTTP struct {
    method         string  // GET, POST, etc.
    request_target string  // URI path
    protocol       string  // HTTP/1.1
    body           string  // Request body (after CRLF separator)
}
```

Key parsing steps:
- Split headers and body using `\r\n\r\n` (CRLF double separator)
- Extract request line components (method, path, version)
- Parse request body when present

### 4. HTTP Response Generation
- Status codes: `200 OK`, `404 Not Found`, `405 Method Not Allowed`
- Proper HTTP/1.1 response format with CRLF line endings
- Method-based routing with `switch` statements

### 5. Request/Response Flow
```
Client → TCP Connection → Read Buffer → Parse Request → Route → Send Response → Close
```

## Technical Details

### CRLF Handling
HTTP uses `\r\n` (Carriage Return + Line Feed) for line breaks:
- Headers separated by `\r\n`
- Header/body boundary marked by `\r\n\r\n`
- Critical for HTTP protocol compliance (see [RFC 9112](https://datatracker.ietf.org/doc/html/rfc9112))

### Buffer Management
- 1024-byte read buffer for incoming requests
- Manual byte-to-string conversion
- No streaming (entire request loaded into memory)

### Connection Lifecycle
1. `Accept()` - Block until client connects
2. `Read()` - Read request data into buffer
3. Parse - Extract HTTP components
4. `Write()` - Send response
5. `Close()` - Terminate connection (via `defer`)

## Project Structure

```
.
├── app/
│   └── main.go       # Complete HTTP server implementation
├── go.mod            # Module definition (no external dependencies)
├── run.sh            # Build and run script
├── LICENSE           # MIT License
└── README.md         # This file
```

## Getting Started

### Prerequisites
- Go 1.25+

### Run Server
```bash
./run.sh
# or
go run app/main.go
```

Server listens on `0.0.0.0:4221`

### Test Endpoints
```bash
# GET request
curl http://localhost:4221/

# POST request
curl -X POST -d "data" http://localhost:4221/

# Verbose output (see full HTTP exchange)
curl -v http://localhost:4221/

# Test unsupported method
curl -X DELETE http://localhost:4221/
```

## Learning Resources

This project was built while studying:

- [MDN HTTP Guide](https://developer.mozilla.org/en-US/docs/Web/HTTP/Guides) - HTTP fundamentals
- [RFC 9112](https://datatracker.ietf.org/doc/html/rfc9112) - HTTP/1.1 message syntax and routing
- [CRLF Glossary](https://developer.mozilla.org/en-US/docs/Glossary/CRLF) - Line ending conventions

## Limitations

Educational project with intentional constraints:

- **No concurrency** - Single-threaded, handles one connection at a time
- **No header parsing** - Only parses request line, ignores headers (Content-Type, Content-Length, etc.)
- **Limited methods** - Only GET/POST supported
- **No body validation** - Assumes well-formed requests
- **No timeout handling** - Connections can hang indefinitely
- **No TLS/HTTPS** - Plain TCP only

**Not production-ready.**

## Key Takeaways

1. HTTP is text-based (understanding raw format is crucial)
2. CRLF (`\r\n`) matters (protocol compliance depends on it)
3. TCP provides reliable byte streams (no message boundaries)
4. Resource management requires discipline (`defer` for cleanup)
5. String manipulation can be done without stdlib (educational exercise)

## License

MIT

## Author

[ismailtrm](https://github.com/ismailtrm)