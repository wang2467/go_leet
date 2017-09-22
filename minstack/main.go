package main

import (
	"fmt"
)

func upwardHeapify(arr []int, x int, y int) []int {
	child := y
	for child > 0 {
		parent := (child - 1) / 2
		if arr[parent] > arr[child] {
			tmp := arr[parent]
			arr[parent] = arr[child]
			arr[child] = tmp
		}
		child = parent
	}
	return arr
}

func main() {
	r := []int{2, 4, 6, 5, 9, 1}
	a := upwardHeapify(r, 0, 5)
	fmt.Println(a)
}
