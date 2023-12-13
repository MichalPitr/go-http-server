package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/MichalPitr/go-http-server/parser"
)

func handleConnection(c net.Conn) {
	defer c.Close()

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

	// TODO: Add helper that only lets through valid and supported requests.
	//
	// For instance, Only GET, POST is supported, only HTTP/1.1, ... etc
	if httpReq.Method != "GET" || httpReq.Protocol != "HTTP/1.1" {
		c.Write([]byte("HTTP/1.1 400 BAD REQUEST\r\n"))
		return
	}

	// TODO: Write a helper to return file stored at path.
	// TODO: index should be accessible both with "/" and "/index.html"
	f, err := os.ReadFile("./www" + httpReq.Path)
	if err != nil {
		log.Println(err)
		c.Write([]byte("HTTP/1.1 404 NOT FOUND\r\n"))
		return
	}

	// TODO: write helper that takes parsed request and returns appropriate response
	c.Write([]byte(fmt.Sprintf("HTTP/1.1 200 OK\r\n\r\n%s\r\n", f)))
}

func main() {
	port := ":2020"
	ln, err := net.Listen("tcp", port)
	defer ln.Close()

	if err != nil {
		log.Printf("Error creating a server: %v", err)
		os.Exit(1)
	}
	log.Printf("Starting server on port %s\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}
