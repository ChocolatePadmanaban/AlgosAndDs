---
layout: default
---
### Before We Run this code do 

```
git clone https://github.com/ChocolatePadmanaban/AlgosAndDs.git
export ALGOS_DS_REPO=`pwd`/AlgosAndDs
cd $ALGOS_DS_REPO/Data_Structures/Linked_List/
```


# Doubly Linked List

**Doubly linked list**, supporting additional constant-time operations. Each node x of a doubly linked
list maintains an x.prev pointer to the node preceeding it in the sequence, in addition to an
x.next pointer to the node following it in the sequence. A doubly linked list L maintains a pointer
to L.tail, the last node in the sequence, in addition to L.head, the ﬁrst node in the sequence.

```
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
```

## O(1) Time Operations

the following algorithms are implemented the following sequence operations, each in O(1) time.

### insert_first(x)

```
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
```

### insert_last(x)

```
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
```

### delete_first(x)

```
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
```

### delete_last(x)
```
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
```


### Remove nodes from x1 to x2 

Given two nodes x1 and x2 from a doubly linked list L, where x1 occurs
before x2, design a constant-time algorithm to remove all nodes from x1 to x2 inclusive from L, and return them as a new doubly linked list.

```
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
```

### Splice list L2 into list L1 after node x 

Given node x from a doubly linked list L1 and second doubly linked list L2,
describe a constant-time algorithm to splice list L2 into list L1 after node x. After the splice operation, L1 should contain all items previously in either list, and L2 should be empty.

```
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
```


## Unit Test to Check Doubly-Linked List and its Function:


In this Folder You can run Unit tests to check Doubly Linked List 
```
go test -v 
```

Output would be :
```
% go test -v  
=== RUN   TestCases
    Double_linked_list_test.go:251: Test 01 Passed
    Double_linked_list_test.go:251: Test 02 Passed
    Double_linked_list_test.go:251: Test 03 Passed
    Double_linked_list_test.go:251: Test 04 Passed
    Double_linked_list_test.go:251: Test 05 Passed
--- PASS: TestCases (0.00s)
PASS
ok      Doubly_Linked_List      0.779s
```


[back](./)