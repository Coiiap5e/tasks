package main

import (
	"Codewars/algorithms"
	"Codewars/structures"
	"Codewars/tasks"
	"context"
	"fmt"
	"time"
)

func showStack[T any](stack *structures.ArrayStack[T], values ...T) {
	fmt.Println("Проверка стека:")
	fmt.Println("IsEmpty: ", stack.IsEmpty())
	for _, value := range values {
		stack.Push(value)
	}
	fmt.Printf("Size после добавления %v: %v \n", values, stack.Size())
	fmt.Println("Удаление последнего элемента Pop(): ", stack.Pop())
	fmt.Println("stack: ", stack)
}

func showQueue[T any](queue *structures.ArrayQueue[T], values ...T) {
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

func showSet[T comparable](set *structures.HashSet[T], deleteValue T, values ...T) {
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

func main() {
	stack := structures.NewArrayStack[int]()
	queue := structures.NewArrayQueue[int]()
	set := structures.NewHashSet[int]()
	incCount, goroutineCount := 10000, 1000

	showStack(stack, 1, 2, 3)

	showQueue(queue, 4, 5, 6)

	showSet(set, 5, 6, 5, 6, 8, 8)

	// Заглушки ф-ций
	algorithms.RunAlgorithm()

	_ = tasks.RunAtomicCounter(incCount, goroutineCount)

	_ = tasks.RunMutexCounter(incCount, goroutineCount)

	//Проверка контекста

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		fmt.Println("CANCELLING CONTEXT!")
		cancel() // Отменяем операцию через 4 секунды
	}()
	dataArray := [][]string{
		{"data1", "data2", "data3"},
		{"data4", "data5", "data6"},
		{"data7", "data8", "data9"},
	}

	channelsArray := make([]chan tasks.Telemetry, len(dataArray))

	for i := 0; i < len(dataArray); i++ {
		channelsArray[i] = make(chan tasks.Telemetry, len(dataArray[i]))
	}

	for i, channel := range channelsArray {
		go func(index int, channel chan tasks.Telemetry) {
			defer close(channel)

			for _, value := range dataArray[index] {
				time.Sleep(1 * time.Second)
				channel <- tasks.Telemetry{SensorData: value}
				fmt.Printf("Sent to channel: %s (message %d)\n", value, index+1)
			}
		}(i, channel)
	}

	onlyReadChannels := make([]<-chan tasks.Telemetry, len(dataArray))
	for i, channel := range channelsArray {
		onlyReadChannels[i] = channel
	}

	mergedChan := tasks.MergeTelemetry(ctx, onlyReadChannels...)

	for data := range mergedChan {
		fmt.Println("Data: ", data)
	}
}
