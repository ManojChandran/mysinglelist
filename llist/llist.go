package llist

import "fmt"

type node struct {
	data interface{}
	next *node
}

func getNode(value interface{}) *node {
	return &node{
		data: value,
	}
}

// Llist is a sturct defined for linked list
type Llist struct {
	head *node
	tail *node
}

// Getlist function will return the list
func Getlist() *Llist {
	return &Llist{}
}

// Check list status
func (l Llist) isEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

// ListAll will list all the values in linked list
func (l Llist) ListAll() {
	for l.tail = l.head; l.tail != nil; l.tail = l.tail.next {
		fmt.Println(l.tail.data)
	}
}

// Add function will add a value to the list
func (l *Llist) Add(value interface{}) bool {
	newnode := getNode(value)
	if l.isEmpty() == true {
		l.head = newnode
		l.tail = newnode
		return true
	}
	for l.tail.next != nil {
		l.tail = l.tail.next
	}
	if l.tail.next == nil {
		l.tail.next = newnode
		return true
	}
	return false
}

// Delete function will delete a value form the linked list
func (l *Llist) Delete(value interface{}) bool {
	prev := l.head
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.data == value {
			l.tail = prev
			l.tail.next = curr.next
			return true
		}
		prev = curr
	}
	return false
}

// Insert function will insert a value into the list
func (l *Llist) Insert(value, after interface{}) bool {
	newnode := getNode(value)
	temp := l.head
	for curr := l.head; curr != nil; curr = curr.next {
		if curr.data == after {
			temp = curr.next
			curr.next = newnode
			curr = curr.next
			curr.next = temp
			return true
		}
	}
	return false
}
