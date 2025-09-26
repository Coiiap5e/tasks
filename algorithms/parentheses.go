package algorithms

import "Codewars/tasks"

func IsBalanced(s string) bool {
	stack := tasks.NewArrayStack[rune]()
	map1 := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, symbol := range s {
		switch symbol {
		case '[', '{', '(':
			stack.Push(symbol)
		case ']', '}', ')':
			if stack.IsEmpty() {
				return false
			}
			last := stack.Pop()
			if last != map1[symbol] {
				return false
			}
		default:
			return false
		}
	}
	return stack.IsEmpty()
}
