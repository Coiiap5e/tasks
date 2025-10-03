package algorithms

/*
Given an array of integers nums and an integer target,
return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution,
and you may not use the same element twice.

You can return the answer in any order
*/

func TwoSum(nums []int, target int) []int {
	diff := make(map[int]int)
	for index, value := range nums {
		diff[target-value] = index
		if result, exists := diff[value]; exists && result != index {
			return []int{result, index}
		}
	}
	return []int{}
}
