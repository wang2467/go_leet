package main

import (
	"fmt"
	"sort"
	"strconv"
)

type sortString []string

func (s sortString) Less(i, j int) bool {
	return s[i]+s[j] < s[j]+s[i]
}

func (s sortString) Len() int {
	return len(s)
}

func (s sortString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func largestNumber(nums []int) string {
	s := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		s[i] = strconv.Itoa(nums[i])
	}
	sort.Sort(sortString(s))
	var result string = ""
	for i := len(nums) - 1; i >= 0; i-- {
		if i != len(nums)-1 || s[i] != "0" {
			result += s[i]
		}
	}
	return result
}

func main() {
	n := []int{3, 34, 30, 320, 9, 5}
	fmt.Println(largestNumber(n))
}
