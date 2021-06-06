package map_test

import "testing"

//use LRUCache to solve leetcode-146.
//https://leetcode-cn.com/problems/lru-cache/

type LRU_Cache struct {
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

func NewLRU_Cache(cap int) *LRU_Cache {
	key2node := map[interface{}]*node{}
	return &LRU_Cache{
		cap,
		key2node,
		newDoublyLinkedList(),
	}
}

func (this *LRU_Cache) Get(key interface{}) interface{} {
	nd, ok := this.key2node[key]
	if !ok {
		return nil
	}
	this.refresh(nd)
	return nd.v
}

func (this *LRU_Cache) Put(key interface{}, value interface{}) interface{} {
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

func (this *LRU_Cache) refresh(nd *node) {
	this.list.remove(nd)
	this.list.addFirst(nd)
}

type LRUCache struct {
	c *LRU_Cache
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		NewLRU_Cache(capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	r := this.c.Get(key)
	if r == nil {
		return -1
	}
	return r.(int)
}

func (this *LRUCache) Put(key int, value int) {
	this.c.Put(key, value)
}

func TestLRUCache(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)

	assertTrue(obj.Get(1) == 1, t)
	obj.Put(3, 3)
	assertTrue(obj.Get(2) == -1, t)
	obj.Put(4, 4)
	assertTrue(obj.Get(1) == -1, t)
	assertTrue(obj.Get(3) == 3, t)
	assertTrue(obj.Get(4) == 4, t)
}

func assertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
