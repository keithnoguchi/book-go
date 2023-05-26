// fetch fetches the web contents
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan string) {
	_, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintln(err)
	}
	ch <- url
}
