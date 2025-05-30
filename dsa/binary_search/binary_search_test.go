package binary_search_test

import (
	"github.com/psharaev/go_competitive/dsa/binary_search"
	"github.com/psharaev/go_competitive/utils/slice"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"
)

func TestStress(t *testing.T) {
	if !t.Run("manuals", manuals) {
		return
	}

	for seed := 0; seed < 1000; seed++ {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			testSeed(t, seed)
		})
	}
}

func manuals(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		testCase(t, []int{5}, 4)
	})
	t.Run("2", func(t *testing.T) {
		testCase(t, []int{5}, 5)
	})
	t.Run("3", func(t *testing.T) {
		testCase(t, []int{5}, 6)
	})
}

func testCase(t *testing.T, nums []int, target int) {
	wantLeft := left(nums, target)
	wantRight := right(nums, target)
	wantInsert := left(nums, target)
	gotLeft := binary_search.BinarySearchLeft(nums, target)
	gotRight := binary_search.BinarySearchRight(nums, target)
	gotInsert := binary_search.BinarySearchInsert(nums, target)
	assert.Equal(t, wantLeft, gotLeft, "fail left %s %d", slice.JoinSlice(nums, " ", "[", "]"), target)
	assert.Equal(t, wantRight, gotRight, "fail right %s %d", slice.JoinSlice(nums, " ", "[", "]"), target)
	assert.Equal(t, wantInsert, gotInsert, "fail insert %s %d", slice.JoinSlice(nums, " ", "[", "]"), target)
}

func left(nums []int, target int) int {
	// min i: arr[i] >= target
	if len(nums) == 0 {
		// a[-1] = -INF
		// a[0]  = +INF
		return 0
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < target {
			return i + 1
		}
	}
	return 0
}

func right(nums []int, target int) int {
	// max i: arr[i] <= target

	if len(nums) == 0 {
		// a[-1] = -INF
		// a[0]  = +INF
		return -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			return i - 1
		}
	}
	return len(nums) - 1
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	nums := gen.SliceInt(1, 100, -100, 100)
	target := gen.Int(-100, 100)
	sort.Ints(nums)

	testCase(t, nums, target)

	nums = gen.SliceInt(0, 10, -100, 100)
	target = gen.Int(-100, 100)
	if len(nums) > 1 {
		sort.Ints(nums)
	}

	testCase(t, nums, target)
}
