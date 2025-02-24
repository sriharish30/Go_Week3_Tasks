package main

import "fmt"

func find_duplicates(nums []int) {
	seen := make(map[int]bool)
	duplicate := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			duplicate[num] = true

		} else {
			seen[num] = true
		}

	}
	for num := range duplicate {
		fmt.Println("duplicates:", num)
	}

}
func main() {
	nums := []int{4, 5, 6, 4, 6, 7}
	find_duplicates(nums)

}
