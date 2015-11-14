// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

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
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
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
