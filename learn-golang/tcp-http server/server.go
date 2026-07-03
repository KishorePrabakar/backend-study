package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Port failed to listen: %v", err)
	}

	defer listener.Close()

	fmt.Println("Server listening on port :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to Connect: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Client disconnected: %v", err)
			break
		}

		fmt.Printf("Recieved: %v", message)
		conn.Write([]byte("Echo:" + message))
	}
}
