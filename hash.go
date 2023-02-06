package main

import "fmt"

/*
	hash data structur is just array but when we
	insert the data, we mush hash the data to get the
	index of that data and store the data in that index
	bcs of that the insert, search, and delete in
	this data structure is so fast.
	time complexity is o(1)
	but hash has 1 main problem, that is the collision
	of the data, when we hash the data, there are
	possibility of that data os gonna be collision
	and to solve that problem we have to store the
	data that have same index to linked list, we'll
	cover that in the next data structure (hash table)
*/

const length = 10

type HashDS struct {
	arr [length]string
}

func (h *HashDS) Index(k string) int {
	return Hash(k)
}

func (h *HashDS) Insert(k string) {
	h.arr[Hash(k)] = k
}

func (h *HashDS) Delete(k string) (int, bool) {
	i := Hash(k)
	v := h.arr[i]
	if v == "" {
		return i, false
	}

	return i, true
}

func Hash(k string) int {
	r := 0
	for _, v := range k {
		r += int(v)
	}
	return r % length
}

func main1() {
	hash := HashDS{}
	hash.Insert("nabil")
	hash.Insert("nabil")
	fmt.Println(hash.Index("nabil"))
	fmt.Println(hash.arr)
}
