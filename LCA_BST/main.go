package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func lowest_common_ancestor_bst(root *tNode, x int, y int) *tNode {
	if root == nil || root.val >= x && root.val <= y {
		return root
	}
	if y <= root.val {
		return lowest_common_ancestor_bst(root.left, x, y)
	}
	if x >= root.val {
		return lowest_common_ancestor_bst(root.right, x, y)
	}
	return nil
}

func main() {
	t1 := &tNode{val: 3, left: nil, right: nil}
	t2 := &tNode{val: 5, left: nil, right: nil}
	t3 := &tNode{val: 0, left: nil, right: nil}
	t4 := &tNode{val: 4, left: t1, right: t2}
	t5 := &tNode{val: 7, left: nil, right: nil}
	t6 := &tNode{val: 9, left: nil, right: nil}
	t7 := &tNode{val: 2, left: t3, right: t4}
	t8 := &tNode{val: 8, left: t5, right: t6}
	t9 := &tNode{val: 6, left: t7, right: t8} //dumbest way of building tree

	fmt.Println(lowest_common_ancestor_bst(t9, 2, 8).val)
}
