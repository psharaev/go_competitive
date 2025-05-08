package fenwick_tree_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/dsa/fenwick_tree"
	"github.com/stretchr/testify/require"
)

func TestStress(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testSeed(t, i)
		})
	}
}

func testCase(t *testing.T, rnd *rand.Rand, arr []int) {
	copied := SliceCopy(arr)
	ft := fenwick_tree.NewFenwickTree(arr)

	for range 1000 {
		cmd := genInt(rnd, 0, 2)
		const (
			sum = iota
			get
			set
			add
		)

		switch cmd {
		case sum:
			l := genInt(rnd, 0, len(arr)-1)
			r := genInt(rnd, l, len(arr)-1)

			want := 0
			for i := l; i <= r; i++ {
				want += copied[i]
			}

			got := ft.Sum(l, r)
			require.Equal(t, want, got, "sum")
		case get:
			idx := genInt(rnd, 0, len(arr)-1)
			require.Equal(t, copied[idx], ft.Get(idx), "get")
		case set:
			idx := genInt(rnd, 0, len(arr)-1)
			val := genInt(rnd, -50, 50)
			copied[idx] = val
			ft.Set(idx, val)
		case add:
			idx := genInt(rnd, 0, len(arr)-1)
			val := genInt(rnd, -50, 50)
			copied[idx] += val
			ft.Add(idx, val)
		}
	}
}

func testSeed(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	testCase(t, rnd, genArray(rnd, 1, 30, -100, 100))
	testCase(t, rnd, genArray(rnd, 1, 100_000, -100, 100))
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

func SliceCopy[T any](arr []T) []T {
	return append([]T(nil), arr...)
}
