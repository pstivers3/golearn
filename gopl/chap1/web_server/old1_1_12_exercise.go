// serve lissajous figures
// doesn't work!!

package main

import (
	"image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
	"net/http"
	"log"
)

//i, err := strconv.Atoi("-42")

var palette = []color.Color{color.White, color.Black} //compostite literal slice

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	    lissajous(w)
    })
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// http.HandleFunc("/", liss_param)
}

/*
func liss_param(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}*/

func lissajous(out io.Writer) {
	const (
        cycles  = strconv.Atoi(r.URL.Path)     // number of complete x oscilator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 64    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )
    freq := rand.Float64() * 3.0 // relative frequency and y oscillator
    anim := gif.GIF{LoopCount: nframes} // composite literal struct
    phase := 0.0 // phase difference
    for i := 0; i < nframes; i++ {   // each iteration produces a single frame for the gif
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t< cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size +int(x*size+0.5), size+int(y*size+0.5),
                blackIndex)
        }
        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
	gif.EncodeAll(out, &anim) // Note: ignoring encoding errors
}

