package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// This program is copy of the main.go program in recursion/main.go

func main() {
	htmlFile, err := os.Open("test.html")
	if err != nil {
		fmt.Println("Error opening the HTML file ", err)
		os.Exit(1)
	}

	// Use this file as a source parse out the HTMl from it and obtain the root node of parsed HTML tree
	data, err := html.Parse(htmlFile)
	if err != nil {
		fmt.Println("Error parsing the HTML file ", err)
		os.Exit(1)
	}

	options := traverse(data, make([]string, 0))
	for _, o := range options {
		fmt.Println(o)
	}

}

func traverse(n *html.Node, list []string) []string {
	if n.Type == html.ElementNode && n.Data == "option" {
		// Traverse over all of it's attributes and push the value attribute onto the string slice
		for _, option := range n.Attr {
			if option.Key == "value" {
				list = append(list, option.Val)
			}
		}
	}

	// Start iterating over current nodes children untill the 'c' variable reaches nil
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		list = traverse(c, list)
		// Since slices are pass by reference types, we're passing the same string slice over and over again so that the function
		// keeps appending this slice
	}

	return list

}
