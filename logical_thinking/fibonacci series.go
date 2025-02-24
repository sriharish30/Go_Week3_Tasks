package main

import (
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciSeries(n int) []int {
	series := make([]int, n)
	for i := 0; i < n; i++ {
		series[i] = fibonacci(i)
	}
	return series
}

func main() {
	N := 7
	series := fibonacciSeries(N)
	for i, num := range series {
		if i != len(series)-1 {
			fmt.Printf("%d, ", num)
		} else {
			fmt.Printf("%d\n", num)
		}
	}
}
