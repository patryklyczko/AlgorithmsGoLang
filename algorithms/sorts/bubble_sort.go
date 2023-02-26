package sort

import "fmt"

type Sorted interface {
	// Check for type needed: works for int, add string
	Sort()
	Length() interface{}
	Add(values []interface{})
	Show()
}

type Array struct {
	items []interface{}
}

func (a *Array) Length() interface{} {
	return len(a.items)
}

func (a *Array) Sort() {
	for i := 0; i < a.Length().(int); i++ {
		for j := 0; j < a.Length().(int)-i-1; j++ {
			if a.items[j+1].(int) < a.items[j].(int) {
				val := a.items[j]
				a.items[j] = a.items[j+1]
				a.items[j+1] = val
			}
		}
	}

}

func (a *Array) Add(values []interface{}) {
	a.items = values
}

func (a *Array) Show() {
	for i := 0; i < a.Length().(int); i++ {
		fmt.Print(a.items[i], " ")
	}
	fmt.Print("\n")
}

func BubbleSort() Sorted {
	return &Array{
		items: make([]interface{}, 0),
	}
}
