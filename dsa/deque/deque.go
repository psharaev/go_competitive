package deque

import (
	"fmt"
	"strings"
)

type Deque[T any] struct {
	data []T
	head int
	tail int
}

func NewDeque[T any](cap int) Deque[T] {
	cap = nextPowerOfTwo(cap)
	return Deque[T]{
		data: make([]T, cap),
		head: 0,
		tail: 0,
	}
}

func nextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	return n
}

func (d *Deque[T]) Size() int {
	mask := len(d.data) - 1
	return (d.tail - d.head + len(d.data)) & mask
}

func (d *Deque[T]) Empty() bool {
	return d.head == d.tail
}

func (d *Deque[T]) Front() T {
	if d.Empty() {
		panic("deque is empty")
	}
	return d.data[d.head]
}

func (d *Deque[T]) Back() T {
	if d.Empty() {
		panic("deque is empty")
	}
	mask := len(d.data) - 1
	return d.data[(d.tail-1)&mask]
}

func (d *Deque[T]) resize() {
	oldSize := len(d.data)
	newSize := oldSize << 1

	newData := make([]T, newSize)
	size := d.Size()

	if d.head < d.tail {
		copy(newData, d.data[d.head:d.tail])
	} else if size > 0 {
		n := copy(newData, d.data[d.head:])
		copy(newData[n:], d.data[:d.tail])
	}

	d.data = newData
	d.head = 0
	d.tail = size
}

func (d *Deque[T]) PushBack(item T) {
	mask := len(d.data) - 1
	if ((d.tail + 1) & mask) == d.head {
		d.resize()
		mask = len(d.data) - 1
	}
	d.data[d.tail] = item
	d.tail = (d.tail + 1) & mask
}

func (d *Deque[T]) PushFront(item T) {
	mask := len(d.data) - 1
	if ((d.tail + 1) & mask) == d.head {
		d.resize()
		mask = len(d.data) - 1
	}
	d.head = (d.head - 1) & mask
	d.data[d.head] = item
}

func (d *Deque[T]) PopFront() T {
	if d.Empty() {
		panic("deque is empty")
	}
	mask := len(d.data) - 1
	item := d.data[d.head]
	d.head = (d.head + 1) & mask
	return item
}

func (d *Deque[T]) PopBack() T {
	if d.Empty() {
		panic("deque is empty")
	}
	mask := len(d.data) - 1
	d.tail = (d.tail - 1) & mask
	return d.data[d.tail]
}

func (d *Deque[T]) String() string {
	if d.Empty() {
		return "[]"
	}

	var sb strings.Builder
	size := d.Size()
	mask := len(d.data) - 1
	fmt.Fprintf(&sb, "[%v", d.data[d.head&mask])
	for i := 1; i < size; i++ {
		fmt.Fprintf(&sb, " %v", d.data[(d.head+i)&mask])
	}
	sb.WriteString("]")
	return sb.String()
}
