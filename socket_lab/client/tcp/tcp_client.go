package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "198.19.249.86:9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	message := "Hello from client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received from server:", string(buffer[:n]))
}
