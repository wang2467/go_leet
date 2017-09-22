package main

import (
	"fmt"
)

func Find_RP(arr []int, x int, y int) int {
	if x >= y {
		return -1
	}
	mid := (x + y) / 2
	var result int
	if mid == 0 && arr[mid] == arr[mid+1] {
		return arr[mid]
	} else if mid == len(arr)-1 && arr[mid] == arr[mid-1] {
		return arr[mid]
	} else if mid != 0 && mid != len(arr)-1 && (arr[mid] == arr[mid+1] || arr[mid] == arr[mid-1]) {
		return arr[mid]
	} else {
		result = Find_RP(arr, x, mid)
		if result != -1 {
			return result
		} else {
			result = Find_RP(arr, mid+1, y)
		}
		return result
	}
	return result
}

func main() {
	s := []int{3, 3, 3, 3, 3}
	fmt.Println(Find_RP(s, 0, len(s)-1))
}
