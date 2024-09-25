package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

const timemout = time.Second * 5

func handleConn(conn net.Conn) {
	defer conn.Close()
	clientChan := make(chan string)
	go clientWriter(conn, clientChan)

	who := conn.RemoteAddr().String()
	clientChan <- "You are : " + who
	messages <- who + " has joined the chat"
	entering <- clientChan

	// start the timer
	timer := time.NewTimer(timemout)
	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	for input.Scan() {
		text := input.Text()
		messages <- who + " : " + text
		timer.Reset(timemout) // just reset the same timer with same timeout
	}

	//wait for either connectoin to close, or client getting kicked due to afk
	leaving <- clientChan
	messages <- who + " has left the chat"

}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
