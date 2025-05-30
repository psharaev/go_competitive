package sort_test

import (
	"github.com/psharaev/go_competitive/dsa/sort"
	"github.com/psharaev/go_competitive/utils/slice"
	"github.com/stretchr/testify/assert"
	"slices"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"
)

func TestStress(t *testing.T) {
	for seed := 0; seed < 1000; seed++ {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			testSeed(t, seed)
		})
	}
}

func testCase(t *testing.T, nums []int) {
	want := slice.SliceCopy(nums)
	slices.Sort(want)

	gotMerge := slice.SliceCopy(nums)
	sort.MergeSort(gotMerge)

	assert.Equal(t, want, gotMerge, "merge sort")
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	nums := gen.SliceInt(0, 100, -100, 100)

	testCase(t, nums)
}
