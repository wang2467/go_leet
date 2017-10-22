package main

import (
	"fmt"
)

func exist(board [][]int, word []int) bool {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if word[0] == board[i][j] {
				if dfs(board, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func dfs(board [][]int, word []int, a int, i int, j int) bool {
	m := len(board)
	n := len(board[0])
	var result bool
	if a == len(word) {
		return true
	}
	if i < 0 || j < 0 || i > m-1 || j > n-1 {
		return false
	}
	if board[i][j] == word[a] {
		temp := board[i][j]
		board[i][j] = -1
		result = dfs(board, word, a+1, i-1, j) || dfs(board, word, a+1, i+1, j) || dfs(board, word, a+1, i, j-1) || dfs(board, word, a+1, i, j+1)
		board[i][j] = temp
	}
	return result
}

func main() {
	board := [][]int{
		{1, 2, 3, 5},
		{24, 6, 3, 24},
		{1, 4, 5, 5},
	}
	word := []int{1, 2, 3, 3, 5, 4, 1}
	fmt.Println(exist(board, word))
}
