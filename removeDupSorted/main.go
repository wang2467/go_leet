package main

import (
	"fmt"
)

func removeDuplicate(nums []int) int {
	diff := nums[0]
	count := 1
	for i := 0; i < len(nums); i++ {
		if diff != nums[i] {
			count++
			diff = nums[i]
		}
	}
	return count
}

func main() {
	fmt.Println(removeDuplicate([]int{1, 1, 2, 2, 3, 4, 4, 5}))
}
