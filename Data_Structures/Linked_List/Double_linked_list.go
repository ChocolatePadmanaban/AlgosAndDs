package Double_linked_list

import (
	"fmt"
	"strings"
)

type DoublyLinkedListNode struct {
	item int
	prev *DoublyLinkedListNode
	next *DoublyLinkedListNode
}

func (node *DoublyLinkedListNode) laterNode(i int) *DoublyLinkedListNode {
	if i == 0 {
		return node
	}
	if node.next == nil {
		panic("Index out of range")
	}
	return node.next.laterNode(i - 1)
}

type DoublyLinkedListSeq struct {
	head *DoublyLinkedListNode
	tail *DoublyLinkedListNode
}

func (list *DoublyLinkedListSeq) String() string {
	var result []string
	for node := list.head; node != nil; node = node.next {
		result = append(result, fmt.Sprintf("(%d)", node.item))
	}
	return strings.Join(result, "-")
}

func (list *DoublyLinkedListSeq) Build(items []int) {
	for _, item := range items {
		list.InsertLast(item)
	}
}

func (list *DoublyLinkedListSeq) GetAt(i int) int {
	node := list.head.laterNode(i)
	return node.item
}

func (list *DoublyLinkedListSeq) SetAt(i int, x int) {
	node := list.head.laterNode(i)
	node.item = x
}

func (list *DoublyLinkedListSeq) InsertFirst(x int) {
	newNode := &DoublyLinkedListNode{item: x}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

func (list *DoublyLinkedListSeq) InsertLast(x int) {
	newNode := &DoublyLinkedListNode{item: x}
	if list.tail == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *DoublyLinkedListSeq) DeleteFirst() int {
	if list.head == nil {
		panic("List is empty")
	}
	x := list.head.item
	list.head = list.head.next
	if list.head == nil {
		list.tail = nil
	} else {
		list.head.prev = nil
	}
	return x
}

func (list *DoublyLinkedListSeq) DeleteLast() int {
	if list.tail == nil {
		panic("List is empty")
	}
	x := list.tail.item
	list.tail = list.tail.prev
	if list.tail == nil {
		list.head = nil
	} else {
		list.tail.next = nil
	}
	return x
}

func (list *DoublyLinkedListSeq) Remove(x1, x2 *DoublyLinkedListNode) *DoublyLinkedListSeq {
	L2 := &DoublyLinkedListSeq{head: x1, tail: x2}
	if x1 == list.head {
		list.head = x2.next
	} else {
		x1.prev.next = x2.next
	}
	if x2 == list.tail {
		list.tail = x1.prev
	} else {
		x2.next.prev = x1.prev
	}
	x1.prev = nil
	x2.next = nil
	return L2
}

func (list *DoublyLinkedListSeq) Splice(x *DoublyLinkedListNode, L2 *DoublyLinkedListSeq) {
	xn := x.next
	x1 := L2.head
	x2 := L2.tail
	L2.head = nil
	L2.tail = nil
	x1.prev = x
	x.next = x1
	x2.next = xn
	if xn != nil {
		xn.prev = x2
	} else {
		list.tail = x2
	}
}
