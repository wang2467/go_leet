package main

import (
	"fmt"
	"sort"
)

func maxProduct(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	sort.Ints(nums)
	max := 0
	for i := len(nums) - 1; i > 0; i-- {
		var a int
		if nums[i] > nums[i-1] {
			a = product(nums, 0, i, 1, m)
		} else {
			a = product(nums, 0, i, 0, m)
		}
		if max < a {
			max = a
		}
	}
	return max
}

func product(nums []int, a int, b int, d int, m map[int]int) int {
	if a >= b {
		return nums[a]
	}
	if d == 0 {
		if m[nums[b]]+1 == m[nums[b-1]] {
			return nums[b] * product(nums, a, b-1, d, m)
		} else {
			return nums[b]
		}
	} else if d == 1 {
		if m[nums[b]]-1 == m[nums[b-1]] {
			return nums[b] * product(nums, a, b-1, d, m)
		} else {
			return nums[b]
		}
	}
	return 1
}

func main() {
	s := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(s))
}
