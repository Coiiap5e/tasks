package algorithms

func IsAnagram(text string, target string) bool {
	lenText := len([]rune(text))
	lenTarget := len([]rune(target))
	if lenText != lenTarget {
		return false
	}
	counter := make(map[rune]int)
	for _, char := range text {
		counter[char]++
	}
	for _, char := range target {
		counter[char]--
		if counter[char] < 0 {
			return false
		}
	}
	for _, count := range counter {
		if count != 0 {
			return false
		}
	}
	return true
}
