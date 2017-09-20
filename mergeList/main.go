package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	var tail *ListNode = nil
	for l1 != nil || l2 != nil {
		if l1 == nil {
			for l2 != nil {
				head, tail = addNode(l2.Val, tail, head)
				l2 = l2.Next
			}
			return head
		}
		if l2 == nil {
			for l1 != nil {
				head, tail = addNode(l1.Val, tail, head)
				l1 = l1.Next
			}
			return head
		}
		if l1.Val < l2.Val {
			head, tail = addNode(l1.Val, tail, head)
			l1 = l1.Next
		} else {
			head, tail = addNode(l2.Val, tail, head)
			l2 = l2.Next
		}
	}
	return head
}

func addNode(x int, tail *ListNode, head *ListNode) (*ListNode, *ListNode) {
	if head == nil && tail == nil {
		tmp := &ListNode{Val: x, Next: nil}
		tail = tmp
		head = tmp
	} else {
		tmp := &ListNode{Val: x, Next: nil}
		tail.Next = tmp
		tail = tmp
	}
	return head, tail
}

func main() {
	l_h := &ListNode{Val: 1, Next: nil}
	l_t := l_h
	l_h, l_t = addNode(3, l_t, l_h)
	l_h, l_t = addNode(5, l_t, l_h)
	l_h, l_t = addNode(7, l_t, l_h)
	l_h, l_t = addNode(9, l_t, l_h)
	l_h, l_t = addNode(11, l_t, l_h)

	j_h := &ListNode{Val: 2, Next: nil}
	j_t := j_h
	j_h, j_t = addNode(4, j_t, j_h)
	j_h, j_t = addNode(6, j_t, j_h)
	j_h, j_t = addNode(8, j_t, j_h)
	j_h, j_t = addNode(10, j_t, j_h)
	j_h, j_t = addNode(12, j_t, j_h)

	n := mergeTwoLists(l_h, j_h)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}
