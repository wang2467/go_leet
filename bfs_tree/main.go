package main

import (
	"fmt"
	lane "gopkg.in/oleiade/lane.v1"
)

type tNode struct {
	value int
	left  *tNode
	right *tNode
}

func main() {
	s := []int{4, 7, 5, 2, 3, 6}
	r := preorder_construct(s)
	//preorder_print(r)
	//bfs_recur(r)
	//fmt.Println(r.right.left.left.value)
	//preorder_stack(r)
	postorder_stack(r)
	fmt.Println("h:", computeHeight(r))
}

func preorder_construct(arr []int) *tNode {
	root := &tNode{value: arr[0], left: nil, right: nil}
	stack := lane.NewStack()
	stack.Push(root)
	for i := 1; i < len(arr); i++ {
		var tmp *tNode = nil
		for !stack.Empty() && arr[i] > stack.Head().(*tNode).value {
			tmp = stack.Pop().(*tNode)
		}
		node := &tNode{value: arr[i], left: nil, right: nil}
		if tmp == nil {
			stack.Head().(*tNode).left = node
			stack.Push(node)
		} else {
			tmp.right = node
			stack.Push(node)
		}
	}
	return root
}

func bfs(root *tNode) {
	q := lane.NewQueue()
	q.Enqueue(root)
	for !q.Empty() {
		r := q.Dequeue().(*tNode)
		if r.left != nil {
			q.Enqueue(r.left)
		}
		if r.right != nil {
			q.Enqueue(r.right)
		}
		fmt.Println(r.value)
	}
}

func bfs_recur(root *tNode) {
	for i := 1; i <= computeHeight(root); i++ {
		_bfs_recur(root, i)
	}
}

func _bfs_recur(root *tNode, l int) {
	if l == 0 || root == nil {
		return
	}
	if l == 1 && root != nil {
		fmt.Println(root.value)
	} else if l > 1 {
		_bfs_recur(root.left, l-1)
		_bfs_recur(root.right, l-1)
	}
}

func computeHeight(root *tNode) int {
	if root == nil {
		return 0
	}
	lheight := computeHeight(root.left)
	rheight := computeHeight(root.right)
	max := func(x int, y int) int {
		if x < y {
			return y + 1
		} else {
			return x + 1
		}
	}
	return max(lheight, rheight)
}

func preorder_stack(root *tNode) {
	s := lane.NewStack()
	s.Push(root)
	for !s.Empty() {
		tmp := s.Pop().(*tNode)
		fmt.Println(tmp.value)
		if tmp.right != nil {
			s.Push(tmp.right)
		}
		if tmp.left != nil {
			s.Push(tmp.left)
		}
	}
}

func postorder_stack(root *tNode) {
	s1 := lane.NewStack()
	s2 := lane.NewStack()
	s1.Push(root)
	for !s1.Empty() {
		tmp := s1.Pop().(*tNode)
		s2.Push(tmp)
		if tmp.left != nil {
			s1.Push(tmp.left)
		}
		if tmp.right != nil {
			s1.Push(tmp.right)
		}
	}
	for !s2.Empty() {
		fmt.Println(s2.Pop().(*tNode).value)
	}
}

func preorder_print(root *tNode) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	preorder_print(root.left)
	preorder_print(root.right)
}
