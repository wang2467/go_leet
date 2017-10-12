package main

import (
	"fmt"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
)

func isParlindrome(s string) bool {
	l := len(s) - 1
	for i := 0; i <= l; i++ {
		if s[i] != s[l] {
			return false
		}
		l--
	}
	return true
}

func partitionParlindrome(s string, a int, b int, mystack *lls.Stack) {
	if a > b {
		fmt.Println(mystack.Values())
		return
	}
	for i := a; i <= b; i++ {
		if isParlindrome(s[a : i+1]) {
			mystack.Push(s[a : i+1])
			partitionParlindrome(s, i+1, b, mystack)
			mystack.Pop()
		}
	}

}

func main() {
	mystack := lls.New()
	s := "geekss"
	partitionParlindrome(s, 0, len(s)-1, mystack)
}
