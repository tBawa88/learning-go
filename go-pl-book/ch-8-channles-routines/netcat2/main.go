package main

import (
	"io"
	"log"
	"net"
	"os"
)

// This is another version of netcat1 which connects to the reverb1 server and they exchange message over the connection
// Whatever we write in the terminal is sent to the server over the tcp connection
// The server then echoes it back to us

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server ", err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn) // goroutine stays the same, it reads from the connection and writes to os.Stdout
	mustCopy(conn, os.Stdin)     // the main routine is blocked here, it's reading from os.Stdin since the io.Copy() calls the Read() function of the Reader
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal("Error copying data from connection ", err)
	}
}
