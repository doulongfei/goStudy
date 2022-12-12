package main

/*
*
解法一 ： 哈希表
*/
func duplicate(numbers []int) int {
	m := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		_, b := m[numbers[i]]
		if b {
			return numbers[i]
		} else {
			m[numbers[i]] = numbers[i]
		}
	}
	return -1
}

/*
*
解法二：在一个长度为n的数组里的所有数字都在0到n-1的范围内。
数组中某些数字是重复的，但不知道有几个数字是重复的。也不知道每个数字重复几次。请找出数组中任意一个重复的数字。 例如，如果输入长度为7的数组[2,3,1,0,2,5,3]
，那么对应的输出是2或者3。存在不合法的输入的话输出-1

所有数字都在0到n-1的范围内。所以排序
*/
func duplicate2(numbers []int) int {

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == i {
			continue
		} else {
			if numbers[i] == numbers[numbers[i]] {
				return numbers[i]
			} else {
				numbers[i], numbers[numbers[i]] = numbers[numbers[i]], numbers[i]
				i--
			}
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 1, 0, 2, 5, 3}
	j := duplicate2(arr)
	//i := duplicate2(arr)
	//println(i)
	println(j)
}
