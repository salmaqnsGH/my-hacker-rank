package main

import "fmt"

type Stack struct {
	items []int32
}

// Push adds value at the end
func (s *Stack) Push(num int32) {
	s.items = append(s.items, num)
}

// Pop removes value at the end
func (s *Stack) Pop() int32 {
	l := len(s.items) - 1

	toRemove := s.items[l]
	s.items = s.items[:l]

	return toRemove
}

func main() {
	myStack := Stack{}

	fmt.Println(myStack)

	myStack.Push(100)
	myStack.Push(101)
	myStack.Push(102)

	fmt.Println(myStack)

	myStack.Pop()

	fmt.Println(myStack)
}
