package union_find

type UnionFind struct {
	n          int
	components int
	parent     []int
	size       []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}

	return &UnionFind{
		n:          n,
		components: n,
		parent:     parent,
		size:       size,
	}
}

func (uf *UnionFind) Union(a int, b int) {
	if a < 0 || a >= uf.n || b < 0 || b >= uf.n {
		return
	}
	aRoot := uf.FindRoot(a)
	bRoot := uf.FindRoot(b)
	if aRoot == bRoot {
		return
	}
	if uf.size[aRoot] < uf.size[bRoot] {
		uf.parent[aRoot] = bRoot
		uf.size[bRoot] += uf.size[aRoot]
	} else {
		uf.parent[bRoot] = aRoot
		uf.size[aRoot] += uf.size[bRoot]
	}
	uf.components--
}

func (uf *UnionFind) Connected(a int, b int) bool {
	if a < 0 || a >= uf.n || b < 0 || b >= uf.n {
		return false
	}
	return uf.FindRoot(a) == uf.FindRoot(b)
}

func (uf *UnionFind) FindRoot(a int) int {
	if a < 0 || a >= uf.n {
		return -1
	}
	root := a
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for a != root {
		newA := uf.parent[a]
		uf.parent[a] = root
		a = newA
	}
	return root
}

func (uf *UnionFind) CountComponents() int {
	return uf.components
}
