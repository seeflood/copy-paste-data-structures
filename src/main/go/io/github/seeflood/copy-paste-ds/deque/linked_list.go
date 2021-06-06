package deque

type LinkedList struct {
	dummy *node
	size  int
}

type node struct {
	v    interface{}
	prev *node
	next *node
}

func newNode(v interface{}, prev *node, next *node) *node {
	return &node{v: v, prev: prev, next: next}
}

func NewLinkedList() *LinkedList {
	dummy := newNode(nil, nil, nil)
	dummy.prev = dummy
	dummy.next = dummy
	return &LinkedList{
		dummy,
		0,
	}
}

func (list *LinkedList) AddFirst(e interface{}) bool {
	if e == nil {
		return false
	}
	oldHead := list.dummy.next
	node := newNode(e, list.dummy, oldHead)
	list.dummy.next = node
	oldHead.prev = node
	list.size++
	return true
}

func (list *LinkedList) AddLast(e interface{}) bool {
	if e == nil {
		return false
	}
	oldTail := list.dummy.prev
	node := newNode(e, oldTail, list.dummy)
	list.dummy.prev = node
	oldTail.next = node
	list.size++
	return true
}

func (list *LinkedList) PollFirst() interface{} {
	if list.size == 0 {
		return nil
	}
	removed := list.removeNode(list.dummy.next)
	return removed.v
}

func (list *LinkedList) PollLast() interface{} {
	if list.size == 0 {
		return nil
	}
	removed := list.removeNode(list.dummy.prev)
	return removed.v
}

func (list *LinkedList) PeekFirst() interface{} {
	if list.size == 0 {
		return nil
	}
	return list.dummy.next.v
}

func (list *LinkedList) PeekLast() interface{} {
	if list.size == 0 {
		return nil
	}
	return list.dummy.prev.v
}

func (list *LinkedList) Contains(e interface{}) bool {
	prev := list.dummy
	for i := 0; i < list.size; i++ {
		prev = prev.next
		if prev.v == e {
			return true
		}
	}

	return false
}

func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.Size() == 0
}

func (list *LinkedList) removeNode(toRemove *node) *node {
	prev := toRemove.prev
	next := toRemove.next
	prev.next = next
	next.prev = prev
	list.size--
	toRemove.next = nil
	toRemove.prev = nil
	return toRemove
}
