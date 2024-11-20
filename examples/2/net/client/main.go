package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("Failed to connect: %s", err.Error())
	}
	log.Print("Connected to the server")

	go handleCopy(conn, os.Stdin)
	go handleCopy(os.Stdout, conn)

	c := make(chan struct{})
	<-c
}

func handleCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalf("Failed to copy: %s", err.Error())
	}
}
