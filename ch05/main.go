// A web link finder.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go finder(url, ch)
	}
	for range os.Args[1:] {
		fmt.Print(<-ch)
	}
}

func finder(url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get() error: %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	_, err = html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "html.Parse() error: %s\n", err)
		os.Exit(1)
	}
	ch <- "not implemented\n"
}
