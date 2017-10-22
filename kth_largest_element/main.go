package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	result := make([]int, k)
	for i := range nums {
		if i < k {
			result[i] = nums[i]
			upwardHeapify(result, i)
		} else {
			if result[0] < nums[i] {
				result[0] = nums[i]
				downwardHeapify(result, 0)
			}
		}
	}
	return result[0]
}

func upwardHeapify(s []int, n int) {
	child := n
	parent := (child - 1) / 2
	for s[child] < s[parent] && parent >= 0 {
		s[child], s[parent] = s[parent], s[child]
		child = parent
		parent = (child - 1) / 2
	}
}

func downwardHeapify(s []int, n int) {
	parent := n
	child1 := min(2*parent+1, len(s)-1)
	child2 := min(2*parent+2, len(s)-1)
	for s[parent] > s[child1] || s[parent] > s[child2] {
		if s[child1] < s[child2] {
			s[child1], s[parent] = s[parent], s[child1]
			parent = child1
		} else {
			s[child2], s[parent] = s[parent], s[child2]
			parent = child2
		}
		child1 = min(2*parent+1, len(s)-1)
		child2 = min(2*parent+2, len(s)-1)
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(nums, 6))
}
