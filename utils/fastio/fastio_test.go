package fastio_test

import (
	"github.com/psharaev/go_competitive/utils/fastio"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func Test_Fastio(t *testing.T) {
	cases := []struct {
		name  string
		check func(t *testing.T, r fastio.FastReader)
	}{
		{
			name: "123",
			check: func(t *testing.T, r fastio.FastReader) {
				require.Equal(t, 123, r.NextInt())
			},
		},
		{
			name: "123",
			check: func(t *testing.T, r fastio.FastReader) {
				require.Equal(t, "123", r.NextWord())
			},
		},
		{
			name: "line1 aboba\n line2 aboba \n line3 ",
			check: func(t *testing.T, r fastio.FastReader) {
				require.Equal(t, "line1 aboba", r.NextLine())
				require.Equal(t, " line2 aboba ", r.NextLine())
				require.Equal(t, " line3 ", r.NextLine())
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.check(t, fastio.NewFastReader(strings.NewReader(tc.name)))
			tc.check(t, fastio.NewFastReader(strings.NewReader(tc.name+"\n")))
			tc.check(t, fastio.NewFastReader(strings.NewReader(tc.name+"\r\n")))
		})
	}
}
