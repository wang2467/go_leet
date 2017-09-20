package main

import (
	"fmt"
	"math"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func (l *List) append(x int) {
	t := &Node{value: x, next: nil}
	if l.tail == nil && l.head == nil {
		l.tail = t
		l.head = t
	}
	l.tail.next = t
	l.tail = t
}

func reverse(curr *Node, prev *Node) *Node {
	if curr.next == nil {
		head := curr
		curr.next = prev
		return head
	}
	tmp := curr.next
	curr.next = prev
	return reverse(tmp, curr)
}

func isP(num int) bool {
	s := strconv.Itoa(num)
	d := len(s)
	fmt.Println("length: ", d)
	mid := d / 2
	n_num := num
	for i := 0; i < mid; i++ {
		if num%10 != (n_num / int(math.Pow10(d-1))) {
			return false
		}
		num = num / 10
		n_num = n_num % int(math.Pow10(d-1))
		d--
	}
	return true

}

func main() {
	ml := &List{}
	ml.append(2)
	ml.append(3)
	ml.append(4)
	ml.append(5)
	fmt.Println(ml.head.value)
	fmt.Println(ml.tail.value)
	nl := reverse(ml.head, nil)
	fmt.Println(nl.next.value)
	fmt.Println(isP(21313))

}
