package main

import (
	"fmt"
	"sync"
)

type Item interface{}

type Queue struct {
	items []Item
	mutex sync.Mutex
}

//For inserting the new item into the queue we use the append() built-in function, which inserts a new item at the end of the slice.

func (queue *Queue) enqueue(item Item) {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	queue.items = append(queue.items, item)
}

func (queue *Queue) peek() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}
	return queue.items[len(queue.items)-1]
}

func (queue *Queue) dequeue() Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	if len(queue.items) == 0 {
		return nil
	}

	lastItem := queue.items[0]
	queue.items = queue.items[1:]

	return lastItem
}

func (queue *Queue) reset() {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.items = nil
}

func (queue *Queue) isEmpty() bool {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	return len(queue.items) == 0
}

func (queue *Queue) dump() []Item {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	var copiedQueue = make([]Item, len(queue.items))
	copy(copiedQueue, queue.items)
	return copiedQueue
}

func main() {
	var queue Queue

	queue.enqueue(5)
	queue.enqueue(4)
	queue.enqueue(3)
	queue.enqueue(2)
	queue.enqueue(1)

	fmt.Println("Queue:", queue.dump())
	fmt.Println("The last Item:", queue.peek())

	queue.dequeue()

	fmt.Println("Queue:", queue.dump())
	fmt.Println("Queue is Empty:", queue.isEmpty())

	queue.reset()

	fmt.Println("Queue is Empty:", queue.isEmpty())

}
