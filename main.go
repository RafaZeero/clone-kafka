package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Close the connection when the function returns
	defer conn.Close()

	// Make a buffer to hold incoming data
	buffer := make([]byte, 1024)

	// Read data from the connection until it's closed by the client
	for {
		// Read data from the connection
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}

		// Print the received data
		fmt.Println("Received:", string(buffer[:n]))

		// Echo the data back to the client
		_, err = conn.Write([]byte("Henlo!!"))
		if err != nil {
			fmt.Println("Error writing:", err.Error())
			return
		}
	}
}

func main() {
	// Listen for incoming connections on port 3000
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	defer listener.Close()
	fmt.Println("Server listening on port 3000")

	// Accept incoming connections and handle them concurrently
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}
		fmt.Println("Accepted connection from:", conn.RemoteAddr())

		go handleConnection(conn)
	}
}
