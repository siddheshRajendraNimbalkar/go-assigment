package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkList struct {
	head *Node
}

func (ll *LinkList) Insert(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkList) display() {
	if ll.head == nil {
		fmt.Println("Link List is Empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Println("data:", current.data)
		fmt.Println("Address:", current)
		current = current.next
	}
}

func main() {
	ll := &LinkList{}

	ll.Insert(12)
	ll.Insert(13)
	ll.Insert(14)
	ll.Insert(15)

	ll.display()
}
