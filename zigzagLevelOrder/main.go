package main

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
)

type tNode struct {
	value int
	left  *tNode
	right *tNode
}

func zigzagLevelOrder(root *tNode) [][]int {
	q := arraylist.New()
	q.Add(root)
	level := 0
	res := make([][]int, 0)
	for !q.Empty() {
		size := q.Size()
		res_sub := make([]int, size)
		for i := 0; i < size; i++ {
			r, _ := q.Get(0)
			q.Remove(0)
			if level%2 == 0 {
				res_sub[i] = r.(*tNode).value
			} else {
				res_sub[size-i-1] = r.(*tNode).value
			}
			if r.(*tNode).left != nil {
				q.Add(r.(*tNode).left)
			}
			if r.(*tNode).right != nil {
				q.Add(r.(*tNode).right)
			}
		}
		res = append(res, res_sub)
		level++
	}
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

	fmt.Println(zigzagLevelOrder(root))
}
