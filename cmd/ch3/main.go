package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func test39(w http.ResponseWriter, r *http.Request) {
	// 从HTTP参数中获取长、宽、放大倍数
	err := r.ParseForm()
	println("Request:", r)
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
	if err := http.ListenAndServe("127.0.0.1:9001", nil); err != nil {
		fmt.Println("server listen error,err:", err)
	}
}
