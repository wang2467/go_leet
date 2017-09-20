package main

import (
	"fmt"
	"sort"
)

type ByValue []int

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i] < a[j] }

func BinarySearch(arr []int, target int) int {
	idx := len(arr) / 2
	if arr[idx] == target {
		return idx
	}
	if arr[idx] < target {
		return BinarySearch(arr[idx+1:], target)
	} else {
		return BinarySearch(arr[:idx], target)
	}

}

func findPair(nums []int, target int) []int {
	i := 0
	j := 1
	for i < len(nums) && j < len(nums) && j > i {
		if nums[j]-nums[i] == target {
			return []int{i, j}
		} else if nums[j]-nums[i] > target {
			i++
		} else {
			j++
		}
	}
	return nil
}

func TwoSum(nums []int, target int) []int {
	up := len(nums) - 1
	down := 0

	for down < up {
		sum := nums[up] + nums[down]
		if sum == target {
			return []int{down, up}
		} else if sum < target {
			down++
		} else {
			up--
		}
	}
	return nil
}

func main() {
	s := []int{13, 2, 9, 5, 6, 4, 17, 3, 0, 1, 20, 1}
	sort.Sort(ByValue(s))
	fmt.Println(s)
	r := findPair(s, 0)
	//r := TwoSum(s, 28)
	fmt.Println(r)
}
