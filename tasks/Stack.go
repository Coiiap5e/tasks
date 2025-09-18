package tasks

type Stack[T any] interface {
	Push(value T)  // Добавление элемента на вершину стека
	Pop() T        // Удаление и возврат элемента с вершины стека
	Peek() T       // Возврат элемента с вершины без удаления
	IsEmpty() bool // Проверка на пустоту
	Size() int     // Возврат количества элементов
}

type ArrayStack[T any] struct {
	elements []T
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{elements: make([]T, 0)}
}

func (s *ArrayStack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *ArrayStack[T]) Pop() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

func (s *ArrayStack[T]) Peek() T {
	if s.IsEmpty() {
		var zero T
		return zero
	}
	return s.elements[len(s.elements)-1]
}

func (s *ArrayStack[T]) IsEmpty() bool {
	return len(s.elements) == 0

}

func (s *ArrayStack[T]) Size() int {
	return len(s.elements)
}
