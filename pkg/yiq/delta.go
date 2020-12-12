package yiq

import (
	"image/color"
)

func normalize(rgba color.Color) (uint8, uint8, uint8, uint8) {
	r, g, b, a := rgba.RGBA()

	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func rgb2y(r, g, b uint8) float64 {
	return float64(r)*0.29889531 + float64(g)*0.58662247 + float64(b)*0.11448223
}

func rgb2i(r, g, b uint8) float64 {
	return float64(r)*0.59597799 - float64(g)*0.27417610 - float64(b)*0.32180189
}

func rgb2q(r, g, b uint8) float64 {
	return float64(r)*0.21147017 - float64(g)*0.52261711 + float64(b)*0.31114694
}

// Delta between two pixels.
func Delta(pixelA, pixelB color.Color) float64 {
	r1, g1, b1, _ := normalize(pixelA)
	r2, g2, b2, _ := normalize(pixelB)

	y := rgb2y(r1, g1, b1) - rgb2y(r2, g2, b2)
	i := rgb2i(r1, g1, b1) - rgb2i(r2, g2, b2)
	q := rgb2q(r1, g1, b1) - rgb2q(r2, g2, b2)

	return 0.5053*y*y + 0.299*i*i + 0.1957*q*q
}

// MaxDelta is a max value of Delta func.
var MaxDelta = 35215.0
