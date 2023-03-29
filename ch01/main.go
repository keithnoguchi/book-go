// A web server
//
// # Examples
// ```
// $ go run main.go
// 2023/03/28 18:03:29 counter
// 2023/03/28 18:03:29 handler
// 2023/03/28 18:03:30 counter
// 2023/03/28 18:03:30 handler
// ```
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var lock sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler")
	lock.Lock()
	count++
	lock.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	log.Println("counter")
	lock.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	lock.Unlock()
}
