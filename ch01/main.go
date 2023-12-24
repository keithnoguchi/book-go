// A Lissajous web server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", lissajous)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!\n")
}

const (
	whiteIndex = 0
	blackIndex = 1
)

var palette = []color.Color{color.White, color.Black}

func lissajous(w http.ResponseWriter, r *http.Request) {
	const (
		cycles = 10
		res = 0.001
		size = 200
		nframes = 64
		delay = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				blackIndex,
			)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, &anim)
}
