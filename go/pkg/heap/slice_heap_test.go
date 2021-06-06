package slice_heap

import (
	"container/heap"
	"testing"
)

func TestSliceHeap(t *testing.T) {
	h := NewSliceHeap(&IntComparator{})
	AssertTrue(h.IsEmpty(), t)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 1)
	heap.Push(h, 8)
	heap.Push(h, 2)

	AssertTrue(Peek(h) == 1, t)
	AssertTrue(heap.Pop(h) == 1, t)
	AssertTrue(Peek(h) == 2, t)
	AssertTrue(heap.Pop(h) == 2, t)
	AssertTrue(Peek(h) == 3, t)
	AssertTrue(heap.Pop(h) == 3, t)
	AssertTrue(Peek(h) == 5, t)
	AssertTrue(heap.Pop(h) == 5, t)
	AssertTrue(!h.IsEmpty(), t)
	AssertTrue(heap.Pop(h) == 8, t)
	AssertTrue(h.Len() == 0, t)
	AssertTrue(h.IsEmpty(), t)
}

func TestIntSliceHeap(t *testing.T) {
	h := NewIntSliceHeap()
	AssertTrue(h.IsEmpty(), t)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 1)
	heap.Push(h, 8)
	heap.Push(h, 2)

	AssertTrue(Peek(h) == 1, t)
	AssertTrue(heap.Pop(h) == 1, t)
	AssertTrue(Peek(h) == 2, t)
	AssertTrue(heap.Pop(h) == 2, t)
	AssertTrue(Peek(h) == 3, t)
	AssertTrue(heap.Pop(h) == 3, t)
	AssertTrue(Peek(h) == 5, t)
	AssertTrue(heap.Pop(h) == 5, t)
	AssertTrue(!h.IsEmpty(), t)
	AssertTrue(heap.Pop(h) == 8, t)
	AssertTrue(h.Len() == 0, t)
	AssertTrue(h.IsEmpty(), t)
}

func AssertTrue(b bool, t *testing.T) {
	t.Helper()
	if !b {
		t.Error()
	}
}
