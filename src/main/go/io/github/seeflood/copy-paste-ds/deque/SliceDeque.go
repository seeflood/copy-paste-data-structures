package deque

type Deque interface {
	AddFirst(e interface{}) bool

	AddLast(e interface{}) bool

	PollFirst() interface{}

	PollLast() interface{}

	PeekFirst() interface{}

	PeekLast() interface{}

	Contains(e interface{}) bool

	Size() int

	IsEmpty() bool
}

type SliceDeque struct {
	slice []interface{}
}

func (dq *SliceDeque) AddFirst(e interface{}) bool {
	if e == nil {
		return false
	}
	dq.slice = append([]interface{}{e}, dq.slice...)
	return true
}

func (dq *SliceDeque) AddLast(e interface{}) bool {
	if e == nil {
		return false
	}
	dq.slice = append(dq.slice, e)
	return true
}

func (dq *SliceDeque) PollFirst() interface{} {
	if dq.IsEmpty() {
		return nil
	}
	result := dq.slice[0]
	dq.slice = dq.slice[1:]
	return result
}

func (dq *SliceDeque) PollLast() interface{} {
	l := dq.Size()
	if l == 0 {
		return nil
	}
	result := dq.slice[l-1]
	dq.slice = dq.slice[:l]
	return result
}

func (dq *SliceDeque) PeekFirst() interface{} {
	if dq.IsEmpty() {
		return nil
	}
	return dq.slice[0]
}

func (dq *SliceDeque) PeekLast() interface{} {
	l := dq.Size()
	if l == 0 {
		return nil
	}
	return dq.slice[l-1]
}

func (dq *SliceDeque) Contains(o interface{}) bool {
	for _, e := range dq.slice {
		if e == o {
			return true
		}
	}
	return false
}

func (dq *SliceDeque) Size() int {
	return len(dq.slice)
}

func (dq *SliceDeque) IsEmpty() bool {
	return dq.Size() == 0
}

func NewSliceDeques() Deque {
	return &SliceDeque{}
}
