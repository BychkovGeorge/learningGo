package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("result", mergeAlternately("abc", "def"))
}

func mergeAlternately(word1 string, word2 string) string {
	var sb strings.Builder
	for i := 0; i < max(len(word1), len(word2)); i++ {
		if i < len(word1) {
			sb.WriteByte(word1[i])
		}
		if i < len(word2) {
			sb.WriteByte(word2[i])
		}
	}
	return sb.String()
}
