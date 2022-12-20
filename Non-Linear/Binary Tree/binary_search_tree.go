package main

import "fmt"

/*
	BINARY SEARCH TREE
	binary search tree is data structure that hace advantages
	in speed of search, the time complexity of searching in this
	data structure is o(h), h is height of the tree, and h have
	a corelation with n which is log n, so time complexity
	also can describe as o(log n)
*/

/*
	the node contain 3 part, the value, addres of left and right child
*/

// node
type Node struct {
	Value      int
	LeftChild  *Node
	RightChild *Node
}

/*
	if it smaller than the parent, it should be left child, and if
	it larger it should be right child and so on, until the node
	have no child
*/
// insert data
func (n *Node) Insert(k int) {
	if n.Value > k {
		if n.LeftChild == nil {
			n.LeftChild = &Node{Value: k}
		} else {
			n.Insert(k)
		}
	} else if n.Value < k {
		if n.RightChild == nil {
			n.RightChild = &Node{Value: k}
		} else {
			n.Insert(k)
		}
	}
}

/*
	logic behind search is also kind of insert, we compare the data
	with the parent to decide left or right child the data should be,
	and when the data is same to the node, it return true, but if not
	it return false
*/
// search data
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Value > k {
		n.LeftChild.Search(k)
	} else if n.Value < k {
		n.RightChild.Search(k)
	}
	return true
}

func main() {
	node := &Node{Value: 10}
	node.Insert(10)
	node.Insert(100)
	fmt.Println(node.Search(100))
	fmt.Println(node)
}
