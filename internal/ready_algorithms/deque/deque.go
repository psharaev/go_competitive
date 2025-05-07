package deque

import "fmt"

type Deque[T any] struct {
	arr  []T
	head int
	size int
}

func NewDeque[T any](cap int) *Deque[T] {
	return &Deque[T]{
		arr:  make([]T, max(2, cap)),
		head: 0,
		size: 0,
	}
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) Front() T {
	return d.arr[d.head]
}

func (d *Deque[T]) Back() T {
	idx := d.head + d.size - 1
	if idx >= len(d.arr) {
		idx -= len(d.arr)
	}
	return d.arr[idx]
}

func (d *Deque[T]) resize(newCap int) {
	newArr := make([]T, newCap)
	if d.head+d.size <= len(d.arr) {
		copy(newArr, d.arr[d.head:(d.head+d.size)])
	} else {
		n := copy(newArr, d.arr[d.head:])
		copy(newArr[n:], d.arr[:(d.size-n)])
	}
	d.head = 0
	d.arr = newArr
}

func (d *Deque[T]) PushBack(item T) {
	if len(d.arr) == d.size {
		d.resize(2 * d.size)
	}
	d.arr[d.tailInsert()] = item
	d.size++
}

func (d *Deque[T]) PushFront(item T) {
	if len(d.arr) == d.size {
		d.resize(2 * d.size)
	}
	d.head = d.frontInsert()
	d.arr[d.head] = item
	d.size++
}

func (d *Deque[T]) tailInsert() int {
	return (d.head + d.size) % len(d.arr)
}

func (d *Deque[T]) frontInsert() int {
	if d.head-1 >= 0 {
		return d.head - 1
	}
	return len(d.arr) - 1
}

func (d *Deque[T]) PopFront() T {
	res := d.arr[d.head]
	d.head++
	d.size--
	if d.head == len(d.arr) {
		d.head = 0
	}
	return res
}

func (d *Deque[T]) PopBack() T {
	d.size--
	return d.arr[d.tailInsert()]
}

func (d *Deque[T]) String() string {
	res := make([]T, d.Size())
	if d.head+d.size <= len(d.arr) {
		copy(res, d.arr[d.head:(d.head+d.size)])
	} else {
		n := copy(res, d.arr[d.head:])
		copy(res[n:], d.arr[:(d.size-n)])
	}

	return fmt.Sprintf("%v", res)
}
