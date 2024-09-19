package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// A minor problem with reverb1 was that if another input is entered immidiately, it has to wait untill the previou call has finished,
// But this is not realistic, since a real echo reverb starts immidiately
/*
hello?
	HELLO?
heyy
	hello?
	hello?
	HEYY
	heyy
	heyy
*/

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
	defer conn.Close()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		go echoo(conn, input.Text(), time.Millisecond*800) // starting another goroutine to handle the ehco part, since it involves multiple blocking calls t time.Sleep()
		// if another input comes from the connection, a new goroutine will be created to handle that input
	}

}

func echoo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t"+strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t"+shout)
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t"+strings.ToLower(shout))
}
