package queue

type Queue interface {
	Push(value interface{})
	Pop() interface{}
	Peek() interface{}
	Delete()
}

type FIFO struct {
	items []interface{}
}

func (f *FIFO) Push(value interface{}) {
	f.items = append(f.items, value)
}

func (f *FIFO) Pop() interface{} {
	if len(f.items) != 0 {
		item := f.items[0]
		f.items = f.items[1:]
		return item
	}
	return nil
}

func (f *FIFO) Peek() interface{} {
	if len(f.items) != 0 {
		item := f.items[0]
		return item
	}
	return nil
}

func (f *FIFO) Delete() {
	f.items = []interface{}{}
}

func FIFOQueue() Queue {
	return &FIFO{
		items: make([]interface{}, 0),
	}
}
