package main

import "fmt"

/*
	in the hash.go, we cover the hash data structure,
	that data structure has big problem which is collision,
	in this file we will cover the solution by replace the actual
	data of the array with the linked list.
	so whenever the data is collision we store that data into
	next node in the linked list
*/

// hash table struct
// bucket struct
// node struct

const l = 10

type HashTable struct {
	arr [l]Bucket
}

func NewHasTable() *HashTable {
	hashTable := &HashTable{}
	for i := 0; i < l; i++ {
		bucket := Bucket{}
		hashTable.arr[i] = bucket
	}
	return hashTable
}

// insert
func (h *HashTable) insert(k string) int {
	index := Hash2(k)
	if h.arr[index].search(k) {
		return 0
	}
	h.arr[index].insert(k)
	return index
}

// delete
func (h *HashTable) delete(k string) bool {
	index := Hash2(k)
	if !h.arr[index].search(k) {
		return false
	}

	head := h.arr[index].head
	if head.key == k {
		h.arr[index].head = h.arr[index].head.next
		return true
	}
	for head.next != nil {
		if head.next.key == k {
			head.next = head.next.next
			break
		}
		head = head.next
	}
	return true
}

// list
func (h *HashTable) list(k string) {
	i := Hash2(k)
	head := h.arr[i].head
	for head != nil {
		fmt.Println(head.key)
		head = head.next
	}
}

// search
func (h *HashTable) search(k string) bool {
	i := Hash2(k)
	return h.arr[i].search(k)
}

type Bucket struct {
	head *Node
}

// insert
func (b *Bucket) insert(k string) {
	if b.search(k) {
		fmt.Println("value already exist")
	}
	node := NewNode(k)
	scnd := b.head
	b.head = node
	b.head.next = scnd
}

// search
func (b *Bucket) search(k string) bool {
	head := b.head
	for head != nil {
		if head.key == k {
			return true
		}
		head = head.next
	}
	return false
}

type Node struct {
	key  string
	next *Node
}

func NewNode(k string) *Node {
	return &Node{key: k}
}

func Hash2(k string) int {
	r := 0
	for _, v := range k {
		r += int(v)
	}
	return r % l
}

func main() {
	hashTable := NewHasTable()
	hashTable.insert("nabil")
	hashTable.insert("libna")
	hashTable.insert("liban")
	hashTable.insert("eric")
	hashTable.insert("ericdssd")
	fmt.Println(Hash2("ericdssd"))
	fmt.Println(Hash2("eric"))
	hashTable.insert("a")

	hashTable.list("nabil")
	fmt.Println(hashTable)

	fmt.Println("============")
	hashTable.delete("liban")
	fmt.Println(hashTable.search("liban"))
	hashTable.list("liban")
}
