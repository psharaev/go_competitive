package deque_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go_competitive/deque"
	"math/rand/v2"
	"testing"
)

const (
	Size = iota
	Empty

	Front
	Back

	PushBack
	PushFront

	PopBack
	PopFront
)

var _ deq = deque.NewDeque(0)
var _ deq = &slowDeque{}

type deq interface {
	Size() int
	Empty() bool

	Front() int
	Back() int

	PushBack(int)
	PushFront(int)

	PopBack() int
	PopFront() int
}

func Test_Stress(t *testing.T) {
	t.Parallel()

	n := 10_000
	r := rand.New(&s{})
	for range n {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			countOps := 10_000
			m := &MergedDeque{
				slowDeque: newSlowDeque(4 * countOps),
				fastDeque: deque.NewDeque(r.IntN(10)),
			}

			for range countOps {
				op := r.IntN(8)
				switch op {
				case Size:
					m.Size(t, r)
				case Empty:
					m.Empty(t, r)
				case Front:
					m.Front(t, r)
				case Back:
					m.Back(t, r)
				case PushBack:
					m.PushBack(t, r)
				case PushFront:
					m.PushFront(t, r)
				case PopBack:
					m.PopBack(t, r)
				case PopFront:
					m.PopFront(t, r)
				default:
					panic(op)
				}
			}
		})
	}
}

type MergedDeque struct {
	slowDeque *slowDeque
	fastDeque *deque.Deque
}

func (m *MergedDeque) Size(t *testing.T, _ *rand.Rand) int {
	want := m.slowDeque.Size()
	actual := m.fastDeque.Size()
	require.Equal(t, want, actual)
	return want
}

func (m *MergedDeque) Empty(t *testing.T, _ *rand.Rand) bool {
	want := m.slowDeque.Empty()
	actual := m.fastDeque.Empty()
	require.Equal(t, want, actual)
	return want
}

func (m *MergedDeque) Front(t *testing.T, r *rand.Rand) int {
	if m.Empty(t, r) {
		if r.Float64() < 0.5 {
			m.PushBack(t, r)
		} else {
			m.PushFront(t, r)
		}
	}
	want := m.slowDeque.Front()
	actual := m.fastDeque.Front()
	require.Equal(t, want, actual)
	return want
}

func (m *MergedDeque) Back(t *testing.T, r *rand.Rand) int {
	m.ensureElem(t, r)
	want := m.slowDeque.Back()
	actual := m.fastDeque.Back()
	require.Equal(t, want, actual)
	return want
}

func (m *MergedDeque) ensureElem(t *testing.T, r *rand.Rand) {
	if m.Empty(t, r) {
		if r.Float64() < 0.5 {
			m.PushBack(t, r)
		} else {
			m.PushFront(t, r)
		}
	}
}

func (m *MergedDeque) PushBack(t *testing.T, r *rand.Rand) {
	v := r.IntN(1000)
	m.slowDeque.PushBack(v)
	m.fastDeque.PushBack(v)
	m.validateState(t, r)
}

func (m *MergedDeque) validateState(t *testing.T, r *rand.Rand) {
	if !m.Empty(t, r) {
		m.Front(t, r)
		m.Back(t, r)
	}
	m.Size(t, r)
}

func (m *MergedDeque) PushFront(t *testing.T, r *rand.Rand) {
	v := r.IntN(1000)
	m.slowDeque.PushFront(v)
	m.fastDeque.PushFront(v)

	m.validateState(t, r)
}

func (m *MergedDeque) PopBack(t *testing.T, r *rand.Rand) int {
	m.ensureElem(t, r)
	want := m.slowDeque.PopBack()
	actual := m.fastDeque.PopBack()
	require.Equal(t, want, actual)
	m.validateState(t, r)
	return want
}

func (m *MergedDeque) PopFront(t *testing.T, r *rand.Rand) int {
	m.ensureElem(t, r)
	want := m.slowDeque.PopFront()
	actual := m.fastDeque.PopFront()
	require.Equal(t, want, actual)
	m.validateState(t, r)
	return want
}

type s struct{}

func (s *s) Uint64() uint64 {
	return 42
}

type slowDeque struct {
	arr  []int
	head int
	tail int
}

func newSlowDeque(cap int) *slowDeque {
	return &slowDeque{
		arr:  make([]int, cap),
		head: cap / 2,
		tail: cap / 2,
	}
}

func (d *slowDeque) Size() int {
	return d.tail - d.head
}

func (d *slowDeque) Empty() bool {
	return d.Size() == 0
}

func (d *slowDeque) Front() int {
	return d.arr[d.head]
}

func (d *slowDeque) Back() int {
	return d.arr[d.tail-1]
}

func (d *slowDeque) PushBack(item int) {
	d.arr[d.tail] = item
	d.tail++
}

func (d *slowDeque) PushFront(item int) {
	d.arr[d.head-1] = item
	d.head--
}

func (d *slowDeque) PopFront() int {
	d.head++
	return d.arr[d.head-1]
}

func (d *slowDeque) PopBack() int {
	d.tail--
	return d.arr[d.tail+1]
}
