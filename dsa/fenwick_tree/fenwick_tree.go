package fenwick_tree

type FenwickTree struct {
	tree []int
}

func NewFenwickTree(arr []int) *FenwickTree {
	ft := &FenwickTree{tree: make([]int, len(arr))}
	for i, v := range arr {
		ft.Add(i, v)
	}
	return ft
}

func (t *FenwickTree) Sum(l, r int) int {
	return t.sumPrefix(r) - t.sumPrefix(l-1)
}

func (t *FenwickTree) sumPrefix(r int) int {
	sum := 0
	for ; r >= 0; r = r&(r+1) - 1 {
		sum += t.tree[r]
	}
	return sum
}

func (t *FenwickTree) Add(idx int, val int) {
	for ; idx < len(t.tree); idx = idx | (idx + 1) {
		t.tree[idx] += val
	}
}

func (t *FenwickTree) Set(idx int, val int) {
	t.Add(idx, val-t.Get(idx))
}

func (t *FenwickTree) Get(idx int) int {
	return t.Sum(idx, idx)
}
