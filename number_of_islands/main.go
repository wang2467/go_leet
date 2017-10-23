package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	color := make([][]int, len(grid))
	for i := range color {
		color[i] = make([]int, len(grid[0]))
	}
	count := 0
	for i := 0; i < len(color); i++ {
		for j := 0; j < len(color[0]); j++ {
			if grid[i][j] == '1' && color[i][j] == 0 {
				count += floodFill(grid, color, i, j)
			}
		}
	}
	return count
}

func floodFill(gird [][]byte, color [][]int, i int, j int) int {
	if i < 0 || i >= len(gird) || j < 0 || j >= len(gird[0]) {
		return 1
	}
	if gird[i][j] == '0' || color[i][j] == 1 {
		return 1
	}
	color[i][j] = 1
	if (floodFill(gird, color, i-1, j) * floodFill(gird, color, i+1, j) * floodFill(gird, color, i, j-1) * floodFill(gird, color, i, j+1)) == 1 {
		return 1
	}
	return 0
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
