// A web crawler
//
// # Example
//
// 111
// go run . https://golang.org
//```
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	breathFirst(crawl, os.Args[1:])
}

func crawl(url string) []string {
	fmt.Println(url)
	links, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return links
}

func breathFirst(f func(string) []string, work []string) {
	seen := make(map[string]bool)
	for len(work) > 0 {
		items := work
		work = nil
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			work = append(work, f(item)...)
		}
	}
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", url, err)
	}
	var links []string
	visit := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEach(doc, visit, nil)
	return links, nil
}

func forEach(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEach(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
