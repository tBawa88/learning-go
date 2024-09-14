package main

// Goal of this program is to recursivel parse a given HTML file
// IT uses a HTML parser package which is outside of go's std lib
// We will traverse over a given HTML and extract all the links from it

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// Parse() function returns parses the entire HTML tree given to it, and returns the ROOT node of that tree
	htmlFile, err := os.Open("test.html")
	if err != nil {
		log.Fatal("Error opening the file ", err)
	}
	// rootNode, err := html.Parse(os.Stdin)
	rootNode, err := html.Parse(htmlFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks 1:%v\n", err)
		os.Exit(1)
	}

	// we're using this root node to recursively visit every node, and extract all the link elements
	for _, link := range visit(make([]string, 0), rootNode) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			// iterating over all the attributes of node n to find the 'href' attr, since it's confirmed to be <a>
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// Here is the recursive call
	// To descend the tree of node n, visit() function is recrusively calling itself and moving to the nextsibling on each call
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
