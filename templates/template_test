package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"

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
	gen := generator.NewGenerator(seed)

	arr := gen.SliceInt(2, 6, 1, 100)
	testCase(t, arr)
}

func slowSolve(arr []int) int {
	return 0
}

func SliceCopy[T any](arr []T) []T {
	return append([]T(nil), arr...)
}
