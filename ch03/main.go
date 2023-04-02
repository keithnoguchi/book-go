// A Mandelbrot set web server
package main

import (
	"context"
	"fmt"
	_ "image"
	_ "image/color"
	_ "image/png"
	"log"
	_ "math/cmplx"
	"net/http"
	_ "os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	s := http.Server{
		Addr:    "localhost:8080",
		Handler: mandelbrot {},
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	time.Sleep(5 * time.Second)
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}

type mandelbrot struct {}

func (h mandelbrot) ServeHTTP(w http.ResponseWriter, _r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
