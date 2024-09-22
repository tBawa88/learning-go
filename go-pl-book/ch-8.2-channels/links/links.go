package links

import (
	"net/http"

	"golang.org/x/net/html"
)

// Making a pakckage that contains a function called Extract, which makes a HTTP GET request to the given link
// parses the received HTML  and extracts all the links in that response

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rootNode, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	var visitnode func(n *html.Node)
	visitnode = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			// itreate over it's attributes and obtain the href values
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(attr.Val) // URL.Prase() parses the provided url with respect to the response URL
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			// recursively call this function for each child node
			visitnode(child)
		}
	}
	visitnode(rootNode)
	return links, nil
}
