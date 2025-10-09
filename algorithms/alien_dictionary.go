package algorithms

func IsAlienSorted(words []string, order string) bool {
	charOrder := make(map[rune]int)
	index := 0

	for _, char := range order {
		charOrder[char] = index
		index++
	}

	for i := 0; i < len(words)-1; i++ {
		if !compareWords(words[i], words[i+1], charOrder) {
			return false
		}
	}

	return true
}

func compareWords(firstWord string, secondWord string, charOrder map[rune]int) bool {
	for j := 0; j < len([]rune(firstWord)) && j < len([]rune(secondWord)); j++ {
		if firstWord[j] != secondWord[j] {
			return charOrder[rune(firstWord[j])] < charOrder[rune(secondWord[j])]
		}
	}

	return len([]rune(firstWord)) <= len([]rune(secondWord))
}
