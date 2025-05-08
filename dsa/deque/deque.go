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
	if cap < 1 {
		cap = 1
	}
	return Deque[T]{
		data: make([]T, cap),
		head: 0,
		tail: 0,
	}
}

func (d *Deque[T]) Size() int {
	return (d.tail - d.head + len(d.data)) % len(d.data)
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
	lastPos := (d.tail - 1 + len(d.data)) % len(d.data)
	return d.data[lastPos]
}

func (d *Deque[T]) resize() {
	oldSize := len(d.data)
	newSize := oldSize * 2
	if newSize == 0 {
		newSize = 1
	}
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
	if (d.tail+1)%len(d.data) == d.head {
		d.resize()
	}
	d.data[d.tail] = item
	d.tail = (d.tail + 1) % len(d.data)
}

func (d *Deque[T]) PushFront(item T) {
	if (d.tail+1)%len(d.data) == d.head {
		d.resize()
	}
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = item
}

func (d *Deque[T]) PopFront() T {
	if d.Empty() {
		panic("deque is empty")
	}
	item := d.data[d.head]
	d.head = (d.head + 1) % len(d.data)
	return item
}

func (d *Deque[T]) PopBack() T {
	if d.Empty() {
		panic("deque is empty")
	}
	d.tail = (d.tail - 1 + len(d.data)) % len(d.data)
	return d.data[d.tail]
}

func (d *Deque[T]) String() string {
	if d.Empty() {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")
	size := d.Size()
	for i := 0; i < size; i++ {
		pos := (d.head + i) % len(d.data)
		if i > 0 {
			sb.WriteString(" ")
		}
		fmt.Fprintf(&sb, "%v", d.data[pos])
	}
	sb.WriteString("]")
	return sb.String()
}
