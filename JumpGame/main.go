package main

import (
	"fmt"
)

func CanJump(nums []int) bool {
	return Jump(nums, len(nums)-1)
}

func Jump(s []int, end int) bool {
	if end == 0 {
		return true
	}
	for i := 0; i < end; i++ {
		if i+s[i] >= end {
			return Jump(s, i)
		}
	}
	return false
}

func main() {
	//s := []int{2, 3, 1, 1, 4}
	s := []int{3, 2, 1, 0, 4}
	fmt.Println(CanJump(s))
}
