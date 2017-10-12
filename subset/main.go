package main

import (
	"fmt"
)

func printSubset(nums []int) {
	data := make([]int, len(nums)-1)
	findSubsets(nums, 0, len(nums)-1, data, 0)
	for _, x := range nums {
		fmt.Printf(("%d "), x)
	}
	fmt.Println()
}

func findSubsets(s []int, start int, end int, data []int, curr int) {
	if curr == len(data) {
		return
	}
	for j := start; j <= end; j++ {
		data[curr] = s[j]
		findSubsets(s, j+1, end, data, curr+1)
		for i := 0; i <= curr; i++ {
			fmt.Printf("%d ", data[i])
		}
		fmt.Println()
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	printSubset(s)
}
