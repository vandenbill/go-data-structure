package main

import "fmt"

type Node struct {
	value string
	next  *Node
}

type Linkedlist struct {
	head   *Node
	length int
}

func (l *Linkedlist) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length += 1
}

func (l *Linkedlist) printAllData() {
	length := l.length
	printedData := l.head
	for length > 0 {
		fmt.Printf("%s ", printedData.value)
		printedData = printedData.next
		length--
	}
	fmt.Printf("\n")
}

func (l *Linkedlist) printSequenceByValue(value string) {
	if l.head.value == value {
		fmt.Println(0)
		return
	}

	printedData := l.head
	length := l.length
	for length != 0 {
		if printedData.value == value {
			fmt.Println(l.length - length)
			return
		}
		printedData = printedData.next
		length--
	}
}

func (l *Linkedlist) deleteByValue(value string) {
	if l.head.value == value {
		l.head = l.head.next
		l.length--
		return
	}

	length := l.length
	beforeDelete := l.head
	for length != 0 {
		if beforeDelete.next == nil {
			break
		}
		if beforeDelete.next.value == value {
			if beforeDelete.next.next == nil {
				beforeDelete.next = nil
				l.length--
				break
			}
			beforeDelete.next = beforeDelete.next.next
			l.length--
			break
		}
		beforeDelete = beforeDelete.next
		length--
	}
}

func main() {
	linkedList := Linkedlist{}
	node1 := Node{value: "pertama"}
	node2 := Node{value: "ke-2"}
	node3 := Node{value: "ke-3"}
	node4 := Node{value: "ke-4"}
	linkedList.prepend(&node1)
	linkedList.prepend(&node2)
	linkedList.prepend(&node3)
	linkedList.prepend(&node4)
	linkedList.printAllData()
	linkedList.printSequenceByValue("pertama")
	fmt.Println(linkedList.length)
	linkedList.deleteByValue("pertama")
	linkedList.deleteByValue("ke-2")
	linkedList.printAllData()
	fmt.Println(linkedList.length)
}
