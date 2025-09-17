package tasks

type Queue interface {
	Enqueue(value interface{}) // Добавление элемента в конец очереди
	Dequeue() interface{}      // Удаление и возврат элемента из начала очереди
	Peek() interface{}         // Возврат элемента из начала без удаления
	IsEmpty() bool             // Проверка на пустоту
	Size() int                 // Возврат количества элементов
}

type ArrayQueue struct {
	elements []interface{}
}

func NewArrayQueue() *ArrayQueue {
	return &ArrayQueue{make([]interface{}, 0)}
}

func (q *ArrayQueue) Enqueue(value interface{}) {
	q.elements = append(q.elements, value)
}

func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

func (q *ArrayQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.elements[0]
}

func (q *ArrayQueue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *ArrayQueue) Size() int {
	return len(q.elements)
}
