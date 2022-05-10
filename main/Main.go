package main

import "fmt"

func main() {
	list := LinkedList{}
	list.addNode()
	list.addNode()
	list.showList()
}

type Element struct {
	label string
	rank float32
}

type Node struct {
	element Element
	nextNode *Node
}

type LinkedList struct {
	length int
	head *Node
	tail *Node
}

func (list LinkedList) addNode(){
	chars := []rune("abcdefghijklmnopqrtuvwxyz")
	var charsIndex int
	if list.length == 0 {
		charsIndex = 0
		e := Element{string(chars[charsIndex]), 0}
		n := Node{e, nil}
		charsIndex++
		list.length = charsIndex
		list = LinkedList{charsIndex, &n, &n}
	}else{
		n := Node{Element{string(chars[charsIndex]), 0}, nil}
		charsIndex++
		list.length = charsIndex
		list.tail = &n
	}
}

func (list LinkedList) showList(){
	fmt.Println(list)
}


