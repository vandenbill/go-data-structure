package main

import "fmt"

/*
	heap is data structure that like a tree but actually
	not bcs heap is structured in an array and has a
	advantages in search the max or min value depend on
	MaxHeap or MinHeap
*/

// heap represented as an array / slice (dynamic array)
type Heap struct {
	arr []int
}

/*
	inserting data to heap need 2 step, first we need to
	append the dta to the last index and 2 is heapify up
	to reorganize the tree
*/
func (h *Heap) Insert(v int) {
	h.arr = append(h.arr, v)
	h.HeapifyUp(len(h.arr) - 1)
}

func (h *Heap) Extract() int {
	if len(h.arr) == 0 {
		fmt.Println("cant extract bcs lenght is 0")
		return -1
	}
	r := h.arr[0]
	l := len(h.arr) - 1
	h.arr[0] = h.arr[l]
	h.arr = h.arr[:l]
	h.HeapifyDown(0)
	return r
}

func (h *Heap) HeapifyUp(i int) {
	for h.arr[h.parent(i)] < h.arr[i] {
		h.swap(h.parent(i), i)
		i = h.parent(i)
	}
}

func (h *Heap) HeapifyDown(i int) {
	lc, rc := h.leftChild(i), h.rightChild(i)
	compr := 0
	lastc := len(h.arr) - 1

	for lc <= lastc {
		if lc == lastc {
			compr = lc
		} else if h.arr[lc] > h.arr[rc] {
			compr = lc
		} else {
			compr = rc
		}
		if h.arr[compr] > h.arr[i] {
			h.swap(i, compr)
			i = compr
			lc, rc = h.leftChild(i), h.rightChild(i)
		} else {
			break
		}
	}
}

func (h *Heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *Heap) leftChild(i int) int {
	return 2*i + 1
}

func (h *Heap) rightChild(i int) int {
	return 2*i + 2
}

func (h *Heap) swap(i1, i2 int) {
	h.arr[i1], h.arr[i2] = h.arr[i2], h.arr[i1]
}

func main() {
	heap := &Heap{}
	nmbrs := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for _, v := range nmbrs {
		heap.Insert(v)
	}
	fmt.Println(heap)
	for i := 0; i < 5; i++ {
		heap.Extract()
		fmt.Println(heap)
	}

}
