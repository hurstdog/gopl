// Server3 is a minimal "echo" and counter server, plus more headers.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

var palette = []color.Color{
	color.Black,
	color.RGBA{0xff, 0x00, 0x0, 0xff}, // red
	color.RGBA{0x0, 0xff, 0x0, 0xff},  // green
	color.RGBA{0x0, 0x0, 0x80, 0xff},  // blue
}

const (
	backgroundIndex = 0 // first color in palette
)

const minColorIndex = 1
const maxColorIndex = 3

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echos the Path component of the requested URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	// Count isn't in the original, but leaving it here for fun.
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echos the number of calls so far
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer, r *http.Request) {
	cycles := 5.0 // number of complete x oscillator revolutions
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	if r.Form["cycles"] != nil {
		formCycles, err := strconv.Atoi(r.Form["cycles"][0])
		if err != nil {
			fmt.Fprintf(out, "Error reading form(%q): %v", r.Form, err)
		}
		cycles = float64(formCycles)
	}
	freq := rand.Float64() * 3.0 // relative frequncy of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	colorCount := 0
	// Change the color every 10 frames
	colorIndex := minColorIndex
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(colorIndex))
		}
		colorCount++
		if colorCount >= 10 {
			colorCount = 0
			colorIndex++
			if colorIndex > maxColorIndex {
				colorIndex = minColorIndex
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: Ignoring encoding errors
}
