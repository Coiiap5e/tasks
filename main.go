package main

import (
	"Codewars/tasks"
	"fmt"
)

func IsBalanced(s string) bool {
	stack := tasks.NewArrayStack()
	map1 := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, symbol := range []rune(s) {
		switch symbol {
		case '[', '{', '(':
			stack.Push(symbol)
		case ']', '}', ')':
			if stack.IsEmpty() {
				return false
			}
			last, ok := stack.Pop().(rune)
			if !ok {
				return false
			}
			if last != map1[symbol] {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func ReverseStrings(s string) string {
	stack := tasks.NewArrayStack()
	var reversedString []rune
	for _, symbol := range []rune(s) {
		stack.Push(symbol)
	}
	for !stack.IsEmpty() {
		reversedString = append(reversedString, stack.Pop().(rune))
	}
	return string(reversedString)
}

func main() {
	stack := tasks.NewArrayStack()
	queue := tasks.NewArrayQueue()
	set := tasks.NewHashSet()
	fmt.Println("Проверка стека:")
	fmt.Println(stack.IsEmpty())
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Size())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
	fmt.Println("Проверка очереди:")
	fmt.Println(queue.IsEmpty())
	queue.Enqueue(3)
	queue.Enqueue(4)
	fmt.Println(queue.Size())
	fmt.Println(queue.IsEmpty())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue)
	fmt.Println(queue.Peek())
	fmt.Println(queue)
	fmt.Println("Проверка сета:")
	fmt.Println(set.IsEmpty())
	set.Add(5)
	set.Add(6)
	set.Add(5)
	fmt.Println(set.Size())
	fmt.Println(set.IsEmpty())
	slice := set.ToSlice()
	fmt.Println(slice)
	fmt.Println(set.Contains(5))
	set.Remove(5)
	fmt.Println(set.Contains(5))
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
