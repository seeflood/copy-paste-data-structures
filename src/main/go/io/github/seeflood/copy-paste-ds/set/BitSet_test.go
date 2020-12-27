package set

import "testing"

func TestBitSet(t *testing.T) {
	set := NewBitSet()

	assertTrue(set.IsEmpty(), t)
	set.Add(5)
	assertTrue(set.Cardinality() == 1, t)
	//assertTrue(!set.IsEmpty(), t)
	set.Add(1)

	assertTrue(set.Contains(5), t)
	assertTrue(set.Contains(1), t)
	//assertTrue(set.Size() == 2, t)
	assertTrue(set.Cardinality() == 2, t)

	set.Remove(5)
	assertTrue(!set.Contains(5), t)
	assertTrue(set.Cardinality() == 1, t)

	set.Clear()
	assertTrue(!set.Contains(5), t)
	assertTrue(!set.Contains(1), t)
	assertTrue(set.Cardinality() == 0, t)
	//assertTrue(set.IsEmpty(), t)

	set.Add(5)
	assertTrue(set.Size() == 1, t)
	assertTrue(!set.IsEmpty(), t)
	set.Add(1)

	assertTrue(set.Contains(5), t)
	assertTrue(set.Contains(1), t)
	assertTrue(set.Cardinality() == 2, t)

	set.Remove(5)
	assertTrue(!set.Contains(5), t)
	assertTrue(set.Cardinality() == 1, t)

}

func assertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
