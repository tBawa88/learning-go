package main

import (
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error starting the server ", err)
	}

	go broadcaster() // start a seperate routine for broadcaster

	for {
		connection, err := listener.Accept() // blocks untill a client makes connection
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(connection)
	}

}
