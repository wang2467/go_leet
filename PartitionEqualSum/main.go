package main

import (
	"fmt"
)

func findPartition(s []int) bool {
	sum := 0
	for _, x := range s {
		sum += x
	}
	if sum%2 == 0 {
		return Partition(s, len(s)-1, sum/2)
	} else {
		return false
	}
}

func Partition(s []int, a int, target int) bool {
	if target == 0 {
		return true
	}
	if a < 0 {
		return false
	}
	return Partition(s, a-1, target-s[a]) || Partition(s, a-1, target)
}

func main() {
	s := []int{4, 5}
	fmt.Println(findPartition(s))
}
