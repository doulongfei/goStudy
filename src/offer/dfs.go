package main

import (
	"fmt"
)

// JZ12 矩阵中的路径
/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param matrix char字符型二维数组
 * @param word string字符串
 * @return bool布尔型
 */
// func hasPath(matrix [][]byte, word string) bool {
// 	// write code here
// 	wordArr:=[]string{word}

// }

func main() {
	wordArr := []byte("abc")

	fmt.Println(string(wordArr[0]), len(wordArr))

	s := "hello, 世界"
	r := []rune(s)
	b := []byte(s)

	fmt.Println(r)
	fmt.Println(b)

}

func hasPath(matrix [][]byte, word string) bool {
	// 哈希表，用于记录已访问的格子
	visited := make(map[int]map[int]bool)

	// 遍历所有格子
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// 从当前格子开始搜索
			if dfs(matrix, visited, i, j, 0, word) {
				return true
			}
		}
	}

	return false
}

func dfs(matrix [][]byte, visited map[int]map[int]bool, i, j, index int, word string) bool {
	// 如果已经找到了所有字符，则返回 true
	if index == len(word) {
		return true
	}

	// 如果越界或者当前格子的字符不匹配，则返回 false
	if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || matrix[i][j] != word[index] {
		return false
	}

	// 如果当前格子已经访问过，则返回 false
	if _, ok := visited[i]; !ok {
		visited[i] = make(map[int]bool)
	}
	if visited[i][j] {
		return false
	}

	// 标记当前格子已访问
	visited[i][j] = true

	// 向四个方向搜索
	if dfs(matrix, visited, i-1, j, index+1, word) ||
		dfs(matrix, visited, i+1, j, index+1, word) ||
		dfs(matrix, visited, i, j-1, index+1, word) ||
		dfs(matrix, visited, i, j+1, index+1, word) {
		return true
	}

	// 清除标记，回溯
	visited[i][j] = false

	return false
}
