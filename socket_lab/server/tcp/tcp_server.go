package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on 0.0.0.0:9000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Received from client:", string(buffer[:n]))

	response := "Hello from server!"
	conn.Write([]byte(response))
}
