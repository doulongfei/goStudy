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

func minNumberInRotateArray2(rotateArray []int) int {
	size := len(rotateArray)
	if size == 1 {
		return rotateArray[0]
	}
	return Min(minNumberInRotateArray2(rotateArray[0:size/2]), minNumberInRotateArray2(rotateArray[(size/2):size]))
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
