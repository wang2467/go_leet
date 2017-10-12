package main

import (
	"fmt"
)

func LIS(s []int, max *int) int {
	if len(s) == 1 {
		return 1
	}
	var r, m int = 1, 1
	for i := 0; i < len(s)-1; i++ {
		r = LIS(s[0:i+1], max)
		if s[len(s)-1] > s[i] {
			m = r + 1
		}
	}
	if *max < m {
		*max = m
	}
	return m
}

func main() {
	s := []int{5, 22, 8, 6, 9}
	var m int = 0
	fmt.Println(LIS(s, &m))
}
