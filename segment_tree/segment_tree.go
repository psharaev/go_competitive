package segment_tree

import "math/bits"

type Vertex struct {
	item int
	idx  int
}

func (a *Vertex) Merge(b *Vertex) *Vertex {
	if a.idx == 0 {
		return b
	}
	if b.idx == 0 {
		return a
	}

	if a.item > b.item {
		return a
	}

	return b
}

type Mergable[T any] interface {
	Merge(T) T
}

type SegmentTree[T Mergable[T]] struct {
	tree    []T
	n       int
	neutral T
}

func (t *SegmentTree[T]) Query(l, r int) T {
	return t.query(0, 0, t.n-1, l, r)
}

func (t *SegmentTree[T]) query(v, vl, vr, l, r int) T {
	if vr < l || vl > r {
		return t.neutral
	}
	if l <= vl && vr <= r {
		return t.tree[v]
	}

	m := (vl + vr) / 2
	lVertex := t.query(leftChild(v), vl, m, l, r)
	rVertex := t.query(rightChild(v), m+1, vr, l, r)

	return lVertex.Merge(rVertex)
}

func CreateSegmentTree[T Mergable[T]](arr []T, neutral T) SegmentTree[T] {
	n := NearestPowerOfTwo(len(arr))
	tree := make([]T, 2*n-1)
	for i := n - 1; i < 2*n-1; i++ {
		tree[i] = neutral
	}
	copy(tree[n-1:], arr)

	for i := n - 2; i >= 0; i-- {
		tree[i] = tree[leftChild(i)].Merge(tree[rightChild(i)])
	}

	return SegmentTree[T]{
		tree:    tree,
		n:       n,
		neutral: neutral,
	}
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func parentNode(i int) int {
	return (i - 1) / 2
}

func NearestPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 << bits.Len64(uint64(n-1))
}
