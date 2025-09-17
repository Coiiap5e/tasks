package tasks

type Set interface {
	Add(value interface{})           // Добавление элемента в множество
	Remove(value interface{})        // Удаление элемента из множества
	Contains(value interface{}) bool // Проверка наличия элемента
	IsEmpty() bool                   // Проверка на пустоту
	Size() int                       // Возврат количества элементов
	Clear()                          // Очистка множества
	ToSlice() []interface{}          // Преобразование в срез
}

type HashSet struct {
	elements map[interface{}]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[interface{}]bool)}
}

func (s *HashSet) Add(value interface{}) {
	s.elements[value] = true
}

func (s *HashSet) Remove(value interface{}) {
	delete(s.elements, value)
}

func (s *HashSet) Contains(value interface{}) bool {
	_, ok := s.elements[value]
	return ok
}

func (s *HashSet) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *HashSet) Size() int {
	return len(s.elements)
}

func (s *HashSet) Clear() {
	clear(s.elements)
}

func (s *HashSet) ToSlice() []interface{} {
	slice := make([]interface{}, 0, len(s.elements))
	for value := range s.elements {
		slice = append(slice, value)
	}
	return slice
}
