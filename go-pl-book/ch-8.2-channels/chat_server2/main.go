package main

import (
	"log"
	"net"
)

/*
	Make the broadcaster announce the current set of clients to each new arrival.
	This requires that the 'clients map' and the entering and leaving channels maintain client names too
*/

/*
	what we need is a seperate channel for each client
	a channel that listens for an entering client channel
	a chhannel that listens for a leaving client channel
	a message channel to which a client can send string messages to

	Process :
		- server maintains a list of all client channels cuurently connected
		- when a client is connected to the tcp server, a new channel is created for it and pushed to the entering channel
		- this channel reaches the broadcaster goroutine via another channel called 'listening', and it saves it in the map
		- this channel also acts as a medium of communication between  the server and the client
		- clients are not directly talkin to each other, they send a message to this channel, and the brodcaster, who is maintaining a map of all client channels,
			forwards this message to those channels which reach the clients
*/

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("Error creating a server ", err)
	}
	go broadcaster()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(connection)
	}
}
