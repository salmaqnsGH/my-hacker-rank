package main

import "fmt"

type Queue struct {
	items []int32
}

// Enqueue will adds value at the end
func (q *Queue) Enqueue(num int32) {
	q.items = append(q.items, num)
}

// Dequeue removes value at the front
func (q *Queue) Dequeue() int32 {
	toRemove := q.items[0]
	q.items = q.items[1:]

	return toRemove
}

func main() {

	myQueue := Queue{}

	fmt.Println(myQueue)

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)

	fmt.Println(myQueue)

	myQueue.Dequeue()

	fmt.Println(myQueue)
}
