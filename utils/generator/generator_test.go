package generator_test

import (
	"math"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"
	"github.com/stretchr/testify/require"
)

func Test_IntInc(t *testing.T) {
	t.Parallel()

	gen := generator.NewGenerator(42)
	set := map[int]bool{}

	for range 1000 {
		val := gen.Int(-5, 5)

		require.GreaterOrEqual(t, val, -5)
		require.LessOrEqual(t, val, 5)

		set[val] = true
	}

	require.Equal(t, 11, len(set))
}

func Test_IntExc(t *testing.T) {
	t.Parallel()

	gen := generator.NewGenerator(42)
	set := map[int]bool{}

	for range 1000 {
		val := gen.IntExc(-5, 5)
		require.GreaterOrEqual(t, val, -5)
		require.Less(t, val, 5)

		set[val] = true
	}

	require.Equal(t, 10, len(set))
}

func Test_Bool(t *testing.T) {
	t.Parallel()

	gen := generator.NewGenerator(42)
	countTrue := 0
	countTest := 1000
	p := 0.3

	for range countTest {
		if gen.Bool(p) {
			countTrue++
		}
	}

	require.NotEqual(t, countTrue, countTest)

	gotP := float64(countTrue) / float64(countTest)

	require.Less(t, math.Abs(gotP-p), 0.02)
}

func Test_SliceInt(t *testing.T) {
	t.Parallel()

	gen := generator.NewGenerator(42)

	sizes := map[int]bool{}
	values := map[int]bool{}

	for range 1000 {
		s := gen.SliceInt(0, 10, -5, 5)

		require.LessOrEqual(t, len(s), 10)
		sizes[len(s)] = true

		for _, item := range s {
			require.GreaterOrEqual(t, item, -5)
			require.LessOrEqual(t, item, 5)

			values[item] = true
		}
	}

	require.Equal(t, 11, len(sizes))
	require.Equal(t, 11, len(values))
}
