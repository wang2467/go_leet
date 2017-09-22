package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func findNode(root *tNode, val int) bool {
	if root == nil {
		return false
	}
	if root.val == val {
		return true
	}
	if findNode(root.left, val) {
		return true
	} else {
		return findNode(root.right, val)
	}
}

func lowest_common_ancestor_bt(root *tNode, x int, y int) *tNode {
	if findNode(root.left, x) && findNode(root.right, y) {
		return root
	}
	if findNode(root.left, x) {
		return lowest_common_ancestor_bt(root.left, x, y)
	}
	if findNode(root.right, y) {
		return lowest_common_ancestor_bt(root.right, x, y)
	}
	return nil
}

func lowest_common_ancestor_bt2(root *tNode, x int, y int) *tNode {
	if root == nil {
		return nil
	}
	if root.val == x || root.val == y {
		return root
	}
	left := lowest_common_ancestor_bt2(root.left, x, y)
	right := lowest_common_ancestor_bt2(root.right, x, y)
	if left == nil {
		return right
	} else if right == nil {
		return left
	} else {
		return root
	}
}

func main() {
	t1 := &tNode{val: 7, left: nil, right: nil}
	t2 := &tNode{val: 4, left: nil, right: nil}
	t3 := &tNode{val: 6, left: nil, right: nil}
	t4 := &tNode{val: 2, left: t1, right: t2}
	t5 := &tNode{val: 0, left: nil, right: nil}
	t6 := &tNode{val: 8, left: nil, right: nil}
	t7 := &tNode{val: 5, left: t3, right: t4}
	t8 := &tNode{val: 1, left: t5, right: t6}
	t9 := &tNode{val: 3, left: t7, right: t8}             //dumbest way of building tree
	fmt.Println(lowest_common_ancestor_bt2(t9, 6, 0).val) //assuming x and y are the left node and the right node, respectively
}
