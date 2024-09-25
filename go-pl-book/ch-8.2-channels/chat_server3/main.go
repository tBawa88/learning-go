package main

import (
	"log"
	"net"
)

/*
	In this variation, we're trying to kick out clients who are inactive for more than 10 seconds
*/

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal("Erro creating a listener ", err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
