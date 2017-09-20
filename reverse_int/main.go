package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	s := []int{1, 3, 4, 6, 2, 15}

	go func() {
		for i := 0; i < len(s); i++ {
			fmt.Println(": ", s[i])
			c <- s[i]
		}
	}()

	s_new := make([]int, 6)
	for i := 5; i >= 0; i-- {
		s_new[i] = <-c
	}
	close(c)
	fmt.Println(s_new)
}
