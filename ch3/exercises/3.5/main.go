// package description
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func main() {
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

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 350
	const contrast = 15

	var v complex128
	for n := int(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			hue := float64(255 * n / iterations)
			saturation := 255.0
			value := 255.0
			return colorful.Hsl(hue, saturation, value)
		}
	}

	return color.Black
}

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}
