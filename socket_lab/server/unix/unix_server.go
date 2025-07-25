package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	socketPath := "/tmp/unix_socket_example.sock"
	os.Remove(socketPath) // ensure clean start

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Unix server is listening on", socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Received from client:", string(buf[:n]))

	_, err = conn.Write([]byte("Hello from Unix server!"))
	if err != nil {
		fmt.Println("Write error:", err)
	}
}
