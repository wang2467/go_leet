package main

import (
	"fmt"
)

func LIS(s []int) int {
	if len(s) == 0 {
		return 1
	} else if s[len(s)-2] < s[len(s)-1] {
		return LIS(s[0:len(s)-1]) + 1
	} else {
		return LIS(s[0 : len(s)-1])
	}
}

func main() {
	s := []int{3, 10, 2, 1, 20}
	fmt.Println(LIS(s))
}
