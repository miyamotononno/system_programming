package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Server is running at localhost:8888");
	conn, err := net.ListenPacket("udp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1500)
	for {
		length, remoteAddress, err := conn.ReadFrom(buffer)
		fmt.Printf("Received from %s: %v\n", remoteAddress, string(buffer[:length]))
		_, err = conn.WriteTo([]byte("Hello from server"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}