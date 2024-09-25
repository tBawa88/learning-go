package main

import (
	"bufio"
	"fmt"
	"net"
)

/*
	This goroutine represents a single client connection
	Things that we must do in this function
		- Create a channel and push it to the entering channel
		- Send a message to the messages channel announcing this clients arrival
		- Start another goroutine that listens to the stdin and sends the entered text to the messages channel
*/

func handleConnection(conn net.Conn) {
	defer conn.Close()

	outgoing := make(chan string)
	go clientWriter(conn, outgoing)

	who := conn.RemoteAddr().String()
	outgoing <- "You are :" + who
	messages <- who + " has joined the chat"

	newClient := clientInfo{channel: outgoing, name: who}
	entering <- newClient

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		fmt.Println("client said ", text)
		messages <- who + " : " + text
	}

	// in the event of leaving
	leaving <- newClient
	messages <- who + " has left the chat"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

// func afkCheck (conn net.Conn)
