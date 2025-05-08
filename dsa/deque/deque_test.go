package deque_test

import (
	"fmt"
	"github.com/psharaev/go_competitive/dsa/deque"
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
	countOps := 1000

	infDeque := newSlowDeque(4 * countOps)
	fastDeque := deque.NewDeque[int](gen.Int(0, 10))

	for range countOps {
		cmd := gen.Int(0, 5)

		const (
			pushBack = iota
			pushFront
			popBack
			popFront
			borders
			state
		)

		switch cmd {
		case pushBack:
			val := gen.Int(-1000, 1000)
			infDeque.PushBack(val)
			fastDeque.PushBack(val)
		case pushFront:
			val := gen.Int(-1000, 1000)
			infDeque.PushFront(val)
			fastDeque.PushFront(val)
		case popBack:
			if isEmpty(t, &infDeque, &fastDeque) {
				break
			}

			infDeque.PopBack()
			fastDeque.PopBack()
		case popFront:
			if isEmpty(t, &infDeque, &fastDeque) {
				break
			}
			infDeque.PopFront()
			fastDeque.PopFront()
		case borders:
			if isEmpty(t, &infDeque, &fastDeque) {
				break
			}

			require.Equal(t, infDeque.Front(), fastDeque.Front())
			require.Equal(t, infDeque.Back(), fastDeque.Back())
		case state:
			isEmpty(t, &infDeque, &fastDeque)
		}
	}
}

func isEmpty(t *testing.T, slowDeque *infinityDeque, fastDeque *deque.Deque[int]) bool {
	require.Equal(t, slowDeque.Size(), fastDeque.Size())
	require.Equal(t, slowDeque.Empty(), fastDeque.Empty())
	require.Equal(t, slowDeque.String(), fastDeque.String())
	return slowDeque.Empty()
}

func testSeed(t *testing.T, seed int) {
	gen := generator.NewGenerator(seed)
	testCase(t, gen)
}

type infinityDeque struct {
	arr  []int
	head int
	tail int
}

func newSlowDeque(cap int) infinityDeque {
	return infinityDeque{
		arr:  make([]int, cap),
		head: cap / 2,
		tail: cap / 2,
	}
}

func (d *infinityDeque) Size() int {
	return d.tail - d.head
}

func (d *infinityDeque) Empty() bool {
	return d.Size() == 0
}

func (d *infinityDeque) Front() int {
	return d.arr[d.head]
}

func (d *infinityDeque) Back() int {
	return d.arr[d.tail-1]
}

func (d *infinityDeque) PushBack(item int) {
	d.arr[d.tail] = item
	d.tail++
}

func (d *infinityDeque) PushFront(item int) {
	d.arr[d.head-1] = item
	d.head--
}

func (d *infinityDeque) PopFront() int {
	d.head++
	return d.arr[d.head-1]
}

func (d *infinityDeque) PopBack() int {
	d.tail--
	return d.arr[d.tail]
}

func (d *infinityDeque) String() string {
	return fmt.Sprintf("%v", d.arr[d.head:d.tail])
}
