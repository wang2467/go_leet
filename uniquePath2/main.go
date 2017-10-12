package main

import (
	"fmt"
)

func uniquePathWithObstacles(obstacleGrid [][]int) int {
	return countPath(obstacleGrid)
}

func countPath(mat [][]int) int {
	m := len(mat)
	n := len(mat[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				if mat[0][0] == 1 {
					res[0][0] = 0
				} else {
					res[0][0] = 1
				}
			} else if i == 0 {
				if mat[i][j] == 1 {
					res[i][j] = 0
				} else {
					res[i][j] = res[i][j-1]
				}
			} else if j == 0 {
				if mat[i][j] == 1 {
					res[i][j] = 0
				} else {
					res[i][j] = res[i-1][j]
				}
			} else {
				if mat[i][j] == 1 {
					res[i][j] = 0
				} else if mat[i-1][j] == 1 {
					res[i][j] = res[i][j-1]
				} else if mat[i][j-1] == 1 {
					res[i][j] = res[i-1][j]
				} else {
					res[i][j] = res[i-1][j] + res[i][j-1]
				}
			}
		}
	}
	return res[m-1][n-1]
}

func main() {
	mat := make([][]int, 2)
	for i := 0; i < 2; i++ {
		mat[i] = make([]int, 2)
	}
	mat[1][1] = 1
	fmt.Println(uniquePathWithObstacles(mat))
}
