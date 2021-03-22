package main

import (
	"fmt"
	"sync"
)

type Item interface{}

type Stack struct {
	items []Item
	mutex sync.Mutex
}

func (stack *Stack) Push(item Item) {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = append(stack.items, item)
}

func (stack *Stack) Pop() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	//check if the lenght of stack
	if len(stack.items) == 0 {
		return nil
	}

	//If the stack is non-empty, we copy the last item from the stack into a new variable
	lastItem := stack.items[len(stack.items)-1]
	//re-slices the stack
	stack.items = stack.items[:len(stack.items)-1]

	return lastItem
}

//IsEmpty check whether the stack is empty or not
func (stack *Stack) IsEmpty() bool {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	return len(stack.items) == 0
}

//The Reset method assigns a zero value (nil) to the slice.
func (stack *Stack) Reset() {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	stack.items = nil
}

//Dump method makes and returns a copy of the original slice. It is recommended to return a copy to avoid modifying the original slice.
func (stack *Stack) Dump() []Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()
	var copiedStack = make([]Item, len(stack.items))
	copy(copiedStack, stack.items)

	return copiedStack
}

//Peek() method is similar to the Pop() method except for one thing — Peek() does not remove an item.
func (stack *Stack) Peek() Item {
	stack.mutex.Lock()
	defer stack.mutex.Unlock()

	if len(stack.items) == 0 {
		return nil
	}
	return stack.items[len(stack.items)-1]
}

func main() {
	var stack Stack
	fmt.Println("Stack:", stack.Dump())

	stack.Push(1)
	stack.Push(6)
	stack.Push("one")
	stack.Push(10)
	stack.Push(10.02)
	stack.Push("two")

	fmt.Println("Stack:", stack.Dump())
	fmt.Println("Last Item:", stack.Peek())
	fmt.Println("Stack Is Empty:", stack.IsEmpty())
	fmt.Println("Last removed item:", stack.Pop())
	fmt.Println("Stack :", stack.Dump())

	stack.Reset()

	fmt.Println("Stack is empty:", stack.IsEmpty())
	fmt.Println("Last Item:", stack.Peek())
}
