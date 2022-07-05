package main

import "fmt"

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func main() {
	list := LinkedList{}
	node1 := Node{data: 1}
	node2 := Node{data: 2}
	node3 := Node{data: 3}
	node4 := Node{data: 4}
	list.PushBack(&node1)
	list.PushBack(&node2)
	list.PushBack(&node3)
	list.PushBack(&node4)
	list.Display()
}

func (l *LinkedList) PushBack(n *Node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.nextNode = n
		l.tail = n
		l.length++
	}
}

func (l LinkedList) Display() {
	for l.head != nil {
		fmt.Printf("%v -> ", l.head.data)
		l.head = l.head.nextNode
	}
	fmt.Println()
}
