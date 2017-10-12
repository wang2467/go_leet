package main

import (
	"fmt"
)

// func letterCombinations(digit string) []string {
// 	// m := map[string]string{
// 	// 	"2": "abc",
// 	// 	"3": "def",
// 	// 	"4": "ghi",
// 	// 	"5": "jkl",
// 	// 	"6": "mno",
// 	// 	"7": "pqrs",
// 	// 	"8": "tuv",
// 	// 	"9": "wxyz",
// 	// }
// }

func combination(s []int, a int, b int, data []int, curr int, n int, memo [][]int) {
	if curr == n {
		for i := 0; i < len(data); i++ {
			//fmt.Printf("%d", data[i])
		}
		//fmt.Println()
		memo = append(memo, data)
		// fmt.Println(memo)
		return
	}

	for k := a; k <= b && k <= b+1-n+curr; k++ {
		data[curr] = s[k]
		combination(s, k+1, b, data, curr+1, n, memo)
	}
}

func findCombination(s []int, num int) {
	memo := make([][]int, 0)
	data := make([]int, num)
	combination(s, 0, len(s)-1, data, 0, num, memo)
	for _, x := range memo {
		fmt.Println(x)
	}
}

func main() {
	//data := make([]int, 3, 3)
	s := []int{0, 1, 2, 3, 4}
	//combination(s, 0, 4, data, 0, 3)
	findCombination(s, 3)

}
