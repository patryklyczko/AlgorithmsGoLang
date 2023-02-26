package linked

import "fmt"

type Linked interface {
	Append(value interface{})
	Prepend(value interface{})
	Delete(val interface{})
	Show()
}

type Node struct {
	item interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Append(value interface{}) {
	if l.head == nil {
		l.head = &Node{
			item: value,
			next: nil,
		}
	} else {
		actualNode := l.head
		for actualNode.next != nil {
			actualNode = actualNode.next
		}
		actualNode.next = &Node{item: value, next: nil}
	}
}

func (l *List) Show() {
	headNode := l.head
	fmt.Print(headNode.item, " -> ")
	for headNode.next != nil {
		headNode = headNode.next
		fmt.Print(headNode.item, " -> ")
	}
	fmt.Print("\n")
}

func (l *List) Prepend(val interface{}) {
	head := l.head
	l.head = &Node{item: val, next: head}
}

func (l *List) Delete(val interface{}) {
	if l.head.item == val {
		l.head = l.head.next
	}
	actualNode := l.head
	for actualNode.next != nil {
		if actualNode.next.item == val {
			actualNode.next = actualNode.next.next
		}
		actualNode = actualNode.next
	}
}

func LinkedList() Linked {
	return &List{head: nil}
}
