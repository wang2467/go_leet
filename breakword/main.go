package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for i := 0; i < len(wordDict); i++ {
		if matches(s, wordDict[i]) {
			t := wordDict[i]
			wordlength := len(wordDict[i])
			wordDict = append(wordDict[:i], wordDict[i+1:]...)
			s = s[wordlength:]
			if !wordBreak(s, wordDict) {
				s = t + s
				if i == 0 {
					wordDict = append([]string{t}, wordDict[i:]...)
				} else {
					tmp := make([]string, len(wordDict[:i]))
					copy(tmp, wordDict[:i])
					tmp = append(tmp, t)
					wordDict = append(tmp, wordDict[i:]...)
				}
			} else {
				return true
			}
		}
	}
	return false
}

func matches(s1 string, s2 string) bool {
	if len(s1) < len(s2) {
		return false
	}
	for i := 0; i < len(s2); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func main() {
	s := "bb"
	wordDict := []string{"a", "b", "bbb", "bbbb"}
	fmt.Println(wordBreak(s, wordDict))
}
