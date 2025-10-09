package algorithms

func RemoveDuplicates(nums []int) int {
	uniqueNumbers := make(map[int]bool)
	nonUniquePos := -1
	for i, num := range nums {
		if !uniqueNumbers[num] {
			uniqueNumbers[num] = true
			nonUniquePos++
		}
		nums[nonUniquePos], nums[i] = nums[i], nums[nonUniquePos]
	}
	return len(uniqueNumbers)
}
