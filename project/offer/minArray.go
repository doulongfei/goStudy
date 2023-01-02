package main

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

// 递归
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

// 先排除 数组长度为1，为2 这种情况
// 长度大于等于3后，从索引1处开始遍历，取索引-1 ，索引+1 与当前比较，都比当前小，返回
func findPeakElement2(nums []int) int {
	// write code here
	length := len(nums)
	if length == 0 || length == 1 {
		return 0
	}
	if nums[length-1] > nums[length-2] {
		return length - 1
	}
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return 0
}
