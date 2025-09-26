package algorithms

import "Codewars/tasks"

func ReverseStrings(s string) string {
	stack := tasks.NewArrayStack[rune](len([]rune(s)))
	var reversedString []rune
	for _, symbol := range s {
		stack.Push(symbol)
	}
	for !stack.IsEmpty() {
		reversedString = append(reversedString, stack.Pop())
	}
	return string(reversedString)
}
