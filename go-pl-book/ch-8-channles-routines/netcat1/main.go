package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Writing a read only TCP client that connects to a TCP server
// IT reads from the connection and writes the data to os.Stdout

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to server ", err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	// the io.Copy() is blocking call, it keeps reading untill it reaches EOF , in this case as long as the data keeps coming to the connection it will keep reading from it
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal("Error copying data from connection ", err)
	}
}
