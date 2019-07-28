package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
  // Split each word in s string to a slice of string
	words := strings.Fields(s)

  // Maka an empty map
	wordCountMap := make(map[string]int)

  // For each word splitted, increment the map with word value by one
	for _, word := range words {
		// Default value if not exists is 0, it is therefore safe to increment value
		// without checking if key exists
		wordCountMap[word] += 1
	}

  // Return the map of word count
	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
