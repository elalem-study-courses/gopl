// package description
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		subPixels              = 2
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = subPixels * 1024, subPixels * 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py += subPixels {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px += subPixels {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
			if px > 0 {
				avg := (colorToInt(img.At(px, py)) + colorToInt(img.At(px-subPixels, py))) / 2
				avgColor := intToColor(avg)
				for i := px - subPixels + 1; i < px; i++ {
					img.Set(i, py, avgColor)
				}
			}
			if py > 0 {
				avg := (colorToInt(img.At(px, py)) + colorToInt(img.At(px, py-subPixels))) / 2
				avgColor := intToColor(avg)
				for i := py - subPixels + 1; i < py; i++ {
					img.Set(px, i, avgColor)
				}
			}
			if px > 0 && py > 0 {
				avg := (colorToInt(img.At(px, py)) +
					colorToInt(img.At(px, py-subPixels)) +
					colorToInt(img.At(px-subPixels, py)) +
					colorToInt(img.At(px-subPixels, py-subPixels))) / 4
				avgColor := intToColor(avg)

				for i := px - subPixels + 1; i < px; i++ {
					for j := py - subPixels + 1; j < py; j++ {
						img.Set(i, j, avgColor)
					}
				}
			}
		}
	}

	png.Encode(os.Stdout, img)
}

func colorToInt(c color.Color) int {
	r, g, b, a := c.RGBA()
	return int((r << 24) + (g << 16) + (b << 8) + a)
}

func intToColor(c int) color.Color {
	r, g, b, a := (c>>24)&0xff, (c>>16)&0xff, (c>>8)&0xff, c&0xff
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
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
