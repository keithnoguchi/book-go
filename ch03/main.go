// A Mandelbrot set web server
package main

import (
	"context"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"sync"
	"time"
)

type Mandelbrot struct{}

func main() {
	var wg sync.WaitGroup
	s := http.Server{
		Addr:    "localhost:8080",
		Handler: Mandelbrot{},
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

func (h Mandelbrot) ServeHTTP(w http.ResponseWriter, _r *http.Request) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
