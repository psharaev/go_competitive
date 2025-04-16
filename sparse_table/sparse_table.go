package sparse_table

type Block struct {
	item int
	idx  int
}

func (a Block) Merge(b Block) Block {
	if a.item >= b.item {
		return a
	}
	return b
}

type Mergable[T any] interface {
	Merge(T) T
}

type SparseTable[T Mergable[T]] struct {
	n        int
	logTable []int
	st       [][]T
}

func (st *SparseTable[T]) Query(l int, r int) T {
	length := r - l + 1
	level := st.logTable[length]
	block1 := st.st[level][l]
	block2 := st.st[level][r-(1<<level)+1]
	return block1.Merge(block2)
}

func CreateSparseTable[T Mergable[T]](arr []T) SparseTable[T] {
	n := len(arr)
	logTable := make([]int, n+1)
	logTable[1] = 0
	for i := 2; i <= n; i++ {
		logTable[i] = logTable[i/2] + 1
	}

	maxLevel := 0
	for (1 << maxLevel) <= n {
		maxLevel++
	}
	maxLevel--

	st := make([][]T, maxLevel+1)
	st[0] = make([]T, n)
	for i := 0; i < n; i++ {
		st[0][i] = arr[i]
	}

	for level := 1; level <= maxLevel; level++ {
		currentLen := 1 << level
		numElements := n - currentLen + 1
		if numElements <= 0 {
			break
		}
		st[level] = make([]T, numElements)
		step := 1 << (level - 1)
		for i := 0; i < numElements; i++ {
			left := st[level-1][i]
			right := st[level-1][i+step]
			st[level][i] = left.Merge(right)
		}
	}

	return SparseTable[T]{
		n:        n,
		logTable: logTable,
		st:       st,
	}
}
