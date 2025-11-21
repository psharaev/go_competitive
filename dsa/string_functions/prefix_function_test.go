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

func Test_FindPattern(t *testing.T) {
	t.Parallel()

	cases := []struct {
		text    string
		pattern string
		symbol  string
		want    []int
	}{
		{
			text:    "",
			pattern: "a",
			symbol:  "#",
			want:    []int{},
		},
		{
			text:    "a",
			pattern: "",
			symbol:  "#",
			want:    []int{-1, 0},
		},
		{
			text:    "",
			pattern: "",
			symbol:  "#",
			want:    []int{-1},
		},
		{
			text:    "a",
			pattern: "a",
			symbol:  "#",
			want:    []int{0},
		},
		{
			text:    "aba",
			pattern: "a",
			symbol:  "#",
			want:    []int{0, 2},
		},
		{
			text:    "aba",
			pattern: "b",
			symbol:  "#",
			want:    []int{1},
		},
	}

	for _, tc := range cases {
		name := tc.pattern + tc.symbol + tc.text
		t.Run(name, func(t *testing.T) {
			got := FindPattern(tc.text, tc.pattern, tc.symbol)

			require.Equal(t, tc.want, got)
		})
	}
}
