package main

//func cutRope(number int) int {
//	// write code here
//	if number <= 3 {
//		return number - 1
//	}
//
//}

func cutRopeShuXue(number int) int {
	// write code here
	if number%3 == 0 {
		return power(3, number/3)
	}
	if number%3 == 1 {
		return power(3, (number-1)/3-1) * 4
	}
	if number%3 == 2 {
		return power(3, (number-2)/3) * 2
	}
	return 0
}

func power(num int, n int) int {
	if n == 0 {
		return 1
	} else {
		return num * power(num, n-1)
	}
}
