package set

type Set interface {
	Size() int
	IsEmpty() bool
	Contains(e interface{}) bool
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
}

type HashSet struct {
	m map[interface{}]struct{}
}

func NewHashSet() Set {
	return &HashSet{
		m: make(map[interface{}]struct{}),
	}
}

func (set *HashSet) Size() int {
	return len(set.m)
}

func (set *HashSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *HashSet) Contains(e interface{}) bool {
	_, ok := set.m[e]
	return ok
}

func (set *HashSet) Clear() {
	set.m = make(map[interface{}]struct{})
}

func (set *HashSet) Add(e interface{}) bool {
	_, ok := set.m[e]
	if !ok {
		set.m[e] = struct{}{}
	}
	return ok
}

func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}
