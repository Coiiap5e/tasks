package tasks

func NaiveSearch(text, target string) int {
	// тест
	runeText := []rune(text)
	runeTarget := []rune(target)
	lenText := len(runeText)
	lenTarget := len(runeTarget)
	if lenTarget == 0 {
		return 0
	}
	if lenText < lenTarget {
		return -1
	}
	for i := 0; i <= lenText-lenTarget; i++ {
		j := 0
		for j < lenTarget && runeText[i+j] == runeTarget[j] {
			j++
		}
		if j == lenTarget {
			return i
		}
	}
	return -1
}
