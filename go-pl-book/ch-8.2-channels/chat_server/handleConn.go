package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
	- This function represents a new client connection , a new goroutine
	- We create a new outgoing channel, to which we send any incoming messages from this connection
	- When this function is called, there are couple of things we need to do right away
		- Start another goroutine, which will write to this channel every time a string is pushed to this outgoing channel
		- Emitt and event to entering channel, which will Add this connection's channel to the golbal set of channels
	- Create an bufio.NewScanner(conn), so that we can read incoming messages from this connection (for broadcasting)
	- Every time a message comes out of the input scanner, send that messgage to global message channel
*/

func handleConnection(conn net.Conn) {
	defer conn.Close()

	outgoing := make(chan string)
	go clientWriter(conn, outgoing)

	who := conn.RemoteAddr().String()
	outgoing <- "You are " + who

	messages <- who + " has entered the chat" //global channel for broadcasting messages
	entering <- outgoing                      // emitt an event to the entering channel, add this client's outgoing channel to the set

	// Read input from this connectoin
	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		messages <- who + " : " + text
	}

	// if the connection is closed, input.Scan() will evaluate to false, hence the loop will end
	leaving <- outgoing
	messages <- who + ": has left the chat"

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // writes to the connection
	}
}
