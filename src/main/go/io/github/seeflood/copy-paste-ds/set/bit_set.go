package set

type BitSet struct {
	words       []uint64
	cardinality int
}

func NewBitSet() *BitSet {
	return &BitSet{
		nil,
		0,
	}
}

func (set *BitSet) Size() int {
	return len(set.words)
}

func (set *BitSet) IsEmpty() bool {
	return set.Size() == 0
}

func (set *BitSet) Contains(x int) bool {
	if x < 0 {
		return false
	}
	word, bit := x/64, uint(x%64)
	return word < len(set.words) && set.words[word]&(1<<bit) != 0
}

func (set *BitSet) Clear() {
	set.words = make([]uint64, 0)
	set.cardinality = 0
}

func (set *BitSet) Add(x int) bool {
	if x < 0 {
		return false
	}
	word, bit := x/64, uint(x%64)
	for word >= len(set.words) {
		set.words = append(set.words, 0)
	}

	exist := set.words[word]&(1<<bit) != 0
	if !exist {
		set.words[word] |= (1 << bit)
		set.cardinality++
	}
	return exist
}

func (set *BitSet) Remove(x int) {
	if x < 0 {
		return
	}
	word, bit := x/64, uint(x%64)
	if word >= len(set.words) {
		return
	}
	exist := set.words[word]&(1<<bit) != 0
	if !exist {
		return
	}
	set.words[word] &= (^(1 << bit))
	set.cardinality--
}

/*
  Returns true if the specified {@code BitSet} has any bits set to
 {@code true} that are also set to {@code true} in this {@code BitSet}.
*/
func (set *BitSet) Intersects(other *BitSet) bool {
	for i := min(len(set.words), len(other.words)) - 1; i >= 0; i-- {
		if set.words[i]&other.words[i] != 0 {
			return true
		}
	}
	return false
}

/*
  Performs a logical <b>OR</b> of this bit set with the bit set
  argument. This bit set is modified so that a bit in it has the
  value {@code true} if and only if it either already had the
  value {@code true} or the corresponding bit in the bit set
  argument has the value {@code true}.
*/
func (set *BitSet) Or(other *BitSet) {
	for i, oword := range other.words {
		if i < len(set.words) {
			old := bitCount(set.words[i])
			set.words[i] |= oword
			set.cardinality += (bitCount(set.words[i]) - old)
		} else {
			set.words = append(set.words, oword)
			set.cardinality += bitCount(set.words[i])
		}
	}
}

func bitCount(x uint64) int {
	cnt := 0
	for x != 0 {
		lowbit := x & -x
		x -= lowbit
		cnt++
	}
	return cnt
}

/*
  Returns the number of bits set to {@code true} in this {@code BitSet}.
*/
func (set *BitSet) Cardinality() int {
	return set.cardinality
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
