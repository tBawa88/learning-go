package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

// This walks the entire tree recursively and prints the outline of the entire html strucutre to Stdout

func main() {
	file, err := os.Open("test.html")
	if err != nil {
		log.Fatal("Error opening the file", err)
	}
	data, err := html.Parse(file)

	if err != nil {
		log.Fatal("Error parsing the html file ", err)
	}

	outline(data, nil)

}

func outline(n *html.Node, stack []string) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) //push the html tag onto the stack
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = n.NextSibling {
		fmt.Println(c.Data)
		outline(c, stack)
	}
}
