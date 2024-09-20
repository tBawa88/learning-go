package main

import (
	"io"
	"log"
	"net"
	"os"
)

// In netcat2, the job of copying the the input from Stdin to the connection is of the main goroutine
// if the Stdin is closed, the main routine will exit the program without any regards to the other goroutine
// To stop this we can use **unbuffered channel** to block the main goroutine untill the child routine closes the connection or something

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Error connecting to the server ", err)
	}
	done := make(chan struct{}) // a channel whose type is an empty struct

	go func() {
		io.Copy(os.Stdout, conn) // remember this is blocking call, therfore unless the conn is closed by the server,
		log.Print("done")
		done <- struct{}{} // putting an empty struct inside the channel
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

// Whats happening? : If this program's stdin is closed by the user, the main goroutine which is inside mustCopy() function , returns, and calls conn.Close() which closes the conn
// when the connection closes, the main routine won't immidiately end the program, instead it will wait for a value to come out of the channel (unbuffered synchronous channel)
// Now since the connection has been close, the background goroutine that is executing the literal function, will also come out of io.Copy() call
// after that it prints "done" and sends a value to the channel
// the main routine receives this value and the program exits after printing "done" to the terminal
// INSTEAD of exiting suddenly

func mustCopy(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal("Error copying data from connection ", err)
	}
}
