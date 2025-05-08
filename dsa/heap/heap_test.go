package heap_test

import (
	"github.com/psharaev/go_competitive/dsa/heap"
	"github.com/psharaev/go_competitive/utils/slice"
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

func testCase(t *testing.T, gen *generator.Generator, arr []int) {
	slowH := newSlowHeap(arr)
	fastH := heap.NewHeap(arr, func(a, b int) int {
		return a - b
	})

	for range 1000 {
		cmd := gen.Int(0, 3)
		const (
			push = iota
			pop
			peek
			size
		)
		switch cmd {
		case push:
			val := gen.Int(-1000, 1000)
			slowH.push(val)
			fastH.Push(val)
		case pop:
			require.Equal(t, slowH.size(), fastH.Size())
			require.Equal(t, slowH.isEmpty(), fastH.IsEmpty())
			if fastH.IsEmpty() {
				break
			}

			want := slowH.pop()
			got := fastH.Pop()
			require.Equal(t, want, got)
		case peek:
			require.Equal(t, slowH.size(), fastH.Size())
			require.Equal(t, slowH.isEmpty(), fastH.IsEmpty())
			if fastH.IsEmpty() {
				break
			}

			want := slowH.peek()
			got := fastH.Peek()
			require.Equal(t, want, got)
		case size:
			require.Equal(t, slowH.size(), fastH.Size())
			require.Equal(t, slowH.isEmpty(), fastH.IsEmpty())
		}
	}
}

type slowHeap struct {
	arr []int
}

func newSlowHeap(arr []int) slowHeap {
	return slowHeap{
		arr: slice.SliceCopy(arr),
	}
}

func (h *slowHeap) push(val int) {
	h.arr = append(h.arr, val)
}

func (h *slowHeap) pop() int {
	posMin := 0
	for i := 1; i < len(h.arr); i++ {
		if h.arr[i] < h.arr[posMin] {
			posMin = i
		}
	}

	val := h.arr[posMin]
	h.arr = slice.RemoveItem(h.arr, posMin)
	return val
}

func (h *slowHeap) peek() int {
	posMin := 0
	for i := 1; i < len(h.arr); i++ {
		if h.arr[i] < h.arr[posMin] {
			posMin = i
		}
	}

	return h.arr[posMin]
}

func (h *slowHeap) size() int {
	return len(h.arr)
}

func (h *slowHeap) isEmpty() bool {
	return len(h.arr) == 0
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)

	testCase(t, gen, gen.SliceInt(0, 30, -100, 100))
	testCase(t, gen, gen.SliceInt(100, 1000, -1000, 1000))
}
