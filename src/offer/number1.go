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
