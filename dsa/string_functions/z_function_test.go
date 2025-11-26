package string_functions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ZFunctionString(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name string
		want []int
	}{
		{
			name: "",
			want: []int{},
		},
		{
			name: "a",
			want: []int{1},
		},
		{
			name: "abbaabababb",
			want: []int{11, 0, 0, 1, 2, 0, 2, 0, 3, 0, 0},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ZFunctionString(tc.name)

			require.Equal(t, tc.want, got)
		})
	}
}
