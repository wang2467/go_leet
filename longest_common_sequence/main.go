package main

import (
	"fmt"
)

func LCS(s1 string, s2 string) int {
	if len(s1) == 1 && len(s2) == 1 && s1 == s2 {
		return 1
	} else if len(s1) == 1 && len(s2) == 1 {
		return 0
	} else if len(s1) == 1 {
		if s1[0] == s2[len(s2)-1] {
			return LCS(s1, s2[0:len(s2)-1]) + 1
		}
		return LCS(s1, s2[0:len(s2)-1])
	} else if len(s2) == 1 {
		if s2[0] == s1[len(s1)-1] {
			return LCS(s1[0:len(s1)-1], s2)
		}
		return LCS(s1[0:len(s1)-1], s2)
	} else if s1[len(s1)-1] == s2[len(s2)-1] {
		return LCS(s1[0:len(s1)-1], s2[0:len(s2)-1]) + 1
	} else {
		return max(LCS(s1[0:len(s1)-1], s2), LCS(s1, s2[0:len(s2)-1]))
	}
}

func LCS_iter(s1 string, s2 string) int {
	mat := make([][]int, len(s1)+1)
	for idx, _ := range mat {
		mat[idx] = make([]int, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 || j == 0 {
				mat[i][j] = 0
			} else if s1[i-1] == s2[j-1] {
				mat[i][j] = mat[i-1][j-1] + 1
			} else {
				mat[i][j] = max(mat[i-1][j], mat[i][j-1])
			}
		}
	}
	return mat[len(s1)][len(s2)]
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	s1 := "ABDGXXAED"
	s2 := "CCCECADGHFGXAE"

	//fmt.Println(LCS(s1, s2))
	fmt.Println(LCS_iter(s1, s2))
}
