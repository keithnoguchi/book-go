// A web link visitor
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	result := make(chan []string)
	for _, url := range os.Args[1:] {
		go visiter(result, url)
	}
	for range os.Args[1:] {
		links := <- result
		for i, link := range links {
			fmt.Println(i, link)
		}
	}
}

func visiter(out chan<- []string, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("get(%q): %v", url, err)
		out <- nil
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Printf("parse(%q): %v", url, err)
		out <- nil
	}
	out <- visit(nil, doc)
}

func visit(link []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				link = append(link, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		link = visit(link, c)
	}
	return link
}
