package main

import (
	"fmt"
	"math"
)

func secondLargest(nums []int) int {
	if len(nums) < 2 {
		return math.MinInt32
	}

	first, second := math.MinInt32, math.MinInt32

	for _, num := range nums {
		if num > first {
			second = first
			first = num
		} else if num > second && num != first {
			second = num
		}
	}

	return second
}

func main() {
	nums := []int{10, 20, 4, 45, 99}
	result := secondLargest(nums)
	fmt.Println("The second largest number is:", result)
}
