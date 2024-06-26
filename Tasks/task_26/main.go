package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]bool)

	for _, r := range str {
		if _, ok := m[r]; ok {
			return false
		}
		m[r] = true
	}

	return true
}

func main() {
	fmt.Printf("abcd: %v\n", isUnique("abcd"))
	fmt.Printf("abCdefAaf: %v\n", isUnique("abCdefAaf"))
	fmt.Printf("aabcd: %v\n", isUnique("aabcd"))
}
