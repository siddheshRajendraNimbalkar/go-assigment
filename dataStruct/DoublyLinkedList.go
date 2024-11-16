package main

import "fmt"

type NodeDLL struct {
	prev *NodeDLL
	data int
	next *NodeDLL
}

type DoubleLinkList struct {
	head *NodeDLL
}

func (dll *DoubleLinkList) Insert(data int) {
	newNode := &NodeDLL{data: data}
	if dll.head == nil {
		dll.head = newNode
		return
	}
	current := dll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.prev = current
}

func (dll *DoubleLinkList) display() {
	if dll.head == nil {
		fmt.Println("Double Link List is empty")
		return
	}

	current := dll.head
	for current != nil {
		fmt.Println(current)
		current = current.next
	}
}

func main() {
	dll := &DoubleLinkList{}

	dll.Insert(12)
	dll.Insert(13)
	dll.Insert(14)
	dll.Insert(15)

	dll.display()
}
