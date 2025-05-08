package sparse_table

import "math/bits"

type SparseTable[Segment any] struct {
	lookup [][]Segment

	merge func(a, b Segment) Segment
}

func (st *SparseTable[Segment]) Min(l int, r int) Segment {
	length := r - l + 1
	level := bits.Len(uint(length)) - 1
	return st.merge(
		st.lookup[level][l],
		st.lookup[level][r-(1<<level)+1],
	)
}

func NewSparseTable[Item any, Segment any](
	arr []Item,
	convert func(int, Item) Segment,
	merge func(Segment, Segment) Segment,
) SparseTable[Segment] {
	n := len(arr)

	maxLevel := bits.Len(uint(n)) - 1

	lookup := make([][]Segment, maxLevel+1)
	lookup[0] = make([]Segment, n)
	for i, item := range arr {
		lookup[0][i] = convert(i, item)
	}

	for level := 1; level <= maxLevel; level++ {
		prevStep := 1 << (level - 1)
		lookup[level] = make([]Segment, n-(1<<level)+1)
		for i := range lookup[level] {
			lookup[level][i] = merge(
				lookup[level-1][i],
				lookup[level-1][i+prevStep],
			)
		}
	}

	return SparseTable[Segment]{
		lookup: lookup,
		merge:  merge,
	}
}
