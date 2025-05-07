package slice_test

import (
	"fmt"
	"testing"

	"github.com/psharaev/go_competitive/internal/ready_algorithms/slice"
	"github.com/stretchr/testify/require"
)

func TestFillSlice(t *testing.T) {
	t.Parallel()

	type testCase struct {
		arr []int
		val int
	}
	cases := []testCase{
		{
			arr: nil,
			val: 2,
		},
		{
			arr: []int{},
			val: 3,
		},
	}

	for l := 1; l < 1000; l++ {
		cases = append(cases, testCase{
			arr: make([]int, l),
			val: 42,
		})
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("%v", len(tc.arr)), func(t *testing.T) {
			slice.FillSlice(tc.arr, tc.val)

			for _, item := range tc.arr {
				require.Equal(t, tc.val, item)
			}
		})
	}
}
