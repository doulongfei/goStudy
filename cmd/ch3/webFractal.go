package main

import (
	"image"
	"image/png"
	"io"
)

func Fractal(w io.Writer, width int, height int) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin

		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandel(z))
		}
	}
	err := png.Encode(w, img)
	if err != nil {
		return
	}
}
