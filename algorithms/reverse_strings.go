package algorithms

import "Codewars/structures"

func ReverseStrings(s string) string {
	stack := structures.NewArrayStack[rune](len([]rune(s)))
	var reversedString []rune
	for _, symbol := range s {
		stack.Push(symbol)
	}
	for !stack.IsEmpty() {
		reversedString = append(reversedString, stack.Pop())
	}
	return string(reversedString)
}
