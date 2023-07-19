package linkedlist

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

import (
	"log"
	"errors"
)

type Node struct {
	Value int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...interface{}) *List {
	list := new(List)
	if elements == nil {
		return list
	}

	for _, el := range elements {
		list.Push(el)
	}

	return list
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	switch v.(type) {
	case int:
		if l.head == nil {
			l.head = new(Node)
			l.head.Value = v.(int)
			l.tail = l.head
		} else {
			item := new(Node)
			item.Value = v.(int)
			item.next = l.head
			l.head.prev = item
			l.head = item
		}
	default:
		log.Fatalf("Unknown type. Got: %T", v)
	}
}

func (l *List) Push(v interface{}) {
	switch v.(type) {
	case int:
		if l.head == nil {
			l.head = new(Node)
			l.head.Value = v.(int)
			l.tail = l.head
		} else {
			item := new(Node)
			item.Value = v.(int)
			item.prev = l.tail
			l.tail.next = item
			l.tail = item
		}
	default:
		log.Fatalf("Unknown type. Got: %T", v)
	}	
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("List is empty")
	}

	if l.head == l.tail {
		l.tail = nil
	}

	item := l.head
	l.head = item.next

	if l.head != nil {
		l.head.prev = nil
	}

	return item.Value, nil
}

func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("List is empty")
	}

	if l.head == l.tail {
		l.head = nil
	}

	item := l.tail
	l.tail = item.prev

	if l.tail != nil {
		l.tail.next = nil
	}

	return item.Value, nil
}

func (l *List) Reverse() {
	for item := l.head; item != nil; {
		next := item.next
		item.next = item.prev
		item.prev = next
		item = item.prev
	}

	head := l.head
	l.head = l.tail
	l.tail = head
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}
