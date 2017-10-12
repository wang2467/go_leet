package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func invert_binary_tree(root *tNode) *tNode {
	if root == nil {
		return nil
	}
	if root.left == nil && root.right == nil {
		return root
	}
	tmp := root.right
	root.right = invert_binary_tree(root.left)
	root.left = invert_binary_tree(tmp)
	return root
}

func printTree(root *tNode) {
	if root == nil {
		return
	}
	printTree(root.left)
	fmt.Println(root.val)
	printTree(root.right)
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

	r := invert_binary_tree(t9)
	printTree(r)
	fmt.Println(":", r.left.val)

}
