package main

import "fmt"

func isSubsetSum(nums []int, n, target int) bool {

	if target == 0 {
		return true
	}

	if n == 0 {
		return false
	}

	if nums[n-1] > target {
		return isSubsetSum(nums, n-1, target)
	}

	return isSubsetSum(nums, n-1, target) || isSubsetSum(nums, n-1, target-nums[n-1])
}

func main() {
	nums := []int{3, 34, 4, 12, 5, 2}
	target := 9
	n := len(nums)

	if isSubsetSum(nums, n, target) {
		fmt.Println("Subset found!")
	} else {
		fmt.Println("No subset found.")
	}
}
