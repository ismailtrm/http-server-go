package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

type Request struct {
	Buffer []byte
	Data   string
}

type HTTP struct {
	method         string
	request_target string
	protocol       string
}

const (
	OK        = "HTTP/1.1 200 OK\r\n\r\n"
	NOT_FOUND = "HTTP/1.1 404 Not Found\r\n\r\n"
)

func (r Request) Parse() *HTTP {
	lines := strings.Split(r.Data, " ")
	http_req := new(HTTP)
	http_req.method = lines[0]
	http_req.request_target = lines[1]
	http_req.protocol = lines[2]
	return http_req
}

func handler(conn net.Listener) {
	request := new(Request)
	request.Buffer = make([]byte, 1024)

	clientConn, err := conn.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	defer clientConn.Close()

	req, err := clientConn.Read(request.Buffer)

	if err == nil {
		fmt.Println("Request catched")

		request.Data = string(request.Buffer[:req])

		http_req := request.Parse()

		if http_req.method == "GET" && http_req.request_target == "/" {
			fmt.Println("OK")
			clientConn.Write([]byte(OK))
			fmt.Println("Received request:", request.Data)
		} else {
			fmt.Println("NOT_FOUND")
			clientConn.Write([]byte(NOT_FOUND))
			fmt.Println("Received request:", request.Data)
		}
	}
}
func main() {
	fmt.Println("Logs will appear here!")

	conn, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		handler(conn)
	}
}
