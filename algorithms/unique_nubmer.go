package algorithms

func UniqueOccurrences(arr []int) bool {
	repeat := make(map[int]int)
	count := make(map[int]bool)
	for _, value := range arr {
		repeat[value]++
	}
	for _, value := range repeat {
		if _, exists := count[value]; exists {
			return false
		} else {
			count[value] = true
		}
	}
	return true
}
