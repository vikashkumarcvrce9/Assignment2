package main

import "fmt"

// Node Struct for pointing
type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

// List struct for head and tail
type List struct {
	head *Node
	tail *Node
}

// Insert function for insert the value into linkedlist
func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head

	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

// Display func for displaying the head and tail
func (l *List) Display() {
	list := l.head

	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// Display
func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

// ShowBackwards pointing towords backwards
func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev

	}
	fmt.Println()
}

// Reverse function
func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next

	}
	l.head = prev
	Display(l.head)
}

func main() {
	list := List{}
	list.Insert(5)
	list.Insert(9)
	list.Insert(13)
	list.Insert(22)
	list.Insert(28)
	list.Insert(36)

	fmt.Println("\n----------------------------\n")
	fmt.Printf("Head:%v\n", list.head.key)
	fmt.Printf("Tail:%v\n", list.tail.key)

	list.Display()

	fmt.Println("\n-------------------------------\n")
	fmt.Printf("Head:%v\n", list.head.key)
	fmt.Printf("Tail:%v\n", list.tail.key)

	list.Reverse()

}
