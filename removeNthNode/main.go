package main

import (
	"fmt"
)

type ListNode struct {
	value int
	next  *ListNode
}

func removeNthNode(head *ListNode, n int) *ListNode {
	senitel := head
	fast := head
	slow := head
	for i := 0; i < n; i++ {
		if i >= (n - 2) {
			fast = fast.next
		}
		senitel = senitel.next
	}
	for senitel != nil {
		senitel = senitel.next
		slow = slow.next
		fast = fast.next
	}
	slow.next = fast
	return head
}

func printList(head *ListNode) {
	temp := head
	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {

	t5 := &ListNode{value: 5, next: nil}
	t4 := &ListNode{value: 4, next: t5}
	t3 := &ListNode{value: 3, next: t4}
	t2 := &ListNode{value: 2, next: t3}
	head := &ListNode{value: 1, next: t2}
	head = removeNthNode(head, 3)
	printList(head)
}
