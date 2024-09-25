package main

type client chan string

type clientInfo struct {
	channel client
	name    string
}

var (
	entering = make(chan clientInfo)
	leaving  = make(chan clientInfo)
	messages = make(chan string)
)

func broadcaster() {
	clientSet := make(map[clientInfo]bool)

	for {
		select {
		case cli := <-entering:
			// iterating over all present clients and sending their names to newly current client
			for client := range clientSet {
				cli.channel <- client.name + " is present"
			}
			clientSet[cli] = true

		case msg := <-messages:
			for cli := range clientSet {
				cli.channel <- msg
			}

		case cli := <-leaving:
			delete(clientSet, cli)
			close(cli.channel)
		}
	}
}
