package main

import (
	"bytes"
	"fmt"
)

func convert(s string, numRows int) string {
	var buffer bytes.Buffer
	for i := 0; i < numRows; i++ {
		if numRows%2 != 0 {
			if i == numRows/2 {
				for j := i; j < len(s); j += numRows/2 + 1 {
					buffer.WriteByte(s[j])
				}
			} else {
				for j := i; j < len(s); j += numRows + 1 {
					buffer.WriteByte(s[j])
				}
			}
		} else {
			for j := i; j < len(s); j += numRows {
				buffer.WriteByte(s[j])
			}
		}
	}
	return buffer.String()
}

func main() {
	fmt.Println(convert("ABC", 2))
}
