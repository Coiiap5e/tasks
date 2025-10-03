package algorithms

import "Codewars/structures"

func IsBalanced(s string) bool {
	stack := structures.NewArrayStack[rune]()
	parenthesesStructure := map[rune]rune{
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
			if stack.Pop() != parenthesesStructure[symbol] {
				return false
			}
		default:
			return false
		}
	}
	return stack.IsEmpty()
}
