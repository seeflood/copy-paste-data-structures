package set

import "testing"

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
}

func AssertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
