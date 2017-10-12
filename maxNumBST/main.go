package main

import (
	"fmt"
)

func _numTrees(num int, memo []int) int {
	if num == 0 {
		return 1
	}
	if num == 1 {
		return 1
	}
	sum := 0
	var l, r int
	for i := 0; i < num; i++ {
		//sum += _numTrees(i) * _numTrees(num-i-1)
		if memo[i] != 0 {
			l = memo[i]
		} else {
			l = _numTrees(i, memo)
			memo[i] = l
		}
		if memo[num-i-1] == 0 {
			r = _numTrees(num-i-1, memo)
			memo[num-i-1] = r
		} else {
			r = memo[num-i-1]
		}
		sum += l * r
	}
	return sum
}

func numTrees(num int) int {
	m := make([]int, num)
	m[0] = 1
	m[1] = 1
	return _numTrees(num, m)
}

func main() {
	fmt.Println(numTrees(1))
}
