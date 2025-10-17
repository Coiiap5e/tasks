package structures

type Queue[T any] interface {
	Enqueue(value T) // Добавление элемента в конец очереди
	Dequeue() T      // Удаление и возврат элемента из начала очереди
	Peek() T         // Возврат элемента из начала без удаления
	IsEmpty() bool   // Проверка на пустоту
	Size() int       // Возврат количества элементов
}

type ArrayQueue[T any] struct {
	elements []T
}

func NewArrayQueue[T any](cap ...int) *ArrayQueue[T] {
	capacity := 0
	if len(cap) > 0 {
		capacity = cap[0]
	}
	return &ArrayQueue[T]{make([]T, 0, capacity)}
}

func (q *ArrayQueue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

func (q *ArrayQueue[T]) Dequeue() T {
	if q.IsEmpty() {
		var zero T
		return zero
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *ArrayQueue[T]) Peek() T {
	if q.IsEmpty() {
		var zero T
		return zero
	}
	return q.elements[0]
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *ArrayQueue[T]) Size() int {
	return len(q.elements)
}
