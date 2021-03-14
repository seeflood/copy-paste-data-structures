package slice_heap

import (
	"container/heap"
)

type Comparator interface {
	compare(a interface{}, b interface{}) int
}
type SliceHeap struct {
	slice      []interface{}
	comparator Comparator
}

func NewSliceHeap(c Comparator) *SliceHeap {
	h := &SliceHeap{
		make([]interface{}, 0),
		c,
	}
	heap.Init(h)
	return h
}

func NewIntSliceHeap() *SliceHeap {
	h := &SliceHeap{
		make([]interface{}, 0),
		&IntComparator{},
	}
	heap.Init(h)
	return h
}

type IntComparator struct {
}

func (*IntComparator) compare(a interface{}, b interface{}) int {
	return a.(int) - b.(int)
}

func (h *SliceHeap) IsEmpty() bool {
	return h.Len() == 0
}
func (h *SliceHeap) Len() int           { return len(h.slice) }
func (h *SliceHeap) Less(i, j int) bool { return h.comparator.compare(h.slice[i], h.slice[j]) < 0 }
func (h *SliceHeap) Swap(i, j int)      { h.slice[i], h.slice[j] = h.slice[j], h.slice[i] }

func (h *SliceHeap) Push(x interface{}) {
	h.slice = append(h.slice, x)
}

func (h *SliceHeap) Pop() interface{} {
	old := h.slice
	n := len(old)
	x := old[n-1]
	h.slice = old[0 : n-1]
	return x
}

func Peek(h *SliceHeap) interface{} {
	if h.IsEmpty() {
		return nil
	}
	return h.slice[0]
}
