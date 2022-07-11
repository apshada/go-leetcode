// package main

// type Node struct {
// 	data int
// 	next *Node
// }

// type LinkedList struct {
// 	head   *Node
// 	tail   *Node
// 	length int
// }

// func main() {
// 	l := LinkedList{}
// 	n1 := Node{data: 4}
// 	n2 := Node{data: 6}
// 	l.PushBack(&n1)
// 	l.PushBack(&n2)
// }

// func (l *LinkedList) PushBack(node *Node) {
// 	if l.head == nil {
// 		l.head = node
// 		l.tail = node
// 		l.length++
// 	} else {
// 		l.tail.next = node
// 		l.tail = node
// 		l.length++
// 	}
// }

package main

import (
	"container/list"
	"fmt"
)

func main() {
	head := list.New()
	head.PushBack(1)
	head.PushFront(2)
	head.PushFront(3)
	c := 0
	for element := head.Front(); element != nil; element = element.Next() {
		// do something with element.Value
		c++
		// fmt.Println(element.Value)
	}
	if c%2 != 0 {
		c++
	}
	c = c / 2
	z := 0
	for el := head.Front(); el != nil; el = el.Next() {
		if z == c {
			fmt.Println(el)
			// return head
		} else {
			z++
		}
	}
	// fmt.Println(c / 2)
	// we now have a linked list with '1' at the back of the list
	// and '2' at the front of the list.
}
