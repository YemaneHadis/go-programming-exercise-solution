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
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			subpixelColors := make([]color.Color, 0)
			for sub_py := 0; sub_py <= 1; sub_py++ {
				for sub_px := 0; sub_px <= 1; sub_px++ {

					x := float64(px)/width*(xmax-xmin) + xmin
					z := complex(x, y)
					color := mandelbrot(z)
					subpixelColors = append(subpixelColors, color)
				}
			}

			// Image point (px, py) represents complex value of z
			img.Set(px, py, averageColor(subpixelColors))
		}
	}
	png.Encode(os.Stdout, img) // note : ignoring error
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		blue := uint8(255 - contrast*n)
		red := uint8(255 - blue)
		green := uint8(123)

		if cmplx.Abs(v) > 2 {
			return color.RGBA{red, green, blue, 255}
		}
	}
	return color.Black
}

func averageColor(colors []color.Color) color.Color {
	var averageRed uint32 = 0
	var averageGreen uint32 = 0
	var averageBlue uint32 = 0
	var averageAlpha uint32 = 0

	for _, color := range colors {
		red, green, blue, alpha := color.RGBA()
		averageRed += uint32(uint8(red))
		averageGreen += uint32(uint8(green))
		averageBlue += uint32(uint8(blue))
		averageAlpha += uint32(uint8(alpha))

	}
	return color.RGBA{
		uint8(averageRed / uint32(len(colors))),
		uint8(averageGreen / uint32(len(colors))),
		uint8(averageBlue / uint32(len(colors))),
		uint8(averageAlpha / uint32(len(colors))),
	}
}
