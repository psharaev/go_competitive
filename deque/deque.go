package deque

type Deque struct {
	arr  []int
	head int
	size int
}

func NewDeque(cap int) *Deque {
	return &Deque{
		arr:  make([]int, max(2, cap)),
		head: 0,
		size: 0,
	}
}

func (d *Deque) Size() int {
	return d.size
}

func (d *Deque) Empty() bool {
	return d.size == 0
}

func (d *Deque) resize(newCap int) {
	newArr := make([]int, newCap)
	if d.head+d.size <= len(d.arr) {
		copy(newArr, d.arr[d.head:(d.head+d.size)])
	} else {
		n := copy(newArr, d.arr[d.head:])
		copy(newArr[n:], d.arr[:(d.size-n)])
	}
	d.head = 0
	d.arr = newArr
}

func (d *Deque) PushBask(item int) {
	if len(d.arr) == d.size {
		d.resize(2 * d.size)
	}
	d.arr[d.tailInsert()] = item
	d.size++
}

func (d *Deque) PushFront(item int) {
	if len(d.arr) == d.size {
		d.resize(2 * d.size)
	}
	d.head = d.frontInsert()
	d.arr[d.head] = item
	d.size++
}

func (d *Deque) tailInsert() int {
	return (d.head + d.size) % len(d.arr)
}

func (d *Deque) frontInsert() int {
	if d.head-1 >= 0 {
		return d.head - 1
	}
	return len(d.arr) - 1
}

func (d *Deque) PopFront() int {
	if d.Empty() {
		panic("PopFront: deque is empty")
	}
	res := d.arr[d.head]
	d.head++
	d.size--
	if d.head == len(d.arr) {
		d.head = 0
	}
	return res
}

func (d *Deque) PopBack() int {
	if d.Empty() {
		panic("PopBack: deque is empty")
	}
	d.size--
	return d.arr[d.tailInsert()]
}
