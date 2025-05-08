package fenwick_tree_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/psharaev/go_competitive/dsa/fenwick_tree"
	"github.com/psharaev/go_competitive/utils/generator"
)

func Test_FenwickTree3dStress(t *testing.T) {
	t.Parallel()

	for i := 0; i < 1000; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			testSeedFenwickTree3d(t, i)
		})
	}
}

func testCaseFenwickTree3d(t *testing.T, gen *generator.Generator) {
	n := gen.Int(1, 30)
	ft := fenwick_tree.NewFenwickTree3D(n)

	cube := make([][][]int, n)
	for i := range cube {
		cube[i] = make([][]int, n)
		for j := range cube[i] {
			cube[i][j] = make([]int, n)
		}
	}

	for range 1000 {
		cmd := gen.Int(1, 2)
		const (
			add = iota + 1
			sum
		)

		switch cmd {
		case add:
			x := gen.IntExc(0, n)
			y := gen.IntExc(0, n)
			z := gen.IntExc(0, n)
			val := gen.Int(-100, 100)

			cube[x][y][z] += val
			ft.Add(x, y, z, val)
		case sum:
			x1 := gen.IntExc(0, n)
			x2 := gen.IntExc(x1, n)
			y1 := gen.IntExc(0, n)
			y2 := gen.IntExc(y1, n)
			z1 := gen.IntExc(0, n)
			z2 := gen.IntExc(z1, n)

			want := 0
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					for z := z1; z <= z2; z++ {
						want += cube[x][y][z]
					}
				}
			}

			got := ft.SumCube(x1, x2, y1, y2, z1, z2)
			require.Equal(t, want, got)
		}
	}
}

func testSeedFenwickTree3d(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	testCaseFenwickTree3d(t, gen)
}
