package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "serarator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
	var i = 1
	i++
	var x = i
	println(x)

	fmt.Println(gcd(12, 15))
}

/*
*
求最大公约数
*/
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
