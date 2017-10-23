package main

import (
	"fmt"
)

func solve(board [][]byte) {
	rows := len(board)
	if rows == 0 {
		return
	}
	cols := len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == 'O' {
				board[i][j] = '-'
			}
		}
	}
	floodFill(board, rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == '-' {
				board[i][j] = 'X'
			}
		}
	}

}

func floodFill(b [][]byte, rows int, cols int) {
	for i := 0; i < cols; i++ {
		if b[0][i] == '-' {
			floodFillUtil(b, 0, i, 'O')
		}
	}
	for i := 0; i < cols; i++ {
		if b[rows-1][i] == '-' {
			floodFillUtil(b, rows-1, i, 'O')
		}
	}
	for i := 0; i < rows; i++ {
		if b[i][0] == '-' {
			floodFillUtil(b, i, 0, 'O')
		}
	}
	for i := 0; i < rows; i++ {
		if b[i][cols-1] == '-' {
			floodFillUtil(b, i, cols-1, 'O')
		}
	}
}

func floodFillUtil(b [][]byte, i int, j int, t byte) {
	if i < 0 || j < 0 || i >= len(b) || j >= len(b[0]) {
		return
	}
	if b[i][j] == 'X' || b[i][j] == 'O' {
		return
	}
	b[i][j] = t
	floodFillUtil(b, i-1, j, t)
	floodFillUtil(b, i+1, j, t)
	floodFillUtil(b, i, j-1, t)
	floodFillUtil(b, i, j+1, t)
}

func main() {
	a := [][]byte{
		{'X', 'O', 'X', 'X'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O'},
		{'O', 'X', 'O', 'X'},
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=")
	solve(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("%c ", a[i][j])
		}
		fmt.Println()
	}

}
