package main

import (
	"fmt"
)

func isBalanced(s string) bool {
	stack := []rune{}
	bracketPairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, char := range s {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != bracketPairs[char] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	input1 := "({[]})"
	input2 := "(]"
	if isBalanced(input1) {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not Balanced")
	}
	if isBalanced(input2) {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Not Balanced")
	}
}
