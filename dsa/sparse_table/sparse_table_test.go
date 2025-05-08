package sparse_table_test

import (
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/dsa/sparse_table"
	"github.com/stretchr/testify/require"

	"github.com/psharaev/go_competitive/utils/generator"
)

func TestStress(t *testing.T) {
	for seed := 0; seed < 1000; seed++ {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			testSeed(t, seed)
		})
	}
}

func testCase(t *testing.T, gen *generator.Generator, arr []int) {
	type segment struct {
		valMin int
		pos    int
	}

	st := sparse_table.NewSparseTable(arr,
		func(idx int, item int) segment {
			return segment{
				valMin: item,
				pos:    idx,
			}
		},
		func(a, b segment) segment {
			if a.valMin <= b.valMin {
				return a
			}
			return b
		},
	)

	for range 1000 {
		l, r := generator.Segment(gen, arr)

		want := arr[l]
		wantPos := l
		for i := l + 1; i <= r; i++ {
			if want > arr[i] {
				want = arr[i]
				wantPos = i
			}
		}

		got := st.Min(l, r)

		require.Equal(t, want, got.valMin)
		require.Equal(t, wantPos, got.pos)
	}
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	testCase(t, gen, gen.SliceInt(1, 20, -50, 50))
	testCase(t, gen, gen.SliceInt(20, 10_000, -500, 500))
}
