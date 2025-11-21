package string_functions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PrefixFunctionString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		want []int
	}{
		{
			name: "",
			want: []int{-1},
		},
		{
			name: "a",
			want: []int{-1, 0},
		},
		{
			name: "abbababb",
			want: []int{-1, 0, 0, 0, 1, 2, 1, 2, 3},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := PrefixFunctionString(tc.name)

			require.Equal(t, tc.want, got)
		})
	}
}
