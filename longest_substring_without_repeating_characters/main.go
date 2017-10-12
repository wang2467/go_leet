package main

import (
	"fmt"
)

func findLength(s string) int {
	m := make(map[byte]int)
	max := 0
	pv := 0
	j := 0
	for i := 0; i < len(s) && j < len(s); i++ {
		_, ok := m[s[i]]
		if ok {
			pv = i
		} else {
			max = maxNum(max, i-pv+1)
		}
		m[s[i]] = i
	}
	return max
}

func main() {
	fmt.Println(findLength("abcabcbb"))
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
