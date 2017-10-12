package main

import (
	"fmt"
)

func permuteNum(s []int) {
	dct := make(map[int]int)
	data := make([]int, len(s))
	for _, x := range s {
		dct[x] = -1
	}
	permute(data, dct, 0, len(s))
}

func permute(data []int, dct map[int]int, curr int, size int) {
	if curr == size-1 {
		for k, v := range dct {
			if v == -1 {
				data[curr] = k
			}
		}
		for _, x := range data {
			fmt.Printf("%d", x)
		}
		fmt.Println()
		return
	}

	for key, value := range dct {
		if value == -1 {
			dct[key] = 1
			data[curr] = key
			permute(data, dct, curr+1, size)
			dct[data[curr]] = -1
		}
	}
}

func permute_2(s []int, l int, r int) {
	if l == r {
		for _, x := range s {
			fmt.Printf("%d", x)
		}
		fmt.Println()
		return
	}

	for i := l; i <= r; i++ {
		s[l], s[i] = s[i], s[l]
		permute_2(s, l+1, r)
		s[l], s[i] = s[i], s[l]

	}
}

func main() {
	s := []int{1, 2, 3, 4}
	//permuteNum(s)
	permute_2(s, 0, 3)
}
