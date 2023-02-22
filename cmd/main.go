package main

import (
	que "Algorithms/algorithms/queues.go"
	"fmt"
)

func main() {
	q := que.FIFOQueue()
	q.Push(1)
	q.push(2)
	q.push(3)
	fmt.Print(q.pop())
}
