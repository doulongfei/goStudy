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
}
