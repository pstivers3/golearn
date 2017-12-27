// Lissajous generates GIF animations of random Lissajous figures.
// to gererate a gif:
// $ go build filename.go
// $ ./filename > gif_filename.gif
// $ open gif_filename.gif

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

// RGBA green = 0 256 0 1 in binary == 0 0 80 1 in hex 
var palette = []color.Color{color.Black, color.RGBA{0x0, 0x80, 0x0, 0x1}} //compostite literal slice

const (
    whiteIndex = 0 // first color in palette
    blackIndex = 1 // next color in palette
)

func main() {
    lissajous (os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5     // number of complete x oscilator revolutions
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

