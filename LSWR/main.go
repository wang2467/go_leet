package main

import (
	"fmt"
)

func findLongestSubstringWithoutRepeatingCharacter(s string) int {
	m := make(map[byte]int)
	var max_val, prv int = 0, 0
	for i := 0; i < len(s); i++ {
		idx, ok := m[s[i]]
		if ok {
			prv = idx + 1
			max_val = max(i-prv, max_val)
			m[s[i]] = i
		} else {
			m[s[i]] = i
			max_val = max(max_val, i-prv+1)
		}
	}
	return max_val
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := "abcadcbb"
	fmt.Println(findLongestSubstringWithoutRepeatingCharacter(s))
}
