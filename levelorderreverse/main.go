package main

import (
	"fmt"
)

type tNode struct {
	val   int
	left  *tNode
	right *tNode
}

func LOP(root *tNode) [][]int {
	r := make([][]int, 0)
	h := computeHeight(root)
	var tmp []int
	for i := h; i > 0; i-- {
		tmp = []int{}
		tmp = traverseGivenOrder(root, tmp, i)
		r = append(r, tmp)
	}
	return r
}

func traverseGivenOrder(root *tNode, tmp []int, l int) []int {
	if l == 1 {
		tmp = append(tmp, root.val)
		return tmp
	}
	tmp = traverseGivenOrder(root.left, tmp, l-1)
	tmp = traverseGivenOrder(root.right, tmp, l-1)
	return tmp
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
	t3 := &tNode{val: 3, left: nil, right: nil}
	t2 := &tNode{val: 2, left: nil, right: nil}
	t1 := &tNode{val: 1, left: nil, right: nil}
	t1.left = t2
	t1.right = t3
	s := LOP(t1)
	fmt.Println(s)
}
