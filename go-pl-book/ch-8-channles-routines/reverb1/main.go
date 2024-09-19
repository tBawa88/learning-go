package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// This is an echo server, which echoes back whatever it reads from the connection

func main() {
	lstner, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal("Error creating a Listener ", err)
	}

	for {
		conn, err := lstner.Accept()
		if err != nil {
			fmt.Println("Connection error ", err)
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	input := bufio.NewScanner(conn) // NewScanner(Reader) requires a Reader and it returns a *bufio.Scanner from which we can read stuff

	// as long it Scan keeps reading from this reader
	for input.Scan() {
		echoo(conn, input.Text(), time.Millisecond*800)
	}

}

func echoo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t"+strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t"+shout)
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t"+strings.ToLower(shout))
}
