package main

import "fmt"

type Node struct {
	name string
	rank float64
	next *Node
}

type LinkedList struct {
	size int
	tail *Node
	head *Node
}

func (list *LinkedList) addNode(n *Node) {
	if list.size == 0 {
		list.tail = n
		list.head = n
		list.size++
		return
	}
	list.tail.next = n
	list.tail = n
	list.size++
}

func New(name string, rank float64) *Node {
	return &Node{name, rank, nil}
}

func main() {
	l := LinkedList{}
	l.addNode(New("Google", 0.5))
	l.addNode(New("Facebook", 0.3))
	l.addNode(New("Amazon", 0.2))

	fmt.Println(l)
}
