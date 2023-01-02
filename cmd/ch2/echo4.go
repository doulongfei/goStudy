package ch2

import (
	"flag"
	"fmt"
	"strings"
	"utils"
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

	fmt.Println(utils.Gcd(12, 15))
}
