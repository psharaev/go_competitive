package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeed_50(t *testing.T) {
	testSeed(t, 50)
}

func TestStress(t *testing.T) {
	if !t.Run("examples", examples) {
		return
	}

	if !t.Run("manuals", manuals) {
		return
	}

	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testSeed(t, i)
		})
	}
}

func examples(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		testCase(t, []int{1})
	})
}

func manuals(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		testCase(t, []int{1})
	})

	seeds := []int{
		50,
	}
	for _, seed := range seeds {
		t.Run(fmt.Sprintf("seed %d", seed), func(t *testing.T) {
			testSeed(t, seed)
		})
	}
}

func testCase(t *testing.T, arr []int) {
	want := slowSolve(arr)
	got := solve(arr)

	if !assert.Equal(t, want, got) {
		t.Logf("%#v", arr)
	}
}

func testSeed(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	arr := genArray(rnd, 2, 6, 1, 100)
	testCase(t, arr)
}

func slowSolve(arr []int) int {
	return 0
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
