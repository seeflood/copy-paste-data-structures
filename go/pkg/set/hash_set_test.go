package set

import (
	"testing"
)

func TestSet(t *testing.T) {
	set := NewHashSet()

	AssertTrue(set.IsEmpty(), t)
	set.Add("a")
	AssertTrue(set.Size() == 1, t)
	AssertTrue(!set.IsEmpty(), t)
	set.Add(1)

	AssertTrue(set.Contains("a"), t)
	AssertTrue(set.Contains(1), t)
	AssertTrue(set.Size() == 2, t)

	set.Remove("a")
	AssertTrue(!set.Contains("a"), t)
	AssertTrue(set.Size() == 1, t)

	set.Clear()
	AssertTrue(!set.Contains("a"), t)
	AssertTrue(!set.Contains(1), t)
	AssertTrue(set.Size() == 0, t)
	AssertTrue(set.IsEmpty(), t)

	set.Add("a")
	AssertTrue(set.Size() == 1, t)
	AssertTrue(!set.IsEmpty(), t)
	set.Add(1)

	AssertTrue(set.Contains("a"), t)
	AssertTrue(set.Contains(1), t)
	AssertTrue(set.Size() == 2, t)

	set.Remove("a")
	AssertTrue(!set.Contains("a"), t)
	AssertTrue(set.Size() == 1, t)

}

func TestForEach(t *testing.T) {
	set := NewHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	b1 := false
	b2 := false
	b3 := false
	set.ForEach(func(e interface{}) {
		var flag *bool
		if e.(int) == 1 {
			flag = &b1
		} else if e.(int) == 2 {
			flag = &b2
		} else if e.(int) == 3 {
			flag = &b3
		} else {
			t.Error()
		}
		AssertFalse(*flag, t)
		*flag = true
	})
	AssertTrue(b1 && b2 && b3, t)
	b1 = false
	b2 = false
	b3 = false
	set.Remove(2)
	set.ForEach(func(e interface{}) {
		var flag *bool
		if e.(int) == 1 {
			flag = &b1
		} else if e.(int) == 2 {
			flag = &b2
		} else if e.(int) == 3 {
			flag = &b3
		} else {
			t.Error()
		}
		AssertFalse(*flag, t)
		*flag = true
	})
	AssertTrue(b1 && !b2 && b3, t)
}

func AssertFalse(b bool, t *testing.T) {
	AssertTrue(!b, t)
}

func TestFilter(t *testing.T) {
	set := NewHashSet()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set = set.Filter(func(e interface{}) bool {
		return e.(int) <= 2
	})
	AssertTrue(set.Size() == 2, t)
	AssertTrue(!set.IsEmpty(), t)
	AssertTrue(set.Contains(1), t)
	AssertTrue(set.Contains(2), t)
}
func AssertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
