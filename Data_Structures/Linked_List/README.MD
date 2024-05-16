---
layout: default
---

# Doubly Linked List

**Doubly linked list**, supporting additional constant-time operations. Each node x of a doubly linked
list maintains an x.prev pointer to the node preceeding it in the sequence, in addition to an
x.next pointer to the node following it in the sequence. A doubly linked list L maintains a pointer
to L.tail, the last node in the sequence, in addition to L.head, the ﬁrst node in the sequence.

```
class Doubly_Linked_List_Node:
    def __init__(self, x):
        self.item = x
        self.prev = None
        self.next = None

    def later_node(self, i):
        if i == 0: return self
        assert self.next
        return self.next.later_node(i - 1)

class Doubly_Linked_List_Seq:
    def __init__(self):
        self.head = None
        self.tail = None

    def __iter__(self):
        node = self.head
        while node:
            yield node.item
            node = node.next

    def __str__(self):
        return '-'.join([('(%s)' % x) for x in self])

    def build(self, X):
        for a in X:
            self.insert_last(a)

    def get_at(self, i):
        node = self.head.later_node(i)
        return node.item

    def set_at(self, i, x):
        node = self.head.later_node(i)
        node.item = x
```

## O(1) Time Operations

the following algorithms are implemented the following sequence operations, each in O(1) time.

### insert_first(x)

```
def insert_first(self, x):
    new_node = Doubly_Linked_List_Node(x)
    if self.head is None:
        self.head = new_node
        self.tail = new_node
    else:
        new_node.next = self.head
        self.head.prev = new_node
        self.head = new_node
```

### insert_last(x)

```
def insert_last(self, x):
    new_node = Doubly_Linked_List_Node(x)
    if self.tail is None:
        self.head = new_node
        self.tail = new_node
    else:
        new_node.prev = self.tail
        self.tail.next = new_node
        self.tail = new_node
```

### delete_first(x)

```
def delete_first(self):
    assert self.head
    x = self.head.item
    self.head = self.head.next
    if self.head is None: self.tail = None
    else : self.head.prev = None
    return x
```

### delete_last(x)
```
def delete_last(self):
    assert self.tail
    x = self.tail.item
    self.tail = self.tail.prev
    if self.tail is None: self.head = None
    else : self.tail.next = None 
    return x
```


### Remove nodes from x1 to x2 

Given two nodes x1 and x2 from a doubly linked list L, where x1 occurs
before x2, design a constant-time algorithm to remove all nodes from x1 to x2 inclusive from L, and return them as a new doubly linked list.

```
def remove(self, x1, x2):
    L2 = Doubly_Linked_List_Seq()
    L2.head = x1
    L2.tail = x2
    if x1 == self.head: self.head = x2.next 
    else : x1.prev.next =x2.next 
    if x2 == self.tail: self.tail = x1.prev
    else: x2.next.prev = x1.prev
    x1.prev= None
    x2.next = None
    return L2
```

### Splice list L2 into list L1 after node x 

Given node x from a doubly linked list L1 and second doubly linked list L2,
describe a constant-time algorithm to splice list L2 into list L1 after node x. After the splice operation, L1 should contain all items previously in either list, and L2 should be empty.

```
def splice(self, x, L2):
    xn = x.next
    x1 = L2.head
    x2 = L2.tail 
    L2.head = None
    L2.tail = None
    x1.prev = x
    x.next = x1
    x2.next = xn 
    if xn: xn.prev = x2
    else: self.tail = x2 
```


## Unit Test to Check Doubly-Linked List and its Function:


In this Folder You can run Unit tests to check Doubly Linked List 
```
python3 tests.py
```

Output would be :
```
$ python tests.py 
test_01 (__main__.TestCases) ... ok
test_02 (__main__.TestCases) ... ok
test_03 (__main__.TestCases) ... ok
test_04 (__main__.TestCases) ... ok
test_05 (__main__.TestCases) ... ok

----------------------------------------------------------------------
Ran 5 tests in 0.001s

OK
```


[back](./)