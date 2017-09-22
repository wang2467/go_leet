package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func reverselist(prev *Node, curr *Node) *Node {
	if curr == nil {
		return nil
	}
	if curr.next == nil {
		var head *Node = curr
		curr.next = prev
		return head
	}
	next := curr.next
	curr.next = prev
	head := reverselist(curr, next)
	return head
}

func iterreverselist(head *Node) *Node {
	curr := head
	var prev *Node = nil
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	n1 := &Node{val: 5, next: nil}
	n2 := &Node{val: 4, next: n1}
	n3 := &Node{val: 3, next: n2}
	n4 := &Node{val: 2, next: n3}
	n5 := &Node{val: 1, next: n4} //head
	printlist(n5)
	n5 = reverselist(nil, n5)
	n5 = iterreverselist(n5)
	printlist(n5)
}

func printlist(h *Node) {
	tmp := h
	for tmp != nil {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}
