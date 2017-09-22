package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func isBalanced(root *tNode) int {
	if root == nil {
		return 0
	}
	l := isBalanced(root.left)
	r := isBalanced(root.right)

	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	if l < r {
		return r + 1
	} else {
		return l + 1
	}
}

func isBalanced_top(root *tNode) bool {
	j := isBalanced(root)
	if j == -1 {
		return false
	} else {
		return true
	}
}

func main() {
	t5 := &tNode{val: 6, left: nil, right: nil}
	t4 := &tNode{val: 5, left: t5, right: nil}
	t3 := &tNode{val: 4, left: nil, right: nil}
	t2 := &tNode{val: 3, left: t3, right: t4}
	t1 := &tNode{val: 2, left: nil, right: nil}
	r := &tNode{val: 1, left: t1, right: t2}
	fmt.Println(isBalanced_top(r))
}
