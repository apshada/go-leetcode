package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func main() {
	n1 := ListNode{Val: 2}
	n2 := ListNode{Val: 3}
	n3 := ListNode{Val: 4}
	n4 := ListNode{Val: 5}
	ll := LinkedList{}
	ll.PushBack(&n1)
	ll.PushBack(&n2)
	ll.PushBack(&n3)
	ll.PushBack(&n4)
	ll.Display()
	ll.ReverseLinkedList()

}

func (ll *LinkedList) ReverseLinkedList() {

}

func (ll *LinkedList) PushBack(n *ListNode) {
	if ll.head == nil {
		ll.head = n
		ll.tail = n
		ll.length++
	} else {
		ll.tail.Next = n
		ll.tail = n
		ll.length++
	}
}

func (ll LinkedList) Display() {
	for ll.head != nil {
		fmt.Printf("%v ->", ll.head.Val)
		ll.head = ll.head.Next
	}
	fmt.Println()
}
