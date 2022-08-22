package main

import "fmt"

func main() {
	var nums = [10]int{1, 3, 4, 6, 7, 1, 21, 32, 11, 45}
	funcResult := twoSum(nums[:], 92)
	fmt.Println(funcResult)
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	// return []int{}  or // return nil
	return nil
}
