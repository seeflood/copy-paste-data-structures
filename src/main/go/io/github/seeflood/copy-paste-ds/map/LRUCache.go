package _map

type LRUCache struct {
	cap      int
	key2node map[interface{}]*node
	list     *doublyLinkedList
}

type node struct {
	key  interface{}
	v    interface{}
	prev *node
	next *node
}

func newNode(key interface{}, v interface{}, prev *node, next *node) *node {
	return &node{key: key, v: v, prev: prev, next: next}
}

type doublyLinkedList struct {
	dummy *node
	size  int
}

func newDoublyLinkedList() *doublyLinkedList {
	dm := newNode(nil, nil, nil, nil)
	dm.next = dm
	dm.prev = dm
	return &doublyLinkedList{
		dm,
		0,
	}
}

func (list *doublyLinkedList) addFirst(nd *node) {
	oldNext := list.dummy.next
	list.dummy.next = nd
	oldNext.prev = nd
	nd.prev = list.dummy
	nd.next = oldNext
	list.size++
}

func (list *doublyLinkedList) remove(nd *node) {
	prev := nd.prev
	next := nd.next
	prev.next = next
	next.prev = prev
	nd.prev = nil
	nd.next = nil
	list.size--
}

func (list *doublyLinkedList) removeLast() *node {
	toRemove := list.dummy.prev
	list.remove(toRemove)
	return toRemove
}

func NewLRUCache(cap int) *LRUCache {
	key2node := map[interface{}]*node{}
	return &LRUCache{
		cap,
		key2node,
		newDoublyLinkedList(),
	}
}

func (this *LRUCache) Get(key interface{}) interface{} {
	nd, ok := this.key2node[key]
	if !ok {
		return nil
	}
	this.refresh(nd)
	return nd.v
}

func (this *LRUCache) Put(key interface{}, value interface{}) interface{} {
	if this.cap <= 0 {
		return value
	}
	nd, ok := this.key2node[key]
	if ok {
		nd.v = value
		this.refresh(nd)
		return nil
	}
	nd = newNode(key, value, nil, nil)
	this.key2node[key] = nd
	this.list.addFirst(nd)
	if this.list.size > this.cap {
		removed := this.list.removeLast()
		delete(this.key2node, removed.key)
		return removed.v
	} else {
		return nil
	}
}

func (this *LRUCache) refresh(nd *node) {
	this.list.remove(nd)
	this.list.addFirst(nd)
}
