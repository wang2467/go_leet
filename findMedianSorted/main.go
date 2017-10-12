package main

import (
	"fmt"
)

func findMedian(arr1 []int, arr2 []int) int {
	i := 0
	j := 0
	k := 0
	result := make([]int, len(arr1)+len(arr2))
	for i < len(arr1) || j < len(arr2) {
		if i == len(arr1) {
			copy(result[k:], arr2[j:])
			break
		} else if j == len(arr2) {
			copy(result[k:], arr1[i:])
			break
		} else if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	fmt.Println(result)
	return result[len(result)/2]
}

func main() {
	s1 := []int{1, 4, 5, 7}
	s2 := []int{-1, 2, 3, 9}
	fmt.Println(findMedian(s1, s2))
}
