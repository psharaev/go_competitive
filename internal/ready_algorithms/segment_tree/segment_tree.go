package segment_tree

import (
	"fmt"
	"math/bits"
	"strings"
)

type Block struct {
	item    int
	idx     int
	isValid bool
}

func (a Block) Merge(b Block) Block {
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

type SegmentTree struct {
	tree    []Block
	n       int
	neutral Block
}

func (t *SegmentTree) Query(l, r int) Block {
	return t.query(0, 0, t.n-1, l, r)
}

func (t *SegmentTree) query(v, vl, vr, l, r int) Block {
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

func CreateSegmentTree(arr []Block, neutral Block) SegmentTree {
	n := NearestPowerOfTwo(len(arr))
	tree := make([]Block, 2*n-1)
	for i := n - 1; i < 2*n-1; i++ {
		tree[i] = neutral
	}
	copy(tree[n-1:], arr)

	for i := n - 2; i >= 0; i-- {
		tree[i] = tree[leftChild(i)].Merge(tree[rightChild(i)])
	}

	return SegmentTree{
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

func (t *SegmentTree) String() string {
	var sb strings.Builder
	t.printTree(0, 0, t.n-1, "", &sb)
	return sb.String()
}

func (t *SegmentTree) printTree(v, vl, vr int, prefix string, sb *strings.Builder) {
	var desc string
	if vl == vr {
		desc = fmt.Sprintf("Node v %d arr[%d] block %v", v, vl, t.tree[v])
	} else {
		desc = fmt.Sprintf("Node v %d vl %d vr %d block %v", v, vl, vr, t.tree[v])
	}
	if !t.tree[v].isValid {
		desc = "invalid " + desc
	}
	sb.WriteString(desc)
	sb.WriteByte('\n')

	type childS struct {
		v  int
		vl int
		vr int
	}

	m := (vl + vr) / 2
	childrens := make([]childS, 0, 2)
	if leftChild(v) < len(t.tree) {
		childrens = append(childrens, childS{
			v:  leftChild(v),
			vl: vl,
			vr: m,
		})
	}

	if rightChild(v) < len(t.tree) {
		childrens = append(childrens, childS{
			v:  rightChild(v),
			vl: m + 1,
			vr: vr,
		})
	}

	for i, child := range childrens {
		connector := "├── "
		newPrefix := "│   "
		if i == len(childrens)-1 {
			connector = "└── "
			newPrefix = "    "
		}

		sb.WriteString(prefix)
		sb.WriteString(connector)
		t.printTree(child.v, child.vl, child.vr, prefix+newPrefix, sb)
	}
}
