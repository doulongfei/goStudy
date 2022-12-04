package main

import "fmt"

func incr(p *int) int {
	*p++
	fmt.Println(*p)
	return *p
}

func main() {
	v := 1
	incr(&v)
	incr(&v)
	fmt.Println(v)

	x := 1
	p := &x //代表某个变量的地址
	fmt.Println("值", *p)
	fmt.Println("地址", p)
	*p = 2 //*xxx 是某个变量，可以直接对其进行操作（赋值）相当于变量
	// 问：*xxx 和变量 x的区别是什么？？？
	//每次我们对一个变量取地址，或者复制指针，我们都是为原变量创建了新的别名。例如，*p就是变量v的别名。
	//指针特别有价值的地方在于我们可以不用名字而访问一个变量，
	fmt.Println(x)

	p1 := f1()
	fmt.Println(p1 == f1())
}

func f1() *int {
	v := 1
	return &v
}
