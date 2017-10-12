package main

import (
	"fmt"
	//"reflect"
)

func FindDifferentCharacterInStrings(s1 string, s2 string) []rune {
	r := make([]rune, 0)
	for _, i := range s1 {
		if !(BinarySearch(s2, 0, len([]rune(s2))-1, i)) {
			r = append(r, i)
		}
	}
	return r
}

func BinarySearch(s string, x int, y int, c rune) bool {
	if x > y {
		return false
	}
	mid := (x + y) / 2
	if rune(s[mid]) == c {
		return true
	}
	if BinarySearch(s, x, mid-1, c) {
		return true
	} else {
		return BinarySearch(s, mid+1, y, c)
	}
}

func main() {
	a := "abc"
	b := "bcd"
	// fmt.Println(reflect.TypeOf(a))
	// fmt.Println(reflect.TypeOf(a[1]))
	// for _, x := range a {
	// 	fmt.Println(reflect.TypeOf(x))

	r := FindDifferentCharacterInStrings(a, b)
	for _, i := range r {
		fmt.Printf("%c\n", i)
	}
}
