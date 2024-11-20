package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Print("Server started")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	log.Printf("Received new connection from %s", conn.RemoteAddr().String())

	if _, err := io.Copy(conn, conn); err != nil {
		log.Printf("Failed to forward to %s: %s", conn.RemoteAddr().String(), err.Error())
	}
}
