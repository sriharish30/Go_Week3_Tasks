package main

import (
	"fmt"
)

func reverseWords(sentence string) string {
	words := []string{}
	word := ""
	for _, char := range sentence {
		if char == ' ' {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(char)
		}
	}
	if word != "" {
		words = append(words, word)
	}

	reversedSentence := ""
	for i := len(words) - 1; i >= 0; i-- {
		reversedSentence += words[i]
		if i > 0 {
			reversedSentence += " "
		}
	}

	return reversedSentence
}

func main() {
	input := "Go is fun"
	output := reverseWords(input)
	fmt.Println(output)
}
