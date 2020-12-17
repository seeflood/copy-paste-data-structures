package segment_tree

import (
	"testing"
)

func TestCase1(t *testing.T) {
	nums := [...]int{1, 3, 5}
	obj := Constructor(nums[:])
	r := obj.SumRange(0, 2)
	if r != 9 {
		t.Error()
	}
	obj.Update(1, 2)
	r = obj.SumRange(0, 2)
	if r != 8 {
		t.Error()
	}
}

type NumArray struct {
	tree Tree
}

func Constructor(nums []int) NumArray {
	n := len(nums)
	tree := CreateSegmentTree(n, new(Adder))
	for i := 0; i < n; i++ {
		tree.Set(i, i, nums[i])
	}
	return NumArray{
		tree,
	}
}

func (this *NumArray) Update(i int, val int) {
	this.tree.Set(i, i, val)
}

func (this *NumArray) SumRange(i int, j int) int {
	query, e := this.tree.Query(i, j)
	if e != nil {
		return 0
	}
	return query.(int)
}
