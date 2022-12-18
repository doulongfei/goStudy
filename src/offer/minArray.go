package offer

func minNumberInRotateArray(rotateArray []int) int {
	left, right := 0, len(rotateArray)
	for left < right {
		middle := left + (right-left)>>1
		if rotateArray[middle] > rotateArray[right] {
			left = middle + 1
		} else if rotateArray[middle] < rotateArray[right] {
			right = middle
		} else {
			right--
		}
	}
	return rotateArray[left]
}
