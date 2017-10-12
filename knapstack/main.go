package main

import (
	"fmt"
)

func findMax(val []int, wgt []int, w int) int {
	if w < 0 || len(val) == 0 || len(wgt) == 0 {
		return -100
	}
	if w == 0 {
		return 0
	}
	l := len(val) - 1
	return max(findMax(val[0:l], wgt[0:l], w), findMax(val, wgt, w-wgt[l])+val[l])
}

func findMax_without_rep(val []int, wgt []int, w int) int {
	if w < 0 || len(val) == 0 || len(wgt) == 0 {
		return -100
	}
	if w == 0 {
		return 0
	}
	l := len(val) - 1
	return max(findMax_without_rep(val[0:l], wgt[0:l], w), findMax_without_rep(val[0:l], wgt[0:l], w-wgt[l])+val[l])
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := []int{60, 100, 120}
	w := []int{10, 20, 30}
	fmt.Println(findMax_without_rep(s, w, 50))
}
