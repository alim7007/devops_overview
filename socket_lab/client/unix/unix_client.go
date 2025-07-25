package main

import (
	"fmt"
	"net"
)

func main() {
	socketPath := "/tmp/unix_socket_example.sock"

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "Hello from Unix client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received from server:", string(buf[:n]))
}
