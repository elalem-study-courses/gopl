package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

type PointPosition int

const (
	width, height               = 600, 320
	cells                       = 100
	xyrange                     = 30.0
	xyscale                     = width / 2 / xyrange
	zscale                      = height * 4.0
	angle                       = math.Pi / 6
	peak          PointPosition = 0
	valley        PointPosition = 1
	middle        PointPosition = 2
)

var sin30, cos30 = math.Y0(angle), math.J0(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		fmt.Fprintln(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
			"style='stroke: grey; fill: white; stroke-width: 0.7' "+
			"width='%d' height='%d'>", width, height)

		for i := 0; i < cells; i++ {
			for j := 0; j < cells; j++ {
				ax, ay, pt1 := corner(i+1, j)
				bx, by, pt2 := corner(i, j)
				cx, cy, pt3 := corner(i, j+1)
				dx, dy, pt4 := corner(i+1, j+1)

				var color string

				if checkHeight(pt1, pt2, pt3, pt4, peak) {
					color = "#ff0000"
				} else if checkHeight(pt1, pt2, pt3, pt4, valley) {
					color = "#0000ff"
				} else {
					color = "#00ff00"
				}
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='%s' />\n", ax, ay, bx, by, cx, cy, dx, dy, color)
			}
		}

		fmt.Fprintln(w, "</svg>")
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func corner(i, j int) (float64, float64, PointPosition) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	pt := getPointType(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, pt
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func getPointType(x, y float64) PointPosition {
	d := math.Hypot(x, y)
	pt := middle

	if math.Abs(d-math.Tan(d)) < 3 {
		pt = peak
		if 2*(math.Sin(d)-d*math.Cos(d))-d*d*math.Sin(d) > 0 {
			pt = valley
		}
	}

	return pt
}

func checkHeight(h1, h2, h3, h4, h PointPosition) bool {
	switch h {
	case h1, h2, h3, h4:
		return true
	}

	return false
}
