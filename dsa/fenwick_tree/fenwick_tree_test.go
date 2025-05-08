package fenwick_tree_test

import (
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"
	"github.com/psharaev/go_competitive/utils/slice"

	"github.com/psharaev/go_competitive/dsa/fenwick_tree"
	"github.com/stretchr/testify/require"
)

func Test_FenwickTreeStress(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testSeed(t, i)
		})
	}
}

func testCase(t *testing.T, gen *generator.Generator, arr []int) {
	copied := slice.SliceCopy(arr)
	ft := fenwick_tree.NewFenwickTree(arr)

	for range 1000 {
		cmd := gen.Int(0, 2)
		const (
			sum = iota
			get
			set
			add
		)

		switch cmd {
		case sum:
			l := gen.Int(0, len(arr)-1)
			r := gen.Int(l, len(arr)-1)

			want := 0
			for i := l; i <= r; i++ {
				want += copied[i]
			}

			got := ft.Sum(l, r)
			require.Equal(t, want, got, "sum")
		case get:
			idx := gen.Int(0, len(arr)-1)
			require.Equal(t, copied[idx], ft.Get(idx), "get")
		case set:
			idx := gen.Int(0, len(arr)-1)
			val := gen.Int(-50, 50)
			copied[idx] = val
			ft.Set(idx, val)
		case add:
			idx := gen.Int(0, len(arr)-1)
			val := gen.Int(-50, 50)
			copied[idx] += val
			ft.Add(idx, val)
		}
	}
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	testCase(t, gen, gen.SliceInt(1, 30, -100, 100))
	testCase(t, gen, gen.SliceInt(1, 100_000, -100, 100))
}
