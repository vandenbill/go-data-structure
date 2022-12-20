package main

import "fmt"

/*
	Queue is a data stucture dan has FIFO term, mean first in first out
	used in like rabbit mq and redis pub/sub etc
	time complexity of queue
	- represented as a slice
	enqueu = o(n) bcs queu has to copy all of the data to new array
	dequeu = o(n) bcs queu has to copy all of the data to new array after taken the
			 data
	- represented as a linkedlist
	enqueue = o(1) queue add on head
	dequeue = o(n) queue has to look from head to the tail node
*/

// queue represented as a linkedlist
type QueueNode struct {
	value int
	next  *QueueNode
}

type QueueLinkedlist struct {
	head   *QueueNode
	length int
}

/*
Enqueue in linkedlist has more better time complexity
compare to Queue in slice
*/
func (l *QueueLinkedlist) enqueue(node *QueueNode) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length += 1
}

/*
Dequeue in linkedlist has same time complexity
compare to Queue in slice
*/
func (l *QueueLinkedlist) dequeue() int {
	result := 0
	if l.length == 0 {
		return 0
	}

	if l.head.next == nil {
		result = l.head.value
		l.head = nil
		l.length--
		return result
	}

	length := l.length
	head := l.head
	for length != 0 {
		if length == 2 {
			result = head.next.value
			head.next = nil
			break
		}
		head = head.next
		length--
	}
	l.length--
	return result
}

func (l *QueueLinkedlist) printAllData() {
	length := l.length
	printedData := l.head
	for length > 0 {
		fmt.Printf("%d ", printedData.value)
		printedData = printedData.next
		length--
	}
	fmt.Printf("\n")
}

// queue represented as a slice
type QueueSlice struct {
	values []int
}

// Enqueue mean add a data to queue
func (q *QueueSlice) enqueue(value int) {
	q.values = append(q.values, value)
}

// Dequeue mean remove a data from queue
func (q *QueueSlice) dequeue() int {
	value := q.values[0]
	q.values = q.values[1:]

	return value
}

func main() {
	queueSlice := QueueSlice{}
	queueSlice.enqueue(1)
	queueSlice.enqueue(2)
	queueSlice.enqueue(3)
	queueSlice.enqueue(4)
	queueSlice.enqueue(5)
	fmt.Println(queueSlice)

	queueSlice.dequeue()
	queueSlice.dequeue()
	queueSlice.dequeue()
	fmt.Println(queueSlice)

	queueLinkedlist := QueueLinkedlist{}
	node1 := QueueNode{value: 1}
	node2 := QueueNode{value: 2}
	node3 := QueueNode{value: 3}
	node4 := QueueNode{value: 4}
	queueLinkedlist.enqueue(&node1)
	queueLinkedlist.enqueue(&node2)
	queueLinkedlist.enqueue(&node3)
	queueLinkedlist.enqueue(&node4)
	queueLinkedlist.printAllData()

	queueLinkedlist.dequeue()
	queueLinkedlist.dequeue()
	queueLinkedlist.printAllData()
}
