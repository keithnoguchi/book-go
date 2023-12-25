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
		go getter(url, ch)
	}
	for range os.Args[1:] {
		resp := <-ch
		fmt.Println(resp)
		resp.Close()
	}
}

// getter gets the url website and returns the io.ReadCloser
// over the channel.
//
// The io.ReadCloser needs to be closed by the consumer.
func getter(url string, ch chan<- io.ReadCloser) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http.Get(%s) error: %s\n", url, err)
		return
	}
	ch <- resp.Body
}
