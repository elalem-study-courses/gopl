// package description
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		widthParam, _ := strconv.ParseInt(r.FormValue("width"), 10, 64)
		heightParam, _ := strconv.ParseInt(r.FormValue("height"), 10, 64)

		width := int(widthParam)
		height := int(heightParam)
		xmin, _ := strconv.ParseFloat(r.FormValue("xmin"), 64)
		ymin, _ := strconv.ParseFloat(r.FormValue("ymin"), 64)
		xmax, _ := strconv.ParseFloat(r.FormValue("xmax"), 64)
		ymax, _ := strconv.ParseFloat(r.FormValue("ymax"), 64)
		// const (
		// 	xmin, ymin, xmax, ymax = -2, -2, +2, +2
		// 	width, height          = 1024, 1024
		// )

		img := image.NewRGBA(image.Rect(0, 0, width, height))
		for py := int(0); py < height; py++ {
			y := float64(py)/float64(height)*(ymax-ymin) + ymin
			for px := int(0); px < width; px++ {
				x := float64(px)/float64(width)*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}

		png.Encode(w, img)
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
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
