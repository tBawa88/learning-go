package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

/*
The strings.NewReader function returns a value that satisfies the io.Reader
interface (and others) by reading from its argument, a string. Implement a simple version of
NewReader yourself, and use it to make the HTML parser (§5.2) take input from a string.
*/
// NewReader(s string) *Reader
// Reader is something from which data can be read, one of the purposes of Reader is that it can be used to write to a Writer type
// So we have to create a new type that has some string property on it, and satisfies the io.Reader interface (meaning has a Read(b []byte)(int, error)) function as it's method

type StringReader struct {
	str string
	pos int //current reading position
}

// The entire job of a Read(b []byte) method is to write the string data to the byte slice passed to it
// When we call some method like io.Copy, it creates it's own byte slice
// passes it to Read()
// then passes the same slice to Write() so that it can be written somewhere

func (sReader *StringReader) Read(buff []byte) (n int, err error) {
	// check if we've already read all the data
	if sReader.pos >= len(sReader.str) {
		return 0, io.EOF // EOF error
	}

	// read the data from str to buffer, but start copying from where the current pos is
	n = copy(buff, sReader.str[sReader.pos:])
	sReader.pos += n // move the pos to n places further
	return n, nil
}

// An initiator function to return a new object of type StringReader
func NewReader(s string) *StringReader {
	return &StringReader{str: s, pos: 0}
}

func main() {
	input := `<html>
			<body><a href="https://example.com">Example</a>
			<a herf="https://google.com">Google it</a>
			<a href="https://instagram.com">BrainRot</a>
			</body>
		</html>`

	reader := NewReader(input)
	node, err := html.Parse(reader)

	if err != nil {
		fmt.Println("Error parsing the html ", err)
		os.Exit(1)
	}

	links := visit(node, make([]string, 0))
	for _, val := range links {
		fmt.Printf("%s ||", val)
	}
	fmt.Println()
}

// recursively visit each node and extract only the link tag values
func visit(node *html.Node, links []string) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(c, links)
	}
	return links
}
