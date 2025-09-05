package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// main downloads an HTML document from the provided URL and
// prints all hyperlinks found in the page.
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL\n", os.Args[0])
		os.Exit(1)
	}

	resp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatalf("failed to fetch URL: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("failed to parse HTML: %v", err)
	}

	// Recursively traverse the HTML node tree and print href attributes.
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					fmt.Println(attr.Val)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
