package main

import (
	"io"
	"log"
	"net"
	"os"
)

// in netcat3, as soon as the stdout is closed by pressing Ctrl + D, the main routine comes out of the mustCopy() function and closes the connection
// which forces the background goroutine to come out of it's literal function

// net package has another type called *net.TCPConn which is an implementation of net.Conn interface
// A TCP connection consists of two halves, a "read" half and a write "half", both of em can be closed indepently using CloseRead() and CloseWrite()
// We want to make a program which will not close the write connection as soon as the read half is closed by the main routine

func main() {
	var conn *net.TCPConn
	// net.DialTCP(network, localAddr, remoteAddr) is specifically for creating a client connection that connects to a TCP server
	conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8080})
	if err != nil {
		log.Fatal("Error connecting to TCP server ", err)
	}

	done := make(chan struct{})
	go func() { // this goroutine is using the "read" half of the TCP connection, since it's reading from the connection and copying it to stdout
		io.Copy(os.Stdout, conn)
		log.Print("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin) // the main goroutine is using the "write" half of the TCP connection, it reads from the stdin file and
	conn.CloseWrite()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal("Error copying data from connection ", err)
	}
}
