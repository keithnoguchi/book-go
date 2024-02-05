// A web link visitor
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(len(os.Args[1:]))
	for _, url := range os.Args[1:] {
		go visitor(&wg, url)
	}
	wg.Wait()
}

func visitor(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "visitor: %v\n", err)
		return
	}
	defer resp.Body.Close()
	n, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "visitor: %v\n", err)
	}
	fmt.Println(url, n)
}
