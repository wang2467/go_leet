package main

import (
	"fmt"
	//"strings"
)

type Matrix [][]int

func findPath(m Matrix, x int, y int) int {
	if x == 5 && y == 5 {
		return 1
		//} else if m[x+1][y] == 0 && m[x][y+1] == 0 {
	} else if x > 5 || y > 5 {
		return 0
	} else if m[x][y] == 0 {
		return 0
	} else {
		fmt.Println(m[x][y])
		return findPath(m, x+1, y) + findPath(m, x, y+1)
	}
}

func _lcp(s1 string, s2 string) string {
	l1 := len([]rune(s1))
	l2 := len([]rune(s2))
	var i, j int
	for i, j = 0, 0; i < l1 && j < l2; i, j = i+1, j+1 {
		if s1[i] != s2[j] {
			break
		}
	}
	return s1[:i]
}

func LCP(s []string, low int, high int) string {
	if low >= high {
		return s[high]
	}
	mid := (low + high) / 2
	x := LCP(s, low, mid)
	y := LCP(s, mid+1, high)
	return _lcp(x, y)
}

func main() {
	s := make(Matrix, 6)
	for i := range s {
		s[i] = make([]int, 0)
	}
	s[0] = append(s[0], 1, 0, 1, 1, 1, 0)
	s[1] = append(s[1], 1, 1, 1, 0, 1, 0)
	s[2] = append(s[2], 0, 0, 1, 1, 0, 1)
	s[3] = append(s[3], 1, 1, 1, 1, 0, 1)
	s[4] = append(s[4], 0, 0, 1, 0, 1, 1)
	s[5] = append(s[5], 1, 1, 1, 1, 1, 1)
	//l := findPath(s, 0, 0)
	fmt.Println(LCP([]string{"geeks", "geeksa", "geekseee", "geeksforgeeks", "geeksaa", "geeksappc"}, 0, 5))
}
