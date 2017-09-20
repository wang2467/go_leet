package main

import (
	"fmt"
)

func FindBP(arr []int, x int, y int) int {
	mid := (x + y) / 2
	var result int
	if y <= x {
		return -1
	}
	if mid == 0 || mid == len(arr)-1 {
		return -1
	}
	if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
		return arr[mid]
	}
	if arr[mid] > arr[mid-1] {
		result = FindBP(arr, mid+1, y)
	} else {
		result = FindBP(arr, x, mid-1)
	}
	return result
}

func FindBP_D(arr []int, x int, y int) int {
	s := 0
	for f := 1; f < len(arr); f++ {
		if arr[s] >= arr[f] && s != 0 {
			return arr[s]
		}
		s++
	}
	return -1
}

func main() {
	s := []int{111, 2, 3, 4, 5, 11, 19, 28}
	fmt.Println(FindBP(s, 0, len(s)-1))
	fmt.Println(FindBP_D(s, 0, len(s)-1))
}
