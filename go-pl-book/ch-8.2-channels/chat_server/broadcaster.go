package main

/*
	High level Approach: Create a new channel for each client
		- Each client can only send messages to this client channel
		- Every time a string comes out of this channel, it is sent into other connected client's channels

	Low level Approach:
		- Create a new type which is a wrapper over a send only client of string type
		- Create 2 new channels which are of this type (channels of channels)
		- Since each client is getting its own channel, the entering channel will keep track of these channels
		- the leaving channel will keep track of clients that want to disconnected

		- We also maintain a 'set'/'dictionary' of all the clients that are currently connected to the server
		- clients := make(map[client] bool)

	Events:
		- If a client connects, an event(client channel) will be pushed onto the entering channel
		- If a client leaves, an event will be pushed onto the leaving channel
		- The responsible of handling these events and pushing clients to these channels is handeled by func handleConnection()
		- If a client sends a message to it's client channel, it will be received by this broadcaster(), and forwarded to all connected clients
*/

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		// listening to multiple channels at once
		select {

		case msg := <-messages:
			// broadcasts incoming message to all clients
			for cli := range clients {
				cli <- msg
			}

		case newClient := <-entering:
			clients[newClient] = true // if new client comes to this channel, add it to the client dictionary

		case disconClient := <-leaving:
			delete(clients, disconClient)
			close(disconClient)

		}
	}
}
