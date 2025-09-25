package algorithms

import (
	"strings"
	"unicode"
)

func IsPalindromeUppercase(s string) bool {
	s = strings.ToLower(s)
	var runeStr []rune
	for _, char := range s {
		if unicode.IsLetter(char) {
			runeStr = append(runeStr, char)
		}
	}
	return isPalindrome(runeStr, 0, len(runeStr)-1)
}

func IsPalindromeWithoutOneLetter(s string) bool {
	runeStr := []rune(s)
	i, j := 0, len(runeStr)-1
	for i < j {
		if runeStr[i] != runeStr[j] {
			return isPalindrome(runeStr, i+1, j) || isPalindrome(runeStr, i, j-1)
		}
		i++
		j--
	}
	return true
}

func isPalindrome(runes []rune, i, j int) bool {
	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
