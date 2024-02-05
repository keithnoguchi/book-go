// A web link finder
package main

import (
	"fmt"
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
	fmt.Println(url)
}
