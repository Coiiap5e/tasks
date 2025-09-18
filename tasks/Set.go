package tasks

type Set[T comparable] interface {
	Add(value T)           // Добавление элемента в множество
	Remove(value T)        // Удаление элемента из множества
	Contains(value T) bool // Проверка наличия элемента
	IsEmpty() bool         // Проверка на пустоту
	Size() int             // Возврат количества элементов
	Clear()                // Очистка множества
	ToSlice() []T          // Преобразование в срез
}

type HashSet[T comparable] struct {
	elements map[T]bool
}

func NewHashSet[T comparable]() *HashSet[T] {
	return &HashSet[T]{make(map[T]bool)}
}

func (s *HashSet[T]) Add(value T) {
	s.elements[value] = true
}

func (s *HashSet[T]) Remove(value T) {
	delete(s.elements, value)
}

func (s *HashSet[T]) Contains(value T) bool {
	_, ok := s.elements[value]
	return ok
}

func (s *HashSet[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *HashSet[T]) Size() int {
	return len(s.elements)
}

func (s *HashSet[T]) Clear() {
	clear(s.elements)
}

func (s *HashSet[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.elements))
	for value := range s.elements {
		slice = append(slice, value)
	}
	return slice
}
