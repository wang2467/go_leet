package main

import (
	"fmt"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return -100
	}
	if beginWord == endWord {
		return 0
	}
	for i := 0; i < len(wordList); i++ {
		if isValid(endWord, wordList[i]) {
			wordList = append(wordList[:i], wordList[i+1:])
			s := ladderLength(beginWord, wordList[i], wordList)
			if s != -100 {
				return s + 1
			}
		}
	}
	return -100
}

func isValid(s1 string, s2 string) bool {
	ct := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			ct++
		}
		if ct > 1 {
			return false
		}
	}
	return true
}

func main() {
	beginword := "hit"
	endword := "cog"
	wordList := []string{"hit", "dot", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginword, endword, wordList))
}
