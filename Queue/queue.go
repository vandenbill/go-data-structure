package main

import "fmt"

/*
	Queue is a data stucture dan has FIFO term, mean first in first out
	used in like rabbit mq and redis pub/sub etc
*/

type Queue struct {
	values []int
}

// Enqueue mean add a data to queue
func (q *Queue) enqueue(value int) {
	q.values = append(q.values, value)
}

// Dequeue mean remove a data from queue
func (q *Queue) dequeue() int {
	value := q.values[0]
	q.values = q.values[1:]

	return value
}

func main() {
	queue := Queue{}
	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)
	queue.enqueue(4)
	queue.enqueue(5)
	fmt.Println(queue)

	queue.dequeue()
	queue.dequeue()
	queue.dequeue()
	fmt.Println(queue)
}
