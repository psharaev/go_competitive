package stack_test

import (
	"github.com/psharaev/go_competitive/dsa/stack"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"

	"github.com/psharaev/go_competitive/utils/generator"
)

func TestStress(t *testing.T) {
	for seed := 0; seed < 1000; seed++ {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			testSeed(t, seed)
		})
	}
}

func testCase(t *testing.T, gen *generator.Generator) {
	arr := make([]int, 1000)
	n := 0

	s := stack.NewStack[int](10)

	for range 1000 {
		cmd := gen.Int(0, 3)

		const (
			push = iota
			pop
			top
			size
		)

		switch cmd {
		case push:
			val := gen.Int(-1000, 1000)

			arr[n] = val
			n++

			s.Push(val)
		case pop:
			require.Equal(t, n, s.Size())
			require.Equal(t, n == 0, s.IsEmpty())
			if s.IsEmpty() {
				break
			}

			want := arr[n-1]
			n--
			got := s.Pop()
			require.Equal(t, want, got)
		case top:
			require.Equal(t, n, s.Size())
			require.Equal(t, n == 0, s.IsEmpty())
			if s.IsEmpty() {
				break
			}

			want := arr[n-1]
			got := s.Top()
			require.Equal(t, want, got)
		case size:
			require.Equal(t, n, s.Size())
			require.Equal(t, n == 0, s.IsEmpty())
		}
	}
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	testCase(t, gen)
}
