package main

import (
	"fmt"
)

func findLPS(s string, x int, y int) int {
	if x == y {
		return 1
	}
	if y == x+1 && s[x] == s[y] {
		return 2
	}
	if s[x] == s[y] {
		return findLPS(s, x+1, y-1) + 2
	}
	if s[x] != s[y] {
		return max(findLPS(s, x+1, y), findLPS(s, x, y-1))
	}
	return -1
}

func findLPS_dp_memo(s string) string {
	n := len(s)
	m := make([][]int, n)
	for a, _ := range m {
		m[a] = make([]int, n)
	}
	x := 0
	r := s[0:1]
	for g := 1; g <= n; g++ {
		for i := 0; i <= n-g; i++ {
			j := i + g - 1
			if j-i == 0 && s[i] == s[j] {
				m[i][j] = 1
			} else if s[i] == s[j] && j-i == 1 {
				m[i][j] = 2
				if max(x, 2) != x {
					x = 2
					r = s[i : j+1]
				}
			} else if s[i] == s[j] {
				m[i][j] = m[i+1][j-1] + 2
				if max(x, m[i][j]) != x {
					x = m[i][j]
					r = s[i : j+1]
				}
			} else {
				m[i][j] = max(m[i+1][j], m[i][j-1])
			}
		}
	}
	return r
}

// func findLPS_recur_with_memotable(s string, x int, y int, memo [][]int) (int, string) {
// 	if x == y {
// 		return 1, s[x : y+1]
// 	}
// 	if y == x+1 && s[x] == s[y] {
// 		if memo[x][y] != -1 {
// 			return memo[x][y]
// 		} else {
// 			memo[x][y] = 2
// 			return 2, s[x : y+1]
// 		}
// 	}
// 	if s[x] == s[y] {
// 		if memo[x+1][y-1] != -1 {
// 			memo[x][y] = memo[x+1][y-1] + 2
// 			return memo[x+1][y-1] + 2
// 		} else {
// 			t1, t2 := findLPS_recur_with_memotable(s, x+1, y-1, memo)
// 			return  + 2
// 		}
// 	} else {
// 		var m1, m2 int
// 		var s1, s2 string
// 		if memo[x+1][y] != -1 {
// 			m1 = memo[x+1][y]
// 		} else {
// 			m1, s1 = findLPS_recur_with_memotable(s, x+1, y, memo)
// 		}
// 		if memo[x][y-1] != -1 {
// 			m2 = memo[x][y-1]
// 		} else {
// 			m2, s2 = findLPS_recur_with_memotable(s, x, y-1, memo)
// 		}
// 		if max(m1, m2) == m1 {
// 			return m1, s1
// 		} else {
// 			return m2, s2
// 		}
// 	}

// }

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s := "ebacacbd"
	fmt.Println(findLPS(s, 0, len(s)-1))
	0
}
