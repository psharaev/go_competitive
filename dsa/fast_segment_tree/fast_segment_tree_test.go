package fast_segment_tree_test

import (
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"

	"github.com/psharaev/go_competitive/dsa/fast_segment_tree"
	"github.com/stretchr/testify/require"
)

func Test_FastSegmentTree(t *testing.T) {
	t.Parallel()

	for i := range 10_000 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testSmallSeed(t, i)
			testBigSeed(t, i)
		})
	}
}

func testSmallSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)
	arr := gen.SliceInt(1, 30, -100, 100)

	testCase(t, gen, arr)
}

func testBigSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	arr := gen.SliceInt(100, 1000, -100, 100)

	testCase(t, gen, arr)
}

func testCase(t *testing.T, gen *generator.Generator, arr []int) {
	type segment struct {
		Val     int
		IsValid bool
	}

	st := fast_segment_tree.NewFastSegmentTree[int, segment](
		arr,
		func(_ int, item int) segment {
			return segment{
				Val:     item,
				IsValid: true,
			}
		},
		func(a, b segment) segment {
			if !a.IsValid && !b.IsValid {
				return segment{}
			}

			if !a.IsValid {
				return b
			}
			if !b.IsValid {
				return a
			}

			return segment{
				Val:     a.Val + b.Val,
				IsValid: true,
			}
		},
	)

	for range 1000 {
		cmd := gen.Int(0, 1)
		const (
			set = iota
			sum
		)

		switch cmd {
		case set:
			idx := generator.Pos(gen, arr)
			val := gen.Int(-100, 100)
			arr[idx] = val
			st.SetVal(idx, val)
		case sum:
			l, r := generator.Segment(gen, arr)

			want := 0
			for i := l; i <= r; i++ {
				want += arr[i]
			}

			got := st.Sum(l, r, segment{IsValid: true}).Val

			require.Equal(t, want, got)
		}
	}
}
