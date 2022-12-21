package offer

func minNumberInRotateArray(rotateArray []int) int {
	left, right := 0, len(rotateArray)-1
	for left < right {
		middle := (right + left) >> 1 // >> 1 代表除以2
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

func search(nums []int, target int) int {
	left1, right1 := 0, len(nums)-1
	for left1 < right1 {
		mid := (left1 + right1) >> 1
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 说明在右区间
			left1 = mid + 1
		} else {
			right1 = mid - 1
		}
	}
	return -1
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @return int整型
 */
func findPeakElement(nums []int) int {
	// write code here
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		//右边是往下，不一定有坡峰
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			//右边是往上，一定能找到波峰
			left = mid + 1
		}
	}
	return right
}
