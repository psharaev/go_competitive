package fast_segment_tree_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/internal/ready_algorithms/fast_segment_tree"
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
	rnd := rand.New(rand.NewSource(int64(seed)))

	arr := genArray(rnd, 1, 30, -100, 100)

	testCase(t, rnd, arr)
}

func testBigSeed(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	arr := genArray(rnd, 100, 1000, -100, 100)

	testCase(t, rnd, arr)
}

func testCase(t *testing.T, rnd *rand.Rand, arr []int) {
	type segment struct {
		Val     int
		IsValid bool
	}

	st := fast_segment_tree.NewFastSegmentTree[int, segment](
		arr,
		func(item int) segment {
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
		cmd := genInt(rnd, 0, 1)
		const (
			set = iota
			sum
		)

		switch cmd {
		case set:
			idx := genInt(rnd, 0, len(arr)-1)
			val := genInt(rnd, -100, 100)
			arr[idx] = val
			st.SetVal(idx, val)
		case sum:
			l := genInt(rnd, 0, len(arr)-1)
			r := genInt(rnd, l, len(arr)-1)

			want := 0
			for i := l; i <= r; i++ {
				want += arr[i]
			}

			got := st.Sum(l, r, segment{IsValid: true}).Val

			require.Equal(t, want, got)
		}
	}
}

func genArray(r *rand.Rand, minSize, maxSize, minValue, maxValueInc int) []int {
	n := genInt(r, minSize, maxSize)
	if n == 0 {
		if r.Intn(2) == 0 {
			return nil
		}
		return []int{}
	}
	a := make([]int, n)
	for i := range a {
		a[i] = genInt(r, minValue, maxValueInc)
	}
	return a
}

func genInt(r *rand.Rand, min, maxInc int) int {
	return r.Intn(maxInc-min+1) + min
}
