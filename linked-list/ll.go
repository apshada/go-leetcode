package main

import (
	"fmt"
)

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
	ll := LinkedList{}
	n1 := &Node{data: 1}
	n2 := &Node{data: 2}
	n3 := &Node{data: 3}
	n4 := &Node{data: 4}
	ll.PushBack(n1)
	ll.PushBack(n2)
	ll.PushBack(n3)
	ll.PushBack(n4)
	// ll.Display()
	// fmt.Println(ll.length)
	res := middleNode(n1)
	fmt.Println(res.data)
}

func middleNode(head *Node) *Node {
	initial := head
	count := 1
	for ; head != nil; head = head.nextNode {
		count++
	}
	count = count /2
	for ; count > 0; initial = initial.nextNode {
		count--
	}
	return initial
}

func (ll *LinkedList) PushBack(n *Node) {
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		ll.length++
	} else {
		ll.tail.nextNode = n
		ll.tail = n
		ll.length++
	}
}

func (ll LinkedList) Display() {
	for ll.head != nil {
		fmt.Printf("%v -> ", ll.head.data)
		ll.head = ll.head.nextNode
	}
	fmt.Println()
}
