package main

import (
	"fmt"
)

func Shuffle(arr []int, x int, y int) {
	if y-x == 1 {
		return
	}
	mid := (x + y) / 2
	mid_n := mid + 1
	mmid := (x+mid)/2 + 1
	for i := mmid; i <= mid; i++ {
		tmp := arr[i]
		arr[i] = arr[mid_n]
		arr[mid_n] = tmp
		mid_n++
	}
	Shuffle(arr, x, mid)
	Shuffle(arr, mid+1, y)
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Shuffle(n, 0, 7)
	fmt.Println(n)
}
