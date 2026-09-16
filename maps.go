package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	resultMap := make(map[string]int)
	wordsSlice := strings.Fields(s)
	for _, word := range wordsSlice {
		count, ok := resultMap[word]
		if ok {
			resultMap[word] = count + 1
		} else {
			resultMap[word] = 1
		}
	}
	return resultMap
}

func main() {
	wc.Test(WordCount)
}
