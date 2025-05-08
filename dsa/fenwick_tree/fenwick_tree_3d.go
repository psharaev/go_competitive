package fenwick_tree

type FenwickTree3D struct {
	tree [][][]int
	n    int
}

func NewFenwickTree3D(n int) FenwickTree3D {
	tree := make([][][]int, n)
	for i := range tree {
		tree[i] = make([][]int, n)
		for j := range tree[i] {
			tree[i][j] = make([]int, n)
		}
	}
	return FenwickTree3D{tree: tree, n: n}
}

func (ft *FenwickTree3D) Add(x, y, z, delta int) {
	n := ft.n
	for i := x; i < n; i = i | (i + 1) {
		for j := y; j < n; j = j | (j + 1) {
			for k := z; k < n; k = k | (k + 1) {
				ft.tree[i][j][k] += delta
			}
		}
	}
}

func (ft *FenwickTree3D) SumPrefix(x, y, z int) int {
	res := 0
	for i := x; i >= 0; i = (i & (i + 1)) - 1 {
		for j := y; j >= 0; j = (j & (j + 1)) - 1 {
			for k := z; k >= 0; k = (k & (k + 1)) - 1 {
				res += ft.tree[i][j][k]
			}
		}
	}
	return res
}

func (ft *FenwickTree3D) SumCube(x1 int, x2 int, y1 int, y2 int, z1 int, z2 int) int {
	a := ft.SumPrefix(x2, y2, z2)
	b := ft.SumPrefix(x1-1, y2, z2)
	c := ft.SumPrefix(x2, y1-1, z2)
	d := ft.SumPrefix(x2, y2, z1-1)
	e := ft.SumPrefix(x1-1, y1-1, z2)
	f := ft.SumPrefix(x1-1, y2, z1-1)
	g := ft.SumPrefix(x2, y1-1, z1-1)
	h := ft.SumPrefix(x1-1, y1-1, z1-1)
	return a - b - c - d + e + f + g - h
}
