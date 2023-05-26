// fetch fetches the web contents
package main

import (
	"fmt"
	"io/ioutil"
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

// fetch fetches the html body for the given url returns over the
// ch channel.
//
// Please note that it returns the error over the channel as a string,
// which will be addressed in the future iteration.
func fetch(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- fmt.Sprint(err)
	}
	ch <- string(b)
}
