package main

import (
	"log"
	"net"
	"os"
	"strings"
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

	// TODO: add parser for requests
	log.Printf("Received request: %q", msg)
	if strings.Split(msg, "\r\n")[0] != "GET / HTTP/1.1" {
		c.Write([]byte("HTTP/1.1 400 BAD REQUEST\r\n"))
		return
	}

	// TODO: write helper that takes parsed request and returns appropriate response
	c.Write([]byte("HTTP/1.1 200 OK\r\n\r\nRequested path: <the path>\r\n"))
}

func main() {
	port := ":2020"
	ln, err := net.Listen("tcp", port)
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
