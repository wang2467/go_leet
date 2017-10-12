package main

import (
	"fmt"
)

type lNode struct {
	val  int
	next *lNode
}

func addTwoNummbers(l1 *lNode, l2 *lNode) *lNode {
	tmp1 := l1
	tmp2 := l2
	c := 0
	var head, tail *lNode = nil, nil
	var tmp_val int
	for tmp1 != nil || tmp2 != nil {
		if tmp1 != nil && tmp2 != nil {
			tmp_val = tmp1.val + tmp2.val + c
			if tmp_val >= 10 {
				tmp_val = tmp_val - 10
				c = 1
			}
		} else if tmp1 != nil {
			tmp_val = tmp1.val + 0 + c
		} else {
			tmp_val = tmp2.val + 0 + c
		}
		head, tail = addNode(tmp_val, head, tail)
		tmp1 = tmp1.next
		tmp2 = tmp2.next
	}
	return head

}

func addNode(val int, head *lNode, tail *lNode) (*lNode, *lNode) {
	if head == nil && tail == nil {
		newNode := &lNode{val: val, next: nil}
		head = newNode
		tail = newNode
	} else {
		newNode := &lNode{val: val, next: nil}
		tail.next = newNode
		tail = newNode
	}
	return head, tail
}

func printList(head *lNode) {
	tmp := head
	for tmp != nil {
		fmt.Println(tmp.val)
		tmp = tmp.next
	}
}
func main() {
	var h1, h2, t1, t2 *lNode = nil, nil, nil, nil
	h1, t1 = addNode(2, h1, t1)
	h1, t1 = addNode(4, h1, t1)
	h1, t1 = addNode(3, h1, t1)
	h2, t2 = addNode(5, h2, t2)
	h2, t2 = addNode(6, h2, t2)
	h2, t2 = addNode(4, h2, t2)
	//printList(h1)
	//printList(h2)
	printList(addTwoNummbers(h1, h2))

}
