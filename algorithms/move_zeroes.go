package algorithms

func MoveZeroes(nums []int) {
	posZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[posZero] = nums[posZero], nums[i]
			posZero++
		}

	}
}
