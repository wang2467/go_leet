package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	countPath(mat, m-1, n-1)
	return mat[m-1][n-1]
}

func countPath(mat [][]int, m int, n int) int {
	if m == 0 && n == 0 {
		mat[m][n] = 1
		return mat[m][n]
	}
	if m == 0 {
		mat[m][n] = countPath(mat, m, n-1)
	} else if n == 0 {
		mat[m][n] = countPath(mat, m-1, n)
	} else {
		mat[m][n] = countPath(mat, m, n-1) + countPath(mat, m-1, n)
	}
	return mat[m][n]
}

func uniquePathsIterative(m int, n int) int {
	mat := make([][]int, m)
	for i := 0; i < m; i++ {
		mat[i] = make([]int, n)
	}
	return countPathIterative(mat, m, n)
}

func countPathIterative(mat [][]int, m int, n int) int {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				mat[i][j] = 1
			} else if i == 0 {
				mat[i][j] = mat[i][j-1]
			} else if j == 0 {
				mat[i][j] = mat[i-1][j]
			} else {
				mat[i][j] = mat[i-1][j] + mat[i][j-1]
			}
		}
	}
	return mat[m-1][n-1]
}

func main() {
	//fmt.Println(uniquePaths(3, 3))
	fmt.Println(uniquePathsIterative(2, 2))
}
