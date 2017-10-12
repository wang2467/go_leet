package main

import (
	"fmt"
	"sort"
)

type sol struct {
	val1 int
	val2 int
}

type ByVal1 []sol

func (a ByVal1) Len() int               { return len(a) }
func (a ByVal1) Swap(x int, y int)      { a[x], a[y] = a[y], a[x] }
func (a ByVal1) Less(x int, y int) bool { return a[x].val1 < a[y].val1 }

func twoSum(arr []int, t int) (int, int) {
	l := len(arr) - 1
	f := 0
	for f < l {
		if arr[f]+arr[l] == t {
			return arr[f], arr[l]
		} else if arr[f]+arr[l] < t {
			f++
		} else {
			l--
		}
	}
	return -99, -99
}

func threeSum(arr []int, t int) [][]int {
	m := make([][]int, 0)
	for idx, x := range arr {
		a1, a2 := twoSum(arr[idx+1:], t-x)
		if a1 != -99 && a2 != 99 {
			m = append(m, []int{x, a1, a2})
		}
	}
	return m
}

func main() {
	s := []int{-4, -1, -1, 0, 1, 2}
	fmt.Println(threeSum(s, 0))
	m := make(map[int]sol)
	m[1] = sol{val1: 2, val2: 3}
	fmt.Println(m[5])
	sort.Ints(s)
	fmt.Println(s)
	x := []sol{{val1: 4, val2: 3}, {val1: 3, val2: 5}}
	fmt.Println(x)
	sort.Sort(ByVal1(x))
	fmt.Println(x)
}
