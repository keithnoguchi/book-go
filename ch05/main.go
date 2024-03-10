// A web crawler
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	bfs(crawl, os.Args[1:])
}

func bfs(f func(string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if seen[item] {
				continue
			}
			seen[item] = true
			worklist = append(worklist, f(item)...)
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", url, err)
	}
	var links []string
	visit := func(n *html.Node) {
		if n.Type != html.ElementNode || n.Data != "a" {
			return
		}
		for _, a := range n.Attr {
			if a.Key == "href" {
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					log.Printf("%s/%s: %v", url, a.Val, err)
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visit, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
