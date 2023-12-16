package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MichalPitr/go-http-server/handler"
	"github.com/MichalPitr/go-http-server/parser"
)

func handleConnection(c net.Conn) {
	defer c.Close()

	// TODO: support arbitrarily large requests.
	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		log.Print(err)
		return
	}

	msg := string(buf[:n])
	log.Printf("Received request: %q", msg)
	httpReq, err := parser.Parse(msg)
	if err != nil {
		log.Println(err)
		c.Write([]byte("HTTP/1.1 500 INTERNAL ERROR\r\n"))
		return
	}

	// Only subset of HTTP/1.1 is supported
	if httpReq.Method != "GET" || httpReq.Protocol != "HTTP/1.1" {
		c.Write([]byte("HTTP/1.1 400 BAD REQUEST\r\n"))
		return
	}

	res := handler.HandleGet(httpReq)
	if _, err := c.Write(res); err != nil {
		log.Printf("Error writing response: %v\n", err)
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", 2020, "server port")
	flag.Parse()
	if port < 1024 {
		log.Println("Using reserved port.")
		os.Exit(1)
	}

	address := fmt.Sprintf("localhost:%d", port)
	ln, err := net.Listen("tcp", address)
	defer ln.Close()

	if err != nil {
		log.Printf("Error creating a server: %v", err)
		os.Exit(1)
	}
	log.Printf("Starting server on %s\n", address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}
