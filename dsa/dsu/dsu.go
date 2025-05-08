package dsu

type Dsu struct {
	parent []int
	rang   []int
}

func NewDsu(n int) *Dsu {
	parent := make([]int, n)
	for i := range n {
		parent[i] = i
	}

	return &Dsu{
		parent: parent,
		rang:   make([]int, n),
	}
}

func (d *Dsu) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *Dsu) Union(x, y int) {
	x = d.Find(x)
	y = d.Find(y)
	if x == y {
		return
	}

	if d.rang[x] > d.rang[y] {
		x, y = y, x
	}

	d.parent[x] = y
	if d.rang[y] == d.rang[x] {
		d.rang[y]++
	}
}
