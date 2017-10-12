package main

import (
	"fmt"
)

func minCost(cost [][]int, x int, y int) int {
	if x < 0 || y < 0 {
		return 10000
	}
	if x == 0 && y == 0 {
		return cost[0][0]
	}
	return min(minCost(cost, x-1, y), minCost(cost, x-1, y-1), minCost(cost, x, y-1)) + cost[x][y]
}

func minCost_iter(cost [][]int, x int, y int) int {
	m := make([][]int, x+2)
	for idx, _ := range m {
		m[idx] = make([]int, x+2)
	}
	for i := 0; i < x+2; i++ {
		for j := 0; j < x+2; j++ {
			if i == 0 && j == 0 {
				m[i][j] = 0
			} else if i == 0 || j == 0 {
				m[i][j] = 10000
			} else {
				m[i][j] = min(m[i-1][j], m[i-1][j-1], m[i][j-1]) + cost[i][j]
			}
		}
	}
	return m[x+1][y+1]
}

func min(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func main() {
	m := make([][]int, 3)
	for idx, _ := range m {
		m[idx] = make([]int, 3)
	}
	m[0][0] = 1
	m[0][1] = 2
	m[0][2] = 3
	m[1][0] = 4
	m[1][1] = 8
	m[1][2] = 2
	m[2][0] = 1
	m[2][1] = 5
	m[2][2] = 4
	//fmt.Println(minCost(m, 2, 2))
	fmt.Println(minCost(m, 2, 2))
}
