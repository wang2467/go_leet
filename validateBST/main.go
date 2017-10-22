package main

import (
	"fmt"
)

type tNode struct {
	value int
	left  *tNode
	right *tNode
}

func findMax(root *tNode) int {
	if root.right == nil {
		return root.value
	} else {
		return findMax(root.right)
	}
}

func findMin(root *tNode) int {
	if root.left == nil {
		return root.value
	} else {
		return findMin(root.left)
	}
}

func ValidateBST(root *tNode) bool {
	if root == nil {
		return true
	}
	if root.left != nil && root.value < findMax(root.left) {
		return false
	}
	if root.right != nil && root.value > findMin(root.right) {
		return false
	}
	return ValidateBST(root.left) && ValidateBST(root.right)
}

func printTreeLevelOrder(root *tNode, h int, res []int) []int {
	if root == nil {
		return res
	}
	if h == 0 {
		res = append(res, root.value)
		return res
	}
	h = h - 1
	res = printTreeLevelOrder(root.left, h, res)
	res = printTreeLevelOrder(root.right, h, res)
	return res
}

func findDepth(root *tNode) int {
	if root == nil {
		return 0
	}
	return max(findDepth(root.left), findDepth(root.right)) + 1
}

func levelOrder(root *tNode) [][]int {
	d := findDepth(root) - 1
	res := make([][]int, 0)
	for i := 0; i <= d; i++ {
		m := make([]int, 0)
		m = printTreeLevelOrder(root, i, m)
		res = append(res, m)
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func levelOrder2(root *tNode) [][]int {
	d := findDepth(root)
	res := make([][]int, d)
	for i := range res {
		res[i] = make([]int, 0)
	}
	res = _levelOrder2(root, 0, res)
	return res
}

func _levelOrder2(root *tNode, h int, res [][]int) [][]int {
	if root == nil {
		return res
	}
	res[h] = append(res[h], root.value)
	h++
	res = _levelOrder2(root.left, h, res)
	res = _levelOrder2(root.right, h, res)
	return res
}

func main() {
	t7 := &tNode{value: 3, left: nil, right: nil}
	t6 := &tNode{value: 2, left: nil, right: t7}
	t5 := &tNode{value: 5, left: nil, right: nil}
	t4 := &tNode{value: 7, left: nil, right: nil}
	t3 := &tNode{value: 6, left: t5, right: t4}
	t2 := &tNode{value: 4, left: t6, right: t3}
	t1 := &tNode{value: 10, left: nil, right: nil}
	root := &tNode{value: 9, left: t2, right: t1}
	// fmt.Println(findMax(root.left))
	// fmt.Println(findMin(root))
	// fmt.Println(ValidateBST(root))
	fmt.Println(levelOrder2(root))
}
