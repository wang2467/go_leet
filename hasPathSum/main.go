package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func hasPathSum(root *tNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.val == sum && root.left == nil && root.right == nil {
		return true
	}
	l := hasPathSum(root.left, sum-root.val)
	r := hasPathSum(root.right, sum-root.val)
	if l {
		return true
	} else {
		return r
	}
}

func main() {
	t1 := &tNode{val: 7, left: nil, right: nil}
	t2 := &tNode{val: 2, left: nil, right: nil}
	t3 := &tNode{val: 1, left: nil, right: nil}
	t4 := &tNode{val: 11, left: t1, right: t2}
	t5 := &tNode{val: 13, left: nil, right: nil}
	t6 := &tNode{val: 4, left: nil, right: t3}
	t7 := &tNode{val: 4, left: t4, right: nil}
	t8 := &tNode{val: 8, left: t5, right: t6}
	t9 := &tNode{val: 5, left: t7, right: t8}

	fmt.Println(hasPathSum(t9, 26))
}
