package segment_tree

import "errors"

type Tree interface {
	Query(lo int, hi int) (interface{}, error)
	Set(lo int, hi int, value interface{}) error
	Update(lo int, hi int, value interface{}) error
}
type AggregateFunction interface {
	aggregate(a interface{}, b interface{}) interface{}
	update(old interface{}, from int, to int, diff interface{}) interface{}
	set(from int, to int, newValue interface{}) interface{}
	getEmpty() interface{}
}

type Adder struct {
}

func (adder *Adder) aggregate(a interface{}, b interface{}) interface{} {
	return a.(int) + b.(int)
}

func (adder *Adder) update(old interface{}, from int, to int, diff interface{}) interface{} {
	return old.(int) + diff.(int)*(to-from+1)
}

func (adder *Adder) set(from int, to int, newValue interface{}) interface{} {
	return newValue.(int) * (to - from + 1)
}

func (adder *Adder) getEmpty() interface{} {
	return 0
}

func CreateSegmentTree(size int, aggregator AggregateFunction) Tree {
	// 1. init tree array
	arrSize := 2*nextPowerOf2(size) - 1
	if arrSize < 0 {
		arrSize = 0
	}
	arr := make([]interface{}, arrSize)
	for i := 0; i < arrSize; i++ {
		arr[i] = aggregator.getEmpty()
	}
	lazy := make([]interface{}, arrSize)
	lazyIsUpdate := make([]bool, arrSize)

	return &segmentTreeImpl{aggregator, arr, lazy, lazyIsUpdate, size}
}

func nextPowerOf2(i int) int {
	if i&(i-1) == 0 {
		return i
	}
	for i&(i-1) != 0 {
		i &= (i - 1)
	}
	return i << 1
}

type segmentTreeImpl struct {
	aggregator   AggregateFunction
	tree         []interface{}
	lazy         []interface{}
	lazyIsUpdate []bool
	n            int
}

func (st *segmentTreeImpl) Query(lo int, hi int) (interface{}, error) {
	return st.query(lo, hi, 0, st.n-1, 0)
}

func (st *segmentTreeImpl) query(qlo int, qhi int, lo int, hi int, idx int) (interface{}, error) {
	if qlo > qhi {
		return nil, errors.New("Illegal params.")
	}
	if hi < qlo || qhi < lo {
		return st.aggregator.getEmpty(), nil
	}
	if qlo <= lo && hi <= qhi {
		return st.nil2Empty(st.tree[idx]), nil
	}
	mid := lo + (hi-lo)/2
	left, _ := st.query(qlo, qhi, lo, mid, 2*idx+1)
	right, _ := st.query(qlo, qhi, mid+1, hi, 2*idx+2)
	return st.aggregator.aggregate(left, right), nil
}

func (st *segmentTreeImpl) nil2Empty(v interface{}) interface{} {
	if v == nil {
		return st.aggregator.getEmpty()
	}
	return v
}

func (st *segmentTreeImpl) Set(lo int, hi int, value interface{}) error {
	return st.update(lo, hi, 0, st.n-1, 0, value, false)
}

func (st *segmentTreeImpl) Update(lo int, hi int, value interface{}) error {
	return st.update(lo, hi, 0, st.n-1, 0, value, true)
}

func (st *segmentTreeImpl) update(qlo int, qhi int, lo int, hi int, idx int, value interface{}, isUpdate bool) error {
	if qlo > qhi {
		return errors.New("Illegal params.")
	}
	if lo > hi {
		return nil
	}
	if hi < qlo || qhi < lo {
		return nil
	}
	st.checkLazy(lo, hi, idx)
	//	cover
	if qlo <= lo && hi <= qhi {
		if isUpdate {
			st.lazyIsUpdate[idx] = true
		} else {
			st.lazyIsUpdate[idx] = false
		}
		st.lazy[idx] = value
		st.checkLazy(lo, hi, idx)
		return nil
	}
	//	overlap
	mid := lo + (hi-lo)/2
	leftI := 2*idx + 1
	rightI := 2*idx + 2
	st.update(qlo, qhi, lo, mid, leftI, value, isUpdate)
	st.update(qlo, qhi, mid+1, hi, rightI, value, isUpdate)
	st.tree[idx] = st.aggregator.aggregate(st.tree[leftI], st.tree[rightI])
	return nil
}

func (st *segmentTreeImpl) checkLazy(lo int, hi int, idx int) {
	if st.lazy[idx] == nil {
		return
	}
	// apply
	st.applyLazy(lo, hi, idx)

	// propagate and merge
	if lo < hi {
		leftI := 2*idx + 1
		rightI := 2*idx + 2
		if st.lazyIsUpdate[idx] {
			st.lazy[leftI] = st.aggregator.aggregate(st.lazy[leftI], st.lazy[idx])
			st.lazy[rightI] = st.aggregator.aggregate(st.lazy[rightI], st.lazy[idx])
		} else {
			st.lazy[leftI] = st.lazy[idx]
			st.lazyIsUpdate[leftI] = false
			st.lazy[rightI] = st.lazy[idx]
			st.lazyIsUpdate[rightI] = false
		}
	}
	//	clear
	st.lazy[idx] = nil
}

func (st *segmentTreeImpl) applyLazy(lo int, hi int, idx int) {
	if st.lazy[idx] == nil {
		return
	}
	if st.lazyIsUpdate[idx] {
		st.tree[idx] = st.aggregator.update(st.tree[idx], lo, hi, st.lazy[idx])
	} else {
		st.tree[idx] = st.aggregator.set(lo, hi, st.lazy[idx])
	}
}
