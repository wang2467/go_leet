package main

import (
	"fmt"
)

func Find_SR(num int, target int) int {
	if num*num <= target && (num+1)*(num+1) > target {
		return num
	}
	return Find_SR(num+1, target)
}

func Find_SR_S(x int, y int, t int) int {
	m := (x + y) / 2
	if m*m <= t && (m+1)*(m+1) > t {
		return m
	}
	var tmp int = m * m
	if tmp < t {
		return Find_SR_S(m+1, y, t)
	} else {
		return Find_SR_S(x, m, t)
	}
}

func Find(t int) int {
	return Find_SR_S(0, t, t)
}

func main() {
	fmt.Println(Find(0))
}
