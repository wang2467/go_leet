package main

import (
	"fmt"
)

func findWater(s []int) int {
	w := 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
		} else if i == len(s)-1 {
		} else {
			r := min(findLeftMax(s[0:i]), findRightMax(s[i+1:]))
			if r > s[i] {
				w += r - s[i]
			}
		}
	}
	return w
}

func findLeftMax(s []int) int {
	l := len(s)
	var i int
	max := 0
	for i = l - 1; i >= 0; i-- {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func findRightMax(s []int) int {
	l := len(s)
	var i int
	max := 0
	for i = 0; i < l; i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func main() {
	//s := []int{3, 0, 0, 2, 0, 4}
	s := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(findWater(s))
	//fmt.Println(len(s[:0]))
}
