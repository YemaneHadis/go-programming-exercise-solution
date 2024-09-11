// package main

// import (
// 	"fmt"
// 	"log"
// 	"math"
// 	"os"
// )

// const (
// 	width, height = 600, 320            // canvas size in pixels
// 	cells         = 100                 // number of grid cells
// 	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
// 	xyscale       = width / 2 / xyrange // pixels per x or y unit
// 	zscale        = height * 0.4        // pixels per z unit
// 	angle         = math.Pi / 6         // angle of x, y axes (=30°)
// )

// var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

// func main() {
// 	out, err := os.Create("plot.svg")
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}

// 	//fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
// 	// "style='stroke: grey; fill: white; stroke-width: 0.7' "+
// 	// "width='%d' height='%d'>", width, height)
// 	s := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
// 		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
// 		"width='%d' height='%d'>", width, height)
// 	out.WriteString(s)
// 	for i := 0; i < cells; i++ {
// 		for j := 0; j < cells; j++ {
// 			ax, ay := corner(i+1, j)
// 			bx, by := corner(i, j)
// 			cx, cy := corner(i, j+1)
// 			dx, dy := corner(i+1, j+1)
// 			if vaidateNumber(ax) || vaidateNumber(ay) || vaidateNumber(bx) || vaidateNumber(by) || vaidateNumber(cx) || vaidateNumber(cy) || vaidateNumber(dx) || vaidateNumber(dy) {
// 				continue
// 			} else {
// 				s = fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
// 					ax, ay, bx, by, cx, cy, dx, dy)
// 				out.WriteString(s)
// 			}
// 		}
// 	}
// 	out.WriteString("</svg>")
// 	defer out.Close()
// }

// func corner(i, j int) (float64, float64) {
// 	// Find point (x,y) at corner of cell (i,j).
// 	x := xyrange * (float64(i)/cells - 0.5)
// 	y := xyrange * (float64(j)/cells - 0.5)

// 	// Compute surface height z.
// 	z := eggbox(x, y)

// 	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
// 	sx := width/2 + (x-y)*cos30*xyscale
// 	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
// 	return sx, sy
// }
// func vaidateNumber(num float64) bool {
// 	return math.IsInf(num, 0) || math.IsNaN(num)
// }

// func f(x, y float64) float64 {
// 	r := math.Hypot(x, y) // distance from (0,0)
// 	return math.Sin(r) / r
// }
