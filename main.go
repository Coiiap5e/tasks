package main

import (
	"Codewars/tasks"
	"fmt"
)

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
		}
	}
	return stack.IsEmpty()
}

func ReverseStrings(s string) string {
	stack := tasks.NewArrayStack[rune]()
	var reversedString []rune
	for _, symbol := range s {
		stack.Push(symbol)
	}
	for !stack.IsEmpty() {
		reversedString = append(reversedString, stack.Pop())
	}
	return string(reversedString)
}

func main() {
	stack := tasks.NewArrayStack[int]()
	queue := tasks.NewArrayQueue[int]()
	set := tasks.NewHashSet[int]()
	fmt.Println("Проверка стека:")
	fmt.Println("IsEmpty ", stack.IsEmpty())
	stack.Push(1)
	stack.Push(2)
	fmt.Println("Size ", stack.Size())
	fmt.Println("Pop ", stack.Pop())
	fmt.Println("stack ", stack)
	fmt.Println("Проверка очереди:")
	fmt.Println("IsEmpty ", queue.IsEmpty())
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println("Size ", queue.Size())
	fmt.Println("IsEmpty ", queue.IsEmpty())
	fmt.Println("Dequeue ", queue.Dequeue())
	fmt.Println("Queue ", queue)
	fmt.Println("Peek ", queue.Peek())
	fmt.Println("Queue ", queue)
	fmt.Println("Проверка сета:")
	fmt.Println("IsEmpty ", set.IsEmpty())
	set.Add(5)
	set.Add(6)
	set.Add(5)
	fmt.Println("Size ", set.Size())
	fmt.Println("IsEmpty ", set.IsEmpty())
	slice := set.ToSlice()
	fmt.Println("slice ", slice)
	fmt.Println("Contains ", set.Contains(5))
	set.Remove(5)
	fmt.Println("Contains ", set.Contains(5))
	fmt.Println(set)
	set.Clear()
	fmt.Println(set)
	fmt.Println("Проверка скобок:")
	fmt.Println(IsBalanced("({[]})"))
	fmt.Println(IsBalanced("([)]"))
	fmt.Println(IsBalanced("{(})"))
	fmt.Println(IsBalanced(""))
	fmt.Println("Проверка реверса строки:")
	fmt.Println(ReverseStrings("Hello World!"))
	fmt.Println(tasks.NaiveSearch("абвгдейка", "где"))
}
