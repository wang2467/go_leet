package main

import (
	"fmt"
)

func F(n int, s []int) int {
	if n < s[0] {
		return 0
	}
	if n == 0 {
		return 1
	}
	if n == s[0] {
		return 1
	}
	r := 0
	for _, m := range s {
		if m < n {
			r += F(n-m, s)
			fmt.Println(r)
		} else {
			break
		}
	}
	return r
}

func F_g(s []int, last int, n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}
	if last < 0 && n > 0 {
		return 0
	}
	return F_g(s, last-1, n) + F_g(s, last, n-s[last])
}

func main() {
	set := []int{6, 5, 3, 2}
	fmt.Println(F_g(set, 3, 10))
}
