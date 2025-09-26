package main

import (
	"Codewars/algorithms"
	"Codewars/tasks"
	"fmt"
)

func showStack[T any](stack *tasks.ArrayStack[T], values ...T) {
	fmt.Println("Проверка стека:")
	fmt.Println("IsEmpty: ", stack.IsEmpty())
	for _, value := range values {
		stack.Push(value)
	}
	fmt.Printf("Size после добавления %v: %v \n", values, stack.Size())
	fmt.Println("Удаление последнего элемента Pop(): ", stack.Pop())
	fmt.Println("stack: ", stack)
}

func showQueue[T any](queue *tasks.ArrayQueue[T], values ...T) {
	fmt.Println("Проверка очереди:")
	fmt.Println("IsEmpty: ", queue.IsEmpty())
	for _, value := range values {
		queue.Enqueue(value)
	}
	fmt.Printf("Size после добавления %v: %v \n", values, queue.Size())
	fmt.Println("IsEmpty: ", queue.IsEmpty())
	fmt.Println("Удаление первого элемента Dequeue(): ", queue.Dequeue())
	fmt.Println("Queue: ", queue)
	fmt.Println("Возвращение 1го элемента Peek(): ", queue.Peek())
	fmt.Println("Queue: ", queue)
}

func showSet[T comparable](set *tasks.HashSet[T], deleteValue T, values ...T) {
	fmt.Println("Проверка сета:")
	fmt.Println("IsEmpty: ", set.IsEmpty())
	for _, value := range values {
		set.Add(value)
	}
	fmt.Printf("Size после добавления %v: %v \n", values, set.Size())
	fmt.Println("IsEmpty: ", set.IsEmpty())
	slice := set.ToSlice()
	fmt.Println("Преобразование в slice: ", slice)
	fmt.Printf("Проверка наличия элемента %v: %v \n", deleteValue, set.Contains(deleteValue))
	fmt.Println("Remove: ", deleteValue)
	set.Remove(deleteValue)
	fmt.Printf("Проверка наличия элемента %v: %v \n", deleteValue, set.Contains(deleteValue))
	fmt.Println("Set: ", set)
	set.Clear()
	fmt.Println("Очистка сета")
	fmt.Println("Set: ", set)
}

func showAlgorithmIsBalanced() {
	fmt.Println("Алгоритм проверки скобок:")
	fmt.Println("({[]})", algorithms.IsBalanced("({[]})"))
	fmt.Println("([)]", algorithms.IsBalanced("([)]"))
	fmt.Println("{(})", algorithms.IsBalanced("{(})"))
	fmt.Println("Пустая строка", algorithms.IsBalanced(""))
	fmt.Println("({[text]})", algorithms.IsBalanced("({[text]})"))
}

func showAlgorithmReverseString() {
	fmt.Println("Алгоритм реверса строки Hello World:")
	fmt.Println(algorithms.ReverseStrings("Hello World!"))
}

func showAlgorithmNaiveSearch(text, target string) {
	fmt.Printf("Алгортим поиска подстроки '%s', в тексте '%s' : \n", target, text)
	fmt.Println(algorithms.NaiveSearch(text, target))
}

func main() {
	stack := tasks.NewArrayStack[int]()
	queue := tasks.NewArrayQueue[int]()
	set := tasks.NewHashSet[int]()
	showStack(stack, 1, 2, 3)
	showQueue(queue, 4, 5, 6)
	showSet(set, 5, 6, 5, 6, 8, 8)
	showAlgorithmIsBalanced()
	showAlgorithmReverseString()
	showAlgorithmNaiveSearch("абвгдейка", "где")
	algorithms.RunAlgorithm()
}
