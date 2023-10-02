package linkedlist

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Add(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}

func (l *List) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		return
	}

	curr := l.head
	for curr.next != nil && curr.next.data != value {
		curr = curr.next
	}

	if curr.next != nil {
		curr.next = curr.next.next
	}
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node = nil
	var next *Node = nil

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next

	}
	l.head = prev
}
func (l *List) PrintList() {
	curr := l.head
	for curr != nil {
		fmt.Printf("%d ", curr.data)
		curr = curr.next
	}
	fmt.Println()
}
