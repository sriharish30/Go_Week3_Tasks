package main

import (
	"fmt"
	"strings"
)

func countWordFrequency(input string) map[string]int {

	words := strings.Fields(input)

	wordFrequency := make(map[string]int)

	for _, word := range words {
		wordFrequency[word]++
	}

	return wordFrequency
}

func main() {
	input := "Go is fun and Go is powerful"
	frequency := countWordFrequency(input)

	for word, count := range frequency {
		fmt.Printf("%s: %d, ", word, count)
	}
}
