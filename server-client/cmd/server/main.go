package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer server.Close()
	fmt.Println("Listening on  " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting on client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	defer connection.Close()
	for {
		mLen, err := connection.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client closed connection")
				break
			} else {
				fmt.Println("Error reading:", err.Error())
			}
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
		_, err = connection.Write([]byte("Thanks! Got your message:" + string(buffer[:mLen])))
	}
}
