// A surface plot of the function sin(r)/r
//
// # Examples
//
// The following command will start the web server
// showing the sin(r)/r drawing on port 8000:
// ```
// $ go run main.go
// ```
package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"sync"
	"time"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cons30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", plot)
	s := &http.Server{Addr: "localhost:8000"}

	var wg sync.WaitGroup
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

func plot(w http.ResponseWriter, _r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(
		w,
		"<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>",
		width,
		height,
	)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(
				w,
				"<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy,
			)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cons30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
