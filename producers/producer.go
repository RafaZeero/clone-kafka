package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// Connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	args := os.Args

	idx := findStringIndex("-m", args)

	var message string

	if idx == -1 {
		fmt.Println("No message provided")
		return
	} else {
		message = args[idx+1]
	}

	// Send data to the server
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err.Error())
		return
	}

	// Read response from the server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading response:", err.Error())
		return
	}

	// Print the response from the server
	fmt.Println("Response:", string(buffer[:n]))
}

func findStringIndex(target string, slice []string) int {
	for i, s := range slice {
		if s == target {
			return i // Return the index of the target string
		}
	}
	return -1 // Return -1 if the string is not found
}
