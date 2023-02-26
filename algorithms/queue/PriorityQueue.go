package queue

type interfacePriorityQueue interface {
	Push(value interface{})
	PullAsc() interface{}
	PullDesc() interface{}
	Delete()
}

type PriorityQueue struct {
	items []interface{}
}

func (p *PriorityQueue) Push(value interface{}) {
	if len(p.items) == 0 {
		p.items = append(p.items, value)
	} else {
		// sort method
	}
}

func (p *PriorityQueue) PullAsc() interface{} {
	if len(p.items) != 0 {
		item := p.items[0]
		p.items = p.items[1:]
		return item
	}
	return nil
}

func (p *PriorityQueue) PullDesc() interface{} {
	if len(p.items) != 0 {
		length := len(p.items)
		item := p.items[length-1]
		p.items = p.items[:length-1]
		return item
	}
	return nil
}
