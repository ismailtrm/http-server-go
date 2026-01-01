package main

import (
	"fmt"
	"net"
	"os"
)

type str string

type Request struct {
	Buffer []byte
	Data   str
}

type HTTP struct {
	method         string
	request_target string
	protocol       string
	body           string
}

const (
	OK                 = "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND          = "HTTP/1.1 404 Not Found\r\n\r\n"
	METHOD_NOT_ALLOWED = "HTTP/1.1 405 Method Not Allowed\r\n\r\n"
)

// Split splits a string by a string separator
func (s str) Split(sep string) []string {
	var result []string
	start := 0

	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && string(s[i:i+len(sep)]) == sep {
			result = append(result, string(s[start:i]))
			start = i + len(sep)
			i += len(sep) - 1 // Skip separator
		}
	}

	result = append(result, string(s[start:]))
	return result
}

// Parse parses the HTTP request into structured format
func (r Request) Parse() *HTTP {
	http_req := new(HTTP)

	parts := r.Data.Split("\r\n\r\n")
	header := str(parts[0]).Split("\r\n")
	firstLine := header[0]

	lines := str(firstLine).Split(" ")
	http_req.method = lines[0]
	http_req.request_target = lines[1]
	http_req.protocol = lines[2]

	if len(parts) > 1 {
		http_req.body = parts[1]
	} else {
		http_req.body = ""
	}

	return http_req
}

// handler handles incoming client connections
func handler(conn net.Listener) {
	request := new(Request)
	request.Buffer = make([]byte, 1024)

	clientConn, err := conn.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
		os.Exit(1)
	}

	defer clientConn.Close()

	req, err := clientConn.Read(request.Buffer)

	if err == nil {
		fmt.Println("Request received")

		request.Data = str(request.Buffer[:req])

		http_req := request.Parse()

		switch http_req.method {
		case "GET":
			if http_req.request_target == "/" {
				fmt.Println("200 OK")
				clientConn.Write([]byte(OK))
			} else {
				fmt.Println("404 Not Found")
				clientConn.Write([]byte(NOT_FOUND))
			}
		case "POST":
			if http_req.request_target == "/" {
				fmt.Println("200 OK")
				fmt.Println(str(request.Buffer))
				clientConn.Write([]byte(OK))
			} else {
				fmt.Println("404 Not Found")
				clientConn.Write([]byte(NOT_FOUND))
			}
		default:
			fmt.Println("405 Method Not Allowed")
			clientConn.Write([]byte(METHOD_NOT_ALLOWED))
		}
	}
}

func main() {
	fmt.Println("HTTP Server starting on 0.0.0.0:4221...")

	conn, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	fmt.Println("Server is listening for connections...")

	for {
		handler(conn)
	}
}
