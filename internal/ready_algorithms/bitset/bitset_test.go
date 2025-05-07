package bitset_test

import (
	"github.com/psharaev/go_competitive/internal/ready_algorithms/bitset"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

func TestStress(t *testing.T) {
	t.Parallel()

	for i := range 1000 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			test(t, i)
		})
	}
}

func test(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	n := genInt(rnd, 1, 10000)
	bs := bitset.NewBitset(n)
	m := map[int]bool{}

	k := genInt(rnd, 1, 100_000)
	for range k {
		pos := genInt(rnd, 0, n-1)
		op := genInt(rnd, 1, 3)
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

func genInt(r *rand.Rand, min, maxInc int) int {
	return r.Intn(maxInc-min+1) + min
}
