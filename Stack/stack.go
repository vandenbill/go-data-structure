package main

import (
	"fmt"
	"math"
)

// TODO represented stack as a linkedlist

// Stack data structure / LIFO (last in first out)
type Stack struct {
	values []int
}

// Push add data to end of stack
func (s *Stack) Push(value int) {
	s.values = append(s.values, value)
}

// Pop take data from end of stack / LIFO
func (s *Stack) Pop() int {
	lAP := len(s.values) - 1
	lV := s.values[lAP]
	s.values = s.values[:lAP]

	return lV
}

func main() {
	math.Sqrt(float64(10))
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	fmt.Println(stack)

	stack.Pop()
	stack.Pop()
	fmt.Println(stack)
}
