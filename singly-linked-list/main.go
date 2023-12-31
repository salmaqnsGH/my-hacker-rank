package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	temp := l.head

	l.head = n
	l.head.next = temp
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)

		toPrint = toPrint.next
		l.length--
	}

	fmt.Println("\nfinish printing data")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func main() {
	myList := linkedList{}

	node1 := &node{data: 48}
	node2 := &node{data: 44}
	node3 := &node{data: 23}
	node4 := &node{data: 11}
	node5 := &node{data: 52}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)

	fmt.Println(myList)

	myList.printListData()
	myList.deleteWithValue(41)
	myList.printListData()
}
