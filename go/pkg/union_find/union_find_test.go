package union_find

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(10)
	AssertTrue(uf.CountComponents() == 10, t)
	AssertTrue(!uf.Connected(1, 2), t)
	AssertTrue(!uf.Connected(3, 4), t)
	AssertTrue(!uf.Connected(3, 4), t)
	uf.Union(1, 2)
	AssertTrue(uf.Connected(1, 2), t)
	AssertTrue(uf.CountComponents() == 9, t)
	uf.Union(1, 4)
	AssertTrue(uf.Connected(1, 2), t)
	AssertTrue(uf.Connected(1, 4), t)
	AssertTrue(uf.Connected(2, 4), t)
	AssertTrue(uf.CountComponents() == 8, t)
	uf.Union(2, 3)
	AssertTrue(uf.Connected(1, 3), t)
	AssertTrue(uf.Connected(1, 4), t)
	AssertTrue(uf.Connected(2, 4), t)
	AssertTrue(uf.Connected(3, 4), t)
	AssertTrue(uf.CountComponents() == 7, t)

}

func AssertTrue(b bool, t *testing.T) {
	if !b {
		t.Error()
	}
}
