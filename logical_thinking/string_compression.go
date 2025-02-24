package main

import (
	"fmt"
	"strconv"
)

func compressString(input string) string {
	if len(input) == 0 {
		return ""
	}

	compressed := ""
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			compressed += string(input[i-1]) + strconv.Itoa(count)
			count = 1
		}
	}

	compressed += string(input[len(input)-1]) + strconv.Itoa(count)

	return compressed
}

func main() {
	input := "aaabbcddd"
	output := compressString(input)
	fmt.Println(output)
}
