package main

import (
	"io"
	"log"
	"net"
	"time"
)

// A concurrent clock server that writes current time to the client, once every second
// net.Listen() returns a listener which is of type net.Listener interface
// it describes how a listener should behave. It has Accept(), Close() and Addr() methods on it
// Accept() starts a new connection im assuming

// NOTE : this code starts a server that connects to the client over the tcp protocol
// Therefore we cannot connect to it using a browser since browser communicates over the HTTP protocol
// to connect to this server, either use telnet or netcat (nc localhost 8080) or write a tcp client in either Go or Python
func main() {
	lstenr, err := net.Listen("tcp", "localhost:8080")
	if err != nil {

		log.Fatal(err)
	}

	// start an infinite loop
	for {
		conn, err := lstenr.Accept() // it's a blocking  call, it WAITS untill a TCP connection is made by the client to the address given to net.Listen()
		// it also must be implementing some sort of queue behind the scene to hold all the client connections, and only when the main goroutine comes here
		// it pop_fronts the queue and passes it to the handleConn() function
		if err != nil {
			log.Print("Failed to create a connection / connection closed")
			continue // if a connection fails, start a another one
		}
		go handleConn(conn)
		// If we don't start another goroutine to handle the connection, we wont' be able to connect to another client
		// because in that case the main goroutine will be stuck inside the infinite loop of handleConn()
		// but if we're spawning a goroutine to handle the connection and write to it
		// the mainroutine will be stuck at lisner.Accept(), waiting for a new connection
		// and every time a new client connects to this server, mainroutine will spawn another routine to handle that connection seperately
	}
}

// con is an interface type, it defines all the behaviours of a connection (what methods it has)
// it has Read() Write() Close() LocalAddr() RemoteAddr()  methods on it,
//
//	which means we can read/write bytes to the connection object and close the connectoin whenever we want
func handleConn(conn net.Conn) {
	defer conn.Close() // close this connection after exiting this function , another good use case for defer function call
	// start another infinite loop which keep writing to the connection once every second
	for {
		_, err := io.WriteString(conn, time.Now().Format("15:05:02\n"))
		if err != nil {
			log.Print("Error writing to the client ", err) // meaning the client closed the connection
			return
		}
		time.Sleep(time.Second * 1) // sleep this goroutine for 1 second and then it exits
	}

}
