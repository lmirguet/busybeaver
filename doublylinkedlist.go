package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
	value int
	prev  *Node
	next  *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// AppendRight adds an element to the end of the list
func (dll *DoublyLinkedList) AppendRight(value int) {
	newNode := &Node{value: value}
	if dll.tail == nil { // If the list is empty
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.tail.next = newNode
		newNode.prev = dll.tail
		dll.tail = newNode
	}
}

// AppendLeft adds an element to the beginning of the list
func (dll *DoublyLinkedList) AppendLeft(value int) {
	newNode := &Node{value: value}
	if dll.head == nil { // If the list is empty
		dll.head = newNode
		dll.tail = newNode
	} else {
		dll.head.prev = newNode
		newNode.next = dll.head
		dll.head = newNode
	}
}

// Display prints the elements of the list from the beginning to the end
func (dll *DoublyLinkedList) Display(max int) {
	current := dll.head
	i := 0
	for current != nil {
		i = i + 1
		if i > max {
			break
		}
		fmt.Print(current.value, " ")
		current = current.next
	}
	fmt.Println()
}

// func main() {
// 	dll := &DoublyLinkedList{}
// 	dll.AppendRight(1)
// 	dll.AppendRight(2)
// 	dll.AppendLeft(0)
// 	dll.AppendLeft(-1)

// 	dll.Display() // Output: -1 0 1 2
// }
