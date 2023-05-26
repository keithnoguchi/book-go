// A web link lister.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	ch := make(chan io.Reader)
	for _, url := range os.Args[1:] {
		go reader(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

// reader retrieves the url website and returns the io.Reader
// over the channel.
func reader(url string, ch chan<- io.Reader) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get(%s) error: %s\n", url, err)
		return
	}
	ch <- resp.Body
}
