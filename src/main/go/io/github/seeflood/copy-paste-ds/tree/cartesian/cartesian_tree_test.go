package cartesian

import "testing"

func TestTree(t *testing.T) {
	arr := []interface{}{4, 2, 5, 1, 6, 3, 7}
	tree := NewTree(arr, func(a interface{}, b interface{}) int {
		return a.(int) - b.(int)
	})
	root := tree.Root()
	assertTrue(root.Value() == 1 && root.Idx() == 3 && root.from == 0 && root.to == 6, t)
	n1 := root.Left()
	assertTrue(n1.Value() == 2 && n1.Idx() == 1 && n1.from == 0 && n1.to == 2, t)
	n5 := root.Right()
	assertTrue(n5.Value() == 3 && n5.Idx() == 5 && n5.from == 4 && n5.to == 6, t)

	n2 := n1.Right()
	assertTrue(n2.Value() == 5 && n2.Idx() == 2 && n2.from == 2 && n2.to == 2, t)

	n0 := n1.Left()
	assertTrue(n0.value == arr[n0.Idx()] && n0.Idx() == 0 && n0.from == 0 && n0.to == 0, t)

	n4 := n5.Left()
	assertTrue(n4.value == arr[n4.Idx()] && n4.Idx() == 4 && n4.from == 4 && n4.to == 4, t)
	n6 := n5.Right()
	assertTrue(n6.value == arr[n6.Idx()] && n6.Idx() == 6 && n6.from == 6 && n6.to == 6, t)
}

func TestIntTree_minOnTop(t *testing.T) {
	arr := []int{4, 2, 5, 1, 6, 3, 7}
	tree := NewIntTree(arr, true)
	root := tree.Root()
	assertTrue(root.Value() == 1 && root.Idx() == 3 && root.from == 0 && root.to == 6, t)
	n1 := root.Left()
	assertTrue(n1.Value() == 2 && n1.Idx() == 1 && n1.from == 0 && n1.to == 2, t)
	n5 := root.Right()
	assertTrue(n5.Value() == 3 && n5.Idx() == 5 && n5.from == 4 && n5.to == 6, t)

	n2 := n1.Right()
	assertTrue(n2.Value() == 5 && n2.Idx() == 2 && n2.from == 2 && n2.to == 2, t)

	n0 := n1.Left()
	assertTrue(n0.value == arr[n0.Idx()] && n0.Idx() == 0 && n0.from == 0 && n0.to == 0, t)

	n4 := n5.Left()
	assertTrue(n4.value == arr[n4.Idx()] && n4.Idx() == 4 && n4.from == 4 && n4.to == 4, t)
	n6 := n5.Right()
	assertTrue(n6.value == arr[n6.Idx()] && n6.Idx() == 6 && n6.from == 6 && n6.to == 6, t)
}

func TestIntTree_maxOnTop(t *testing.T) {
	arr := []int{40, 200, 5, 1000, 60, 300, 7}
	tree := NewIntTree(arr, false)
	root := tree.Root()
	assertTrue(root.Value() == 1000 && root.Idx() == 3 && root.from == 0 && root.to == 6, t)
	n1 := root.Left()
	assertTrue(n1.Value() == 200 && n1.Idx() == 1 && n1.from == 0 && n1.to == 2, t)
	n5 := root.Right()
	assertTrue(n5.Value() == 300 && n5.Idx() == 5 && n5.from == 4 && n5.to == 6, t)

	n2 := n1.Right()
	assertTrue(n2.Value() == 5 && n2.Idx() == 2 && n2.from == 2 && n2.to == 2, t)

	n0 := n1.Left()
	assertTrue(n0.value == arr[n0.Idx()] && n0.Idx() == 0 && n0.from == 0 && n0.to == 0, t)

	n4 := n5.Left()
	assertTrue(n4.value == arr[n4.Idx()] && n4.Idx() == 4 && n4.from == 4 && n4.to == 4, t)
	n6 := n5.Right()
	assertTrue(n6.value == arr[n6.Idx()] && n6.Idx() == 6 && n6.from == 6 && n6.to == 6, t)
}

func assertTrue(b bool, t *testing.T) {
	t.Helper()
	if !b {
		t.Error()
	}
}
