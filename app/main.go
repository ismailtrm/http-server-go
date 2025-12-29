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
	Data string
}

type HTTP struct {
	method         string
	request_target string
	protocol       string
}

func (r Request) Parse() *HTTP {
	lines := strings.Split(r.Data, " ")
	http_req := new(HTTP)
	http_req.method = lines[0]
	http_req.request_target = lines[1]
	http_req.protocol = lines[2]
	return http_req
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO: Uncomment the code below to pass the first stage
	reqBuf := make([]byte, 1024)
	OK := []byte("HTTP/1.1 200 OK\r\n\r\n")
	NOT_FOUND := []byte("HTTP/1.1 404 Not Found\r\n\r\n")

	conn, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		clientConn, err := conn.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		req, err := clientConn.Read(reqBuf)

		if err == nil {
			fmt.Println("request catched")

			reqString := new(Request)
			reqString.Data = string(reqBuf[:req])
			fmt.Println("Received request:", reqString.Data)

			http_req := reqString.Parse()
			if http_req.request_target == "/" {
				fmt.Println(string(OK))
				clientConn.Write(OK)
			} else {
				fmt.Println(string(OK))
				clientConn.Write(NOT_FOUND)
			}
		}
		clientConn.Close()
	}
}
