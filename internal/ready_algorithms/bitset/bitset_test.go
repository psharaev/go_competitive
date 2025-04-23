package bitset_test

import (
	"github.com/psharaev/go_competitive/internal/ready_algorithms/bitset"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
)

func TestStress(t *testing.T) {
	for i := range 1000 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			test(t, i)
		})
	}
	//test(t, 50)
}

func test(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	n := genInt(rnd, 1, 10000)
	bitset := bitset.NewBitset(n)
	m := map[int]bool{}

	k := genInt(rnd, 1, 100_000)
	for range k {
		pos := genInt(rnd, 0, n-1)
		op := genInt(rnd, 1, 3)
		switch op {
		case 0:
			actual := bitset.Get(pos)
			want := m[pos]
			require.Equal(t, want, actual)
		case 1:
			bitset.Clear(pos)
			m[pos] = false
		case 2:
			bitset.Set(pos)
			m[pos] = true
		}
	}
}

func genInt(r *rand.Rand, min, maxInc int) int {
	return r.Intn(maxInc-min+1) + min
}
