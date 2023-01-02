package utils

// Gcd 求最大公约数
func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// Fib 求斐波那契数列的第n个数
func Fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
