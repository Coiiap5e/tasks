package tasks

type Stack interface {
	Push(value interface{}) // Добавление элемента на вершину стека
	Pop() interface{}       // Удаление и возврат элемента с вершины стека
	Peek() interface{}      // Возврат элемента с вершины без удаления
	IsEmpty() bool          // Проверка на пустоту
	Size() int              // Возврат количества элементов
}

type ArrayStack struct {
	elements []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{elements: make([]interface{}, 0)}
}

func (s *ArrayStack) Push(value interface{}) {
	s.elements = append(s.elements, value)
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value
}

func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.elements[len(s.elements)-1]
}

func (s *ArrayStack) IsEmpty() bool {
	return len(s.elements) == 0

}

func (s *ArrayStack) Size() int {
	return len(s.elements)
}
