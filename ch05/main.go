// A web link lister.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	ch := make(chan io.ReadCloser)
	for _, url := range os.Args[1:] {
		go reader(url, ch)
	}
	for range os.Args[1:] {
		resp := <- ch
		fmt.Println(resp)
		resp.Close()
	}
}

// reader retrieves the url website and returns the io.Reader
// over the channel.
//
// A receiver needs to close the io.Reader
func reader(url string, ch chan<- io.ReadCloser) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get(%s) error: %s\n", url, err)
		return
	}
	ch <- resp.Body
}
