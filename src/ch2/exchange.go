package main

import "fmt"

func main() {
	var a, b = 2, 3
	a, b = b, a
	println(a, b)

	x := 1
	p := &x //代表某个变量的地址
	fmt.Println("值", *p)
	fmt.Println("地址", p)
	*p = 2 //*xxx 是某个变量，可以直接对其进行操作（赋值）相当于变量
	// 问：*xxx 和变量 x的区别是什么？？？
	fmt.Println(x)

	p1 := f()
	fmt.Println(p1 == f())
}

func f() *int {
	v := 1
	return &v
}
