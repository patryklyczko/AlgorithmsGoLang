package queue

type Queue interface {
	Push(value interface{})
	Pop() interface{}
}

type FIFO struct {
	items []interface{}
}

func (i *FIFO) Push(value interface{}) {
	i.items = append(i.items, value)
}

func (i *FIFO) Pop() interface{} {
	if len(i.items) != 0 {
		item := i.items[0]
		i.items = i.items[1:]
		return item
	}
	return nil
}

func FIFOQueue() Queue {
	return &FIFO{
		items: make([]interface{}, 0),
	}
}
