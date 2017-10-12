package main

import (
	"fmt"
)

func findMax(s []int, a int, b int) int {
	if a >= b {
		return -1
	}
	mid := (a + b) / 2
	if s[mid] > s[mid+1] {
		return mid
	}
	x := findMax(s, a, mid-1)
	if x == -1 {
		return findMax(s, mid+1, b)
	} else {
		return x
	}
}

func rotatedSearch(num []int, target int) int {
	max := findMax(num, 0, len(num)-1)
	if target == num[0] {
		return 0
	}
	if target == num[max] {
		return max
	}
	if target < num[0] {
		return binarySearch(num, max, len(num), target)
	} else {
		return binarySearch(num, 0, max-1, target)
	}

}

func binarySearch(s []int, a int, b int, t int) int {
	if a >= b {
		return -1
	}
	mid := (a + b) / 2
	if t == s[mid] {
		return mid
	}
	if t < s[mid] {
		return binarySearch(s, a, mid-1, t)
	} else {
		return binarySearch(s, mid+1, b, t)
	}
}

func main() {
	s := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMax(s, 0, len(s)-1))
	fmt.Println(rotatedSearch(s, 4))
}
