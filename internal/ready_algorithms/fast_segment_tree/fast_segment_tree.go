package fast_segment_tree

import (
	"fmt"
	"math/bits"
	"strings"
)

type FastSegmentTree[Item any, Segment any] struct {
	tree []Segment
	n    int

	convert func(item Item) Segment
	merge   func(a, b Segment) Segment
}

func (t *FastSegmentTree[Item, Segment]) Sum(l, r int, res Segment) Segment {
	l = l + t.n - 1
	r = r + t.n - 1

	acc2 := res

	for l <= r {
		if l&1 == 0 {
			res = t.merge(res, t.tree[l])
		}
		l >>= 1

		if r&1 == 1 {
			acc2 = t.merge(t.tree[r], acc2)
		}
		r = (r >> 1) - 1
	}

	return t.merge(res, acc2)
}

func (t *FastSegmentTree[Item, Segment]) SetVal(idx int, val Item) {
	idx = idx + t.n - 1
	t.tree[idx] = t.convert(val)
	for idx != 0 {
		idx = (idx - 1) >> 1
		t.tree[idx] = t.merge(t.tree[(2*idx+1)], t.tree[(2*idx+2)])
	}
}

func NewFastSegmentTree[Item any, Segment any](arr []Item, convert func(item Item) Segment, merge func(a, b Segment) Segment) FastSegmentTree[Item, Segment] {
	n := NearestPowerOfTwo(len(arr))
	tree := make([]Segment, 2*n-1)
	for i := 0; i < len(arr); i++ {
		tree[n-1+i] = convert(arr[i])
	}

	for i := n - 2; i >= 0; i-- {
		tree[i] = merge(tree[(2*i+1)], tree[(2*i+2)])
	}

	return FastSegmentTree[Item, Segment]{
		tree:    tree,
		n:       n,
		convert: convert,
		merge:   merge,
	}
}

func NearestPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	return 1 << bits.Len64(uint64(n-1))
}

func (t *FastSegmentTree[Item, Segment]) String() string {
	var sb strings.Builder
	t.printTree(0, 0, t.n-1, "", &sb)
	return sb.String()
}

func (t *FastSegmentTree[Item, Segment]) printTree(v, vl, vr int, prefix string, sb *strings.Builder) {
	var desc string
	if vl == vr {
		desc = fmt.Sprintf("Node v %d arr[%d] segment %v", v, vl, t.tree[v])
	} else {
		desc = fmt.Sprintf("Node v %d vl %d vr %d segment %v", v, vl, vr, t.tree[v])
	}
	sb.WriteString(desc)
	sb.WriteByte('\n')

	type childS struct {
		v  int
		vl int
		vr int
	}

	m := (vl + vr) / 2
	childs := make([]childS, 0, 2)
	if 2*v+1 < len(t.tree) {
		childs = append(childs, childS{
			v:  2*v + 1,
			vl: vl,
			vr: m,
		})
	}

	if 2*v+2 < len(t.tree) {
		childs = append(childs, childS{
			v:  2*v + 2,
			vl: m + 1,
			vr: vr,
		})
	}

	for i, child := range childs {
		connector := "├── "
		newPrefix := "│   "
		if i == len(childs)-1 {
			connector = "└── "
			newPrefix = "    "
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		t.printTree(child.v, child.vl, child.vr, prefix+newPrefix, sb)
	}
}
