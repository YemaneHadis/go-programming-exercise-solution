package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{
	color.RGBA{R: 255, G: 255, B: 255, A: 255}, // Red
	color.RGBA{R: 0, G: 255, B: 0, A: 255},     // Green
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

//handler echoes the path components of the request URL
// func handler(w http.ResponseWriter, r *http.Request) {
// 	mu.Lock()
// 	count++
// 	mu.Unlock()
// 	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

// }

func handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	cycle := queryParams.Get("cycle")
	if cycle == "" {
		lissajous(w, 5)
	} else {
		val, _ := strconv.Atoi(cycle)
		lissajous(w, val)
	}

}

// handler := func(w http:ResponseWriter, r *http.Request){

// }

// counter echoes the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprint(w, "Count %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, cycle int) {
	const (
		// cycle   = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolutions
		size    = 100   // image canvas covers[-size .. +size]
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0 //relative frequescies of y osicllatoir
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycle)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
