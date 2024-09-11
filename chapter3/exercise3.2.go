package main

import "math"

func eggbox(x, y float64) float64 {
	return 0.1 * (math.Cos(x) + math.Cos(y))
}
