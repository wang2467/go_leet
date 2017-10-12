package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(s []int, a int, b int, target int) []int {
	if a > b {
		return []int{-1, -1}
	}
	mid := (a + b) / 2
	if s[mid] == target {
		var i, j int = mid, mid
		for i >= 0 && i < len(s) && s[i] == s[mid] {
			i++
		}
		for j >= 0 && j < len(s) && s[j] == s[mid] {
			j--
		}
		return []int{j + 1, i - 1}
	}
	if target < s[mid] {
		return binarySearch(s, a, mid-1, target)
	} else {
		return binarySearch(s, mid+1, b, target)
	}
}

func main() {
	s := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(searchRange(s, 10))
}
