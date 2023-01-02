package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find(target int, array [][]int) bool {
	// write code here
	// 右上角开始
	// 有多少列 y
	column := len(array[0])
	// 有多少行  x
	row := len(array)

	for x, y := 0, column-1; x < row && y >= 0; {
		if array[x][y] < target {
			x++
			continue
		}
		if array[x][y] > target {
			y--
			continue
		}
		if array[x][y] == target {
			return true
		}
	}
	return false

	// 左下角开始
	for x, y := row-1, 0; x >= 0 && y < column; {
		if array[x][y] < target {
			y++
			continue
		}
		if array[x][y] > target {
			x--
			continue
		}
		if array[x][y] == target {
			return true
		}
	}
	return false
}
