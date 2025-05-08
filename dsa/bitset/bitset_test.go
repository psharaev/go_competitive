package bitset_test

import (
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"

	"github.com/psharaev/go_competitive/dsa/bitset"
	"github.com/stretchr/testify/require"
)

func TestStress(t *testing.T) {
	t.Parallel()

	for seed := range 1000 {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			test(t, seed)
		})
	}
}

func test(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	n := gen.Int(1, 10000)
	bs := bitset.NewBitset(n)
	m := map[int]bool{}

	k := gen.Int(1, 100_000)
	for range k {
		pos := gen.Int(0, n-1)
		op := gen.Int(1, 3)
		switch op {
		case 0:
			actual := bs.Get(pos)
			want := m[pos]
			require.Equal(t, want, actual)
		case 1:
			bs.Clear(pos)
			m[pos] = false
		case 2:
			bs.Set(pos)
			m[pos] = true
		}
	}
}
