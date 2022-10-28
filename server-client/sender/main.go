package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	connection, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for true {
		fmt.Println("What message do you want to send? (type exit to exit)")

		var message string
		if scanner.Scan() {
			message = scanner.Text()
		}

		if scanner.Err() != nil {
			fmt.Println("error reading input...")
			continue
		}

		if message == "exit" {
			break
		}
		writeMessage(connection, message)
		readResponse(connection)
	}

	connection.Close()
}

func writeMessage(connection net.Conn, message string) {
	fmt.Println("Attempting to send message:", message)
	_, err := connection.Write([]byte(message))
	if err != nil {
		fmt.Println("Failed to send message:", err.Error())
	}
}

func readResponse(connection net.Conn) {
	buffer := make([]byte, 1024)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received: ", string(buffer[:mLen]))
}
