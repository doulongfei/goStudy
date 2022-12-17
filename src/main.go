package main

import (
	"fmt"
)

func main() {
	fmt.Println("helle everyone!!!")
	ints := map[string]int64{
		"first":  32,
		"second": 12,
	}

	floats := map[string]float64{
		"first":  32.21,
		"second": 12.23,
	}
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
