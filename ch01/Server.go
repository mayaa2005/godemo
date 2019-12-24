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
)

var palette2 = []color.Color{color.White, color.Black, color.RGBA{
	R: 0,
	G: 0,
	B: 255,
	A: 128,
}}

func main() {
	http.HandleFunc("/image", handleImage)
	http.HandleFunc("/", handleCommon)
	http.ListenAndServe(":http", nil)
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("parse form error")
	}
	color := r.Form.Get("color")
	colorIndex, err := strconv.Atoi(color)
	if err != nil {
		log.Print("color convert error")
		colorIndex = 1
	}
	lissajous2(uint8(colorIndex), w)
}

func handleCommon(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Url:%s,Method:%s\n", r.RequestURI, r.Method)
	header := r.Header
	for k, v := range header {
		fmt.Fprintf(w, "%s:%q\n", k, v)
	}

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form:[%q]:%q", k, v)
	}
}

func lissajous2(colorIndex uint8, out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{
		Image:           nil,
		Delay:           nil,
		LoopCount:       nframes,
		Disposal:        nil,
		Config:          image.Config{},
		BackgroundIndex: 0,
	}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
