package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"strconv"
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

func test39(w http.ResponseWriter, r *http.Request) {
	// 从HTTP参数中获取长、宽、放大倍数
	err := r.ParseForm()
	if err != nil {
		return
	}
	width, height, zoom := 1024, 1024, 1
	if s := r.Form.Get("x"); s != "" {

		width, _ = strconv.Atoi(s)
	}
	if s := r.Form.Get("y"); s != "" {

		height, _ = strconv.Atoi(s)
	}
	if s := r.Form.Get("zoom"); s != "" {

		zoom, _ = strconv.Atoi(s)
	}
	Fractal(w, width*zoom, height*zoom)
}

func main() {
	http.HandleFunc("/", test39)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println("server listen error,err:", err)
	}
}
