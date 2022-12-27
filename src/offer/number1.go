package offer

func NumberOf1(n int) int {
	// write code here
	res := 0
	for i := 0; i < 32; i++ {
		if n&(1<<i) != 0 {
			res++
		}
	}
	return res
}

func Power(base float64, exponent int) float64 {
	// write code here
	if exponent < 0 {
		base = 1 / base
		exponent = -exponent
	}
	res := 1.0
	//累乘
	for i := 0; i < exponent; i++ {
		res *= base
	}
	return res
}
func printNumbers(n int) []int {
	// write code here
	end := 1
	for i := 1; i < n; i++ {
		end *= 10
	}
	res := make([]int, end-1)
	for i := 1; i < end; i++ {
		res[i-1] = i
	}
	return res
}
