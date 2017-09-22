package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func sortedArrayToBST(n []int, x int, y int) *tNode {
	if x > y {
		return nil
	}
	mid := (x + y) / 2
	r := &tNode{val: n[mid], left: nil, right: nil}
	r.left = sortedArrayToBST(n, x, mid-1)
	r.right = sortedArrayToBST(n, mid+1, y)
	return r
}

func printBST(root *tNode) {
	if root == nil {
		return
	}
	printBST(root.left)
	fmt.Println(root.val)
	printBST(root.right)
}

func computeHeight(root *tNode) int {
	if root == nil {
		return 0
	}
	l := computeHeight(root.left)
	r := computeHeight(root.right)

	if l < r {
		return r + 1
	} else {
		return l + 1
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	r := sortedArrayToBST(s, 0, 6)
	printBST(r)
	fmt.Println("height: ", computeHeight(r))
}
