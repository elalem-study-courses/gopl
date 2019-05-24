package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
	valleyColor   = 0x0000ff
	peakColor     = 0xff0000
	colorRange    = peakColor - valleyColor
	maxHeight     = 310.0
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			if isFinite(ax + ay + bx + by + cx + cy + dx + dy) {
				color := getColor(ay, by, cy, dy)
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: #%x; stroke: #%[9]x'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}
	}
	fmt.Println("</svg>")
}

func getColor(a, b, c, d float64) int {
	max := math.Max(a, b)
	max = math.Max(max, c)
	max = math.Max(max, d)

	color := int(valleyColor + (max/maxHeight)*colorRange)

	return color
}

func isFinite(x float64) bool {
	sign := 1
	if x < 0 {
		sign = -1
	}
	isNan := math.IsNaN(x)
	isInf := math.IsInf(x, sign)
	return !isInf && !isNan
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
