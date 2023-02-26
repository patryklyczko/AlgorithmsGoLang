package queue

type LIFO struct {
	items []interface{}
}

func (l *LIFO) Push(val interface{}) {
	l.items = append(l.items, val)
}

func (l *LIFO) Pop() interface{} {
	length := len(l.items)
	item := l.items[length-1]
	l.items = l.items[:length-1]
	return item
}

func (l *LIFO) Peek() interface{} {
	length := len(l.items)
	item := l.items[length-1]
	return item
}

func (l *LIFO) Delete() {
	l.items = []interface{}{}
}

func LIFOQueue() Queue {
	return &FIFO{
		items: make([]interface{}, 0),
	}
}
