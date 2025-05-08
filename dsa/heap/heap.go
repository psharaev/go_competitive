package heap

type Heap[Item any] struct {
	arr []Item
	cmp func(a, b Item) int
}

func NewHeap[Item any](arr []Item, cmp func(a, b Item) int) Heap[Item] {
	h := Heap[Item]{
		arr: make([]Item, len(arr)),
		cmp: cmp,
	}
	copy(h.arr, arr)

	for i := (len(h.arr) - 2) / 2; i >= 0; i-- {
		h.down(i)
	}

	return h
}

func (h *Heap[Item]) Push(val Item) {
	h.arr = append(h.arr, val)
	h.up(len(h.arr) - 1)
}

func (h *Heap[Item]) Pop() Item {
	val := h.arr[0]
	last := len(h.arr) - 1
	h.arr[0] = h.arr[last]
	h.arr = h.arr[:last]

	if len(h.arr) > 0 {
		h.down(0)
	}

	return val
}

func (h *Heap[Item]) Peek() Item {
	return h.arr[0]
}

func (h *Heap[Item]) Size() int {
	return len(h.arr)
}

func (h *Heap[Item]) IsEmpty() bool {
	return len(h.arr) == 0
}

func (h *Heap[Item]) up(idx int) {
	for {
		p := (idx - 1) / 2
		if p == idx || h.cmp(h.arr[p], h.arr[idx]) <= 0 {
			break
		}
		h.arr[p], h.arr[idx] = h.arr[idx], h.arr[p]
		idx = p
	}
}

func (h *Heap[Item]) down(idx int) {
	last := len(h.arr) - 1
	for {
		l := 2*idx + 1
		r := 2*idx + 2
		swapIdx := idx

		if l <= last && h.cmp(h.arr[swapIdx], h.arr[l]) > 0 {
			swapIdx = l
		}
		if r <= last && h.cmp(h.arr[swapIdx], h.arr[r]) > 0 {
			swapIdx = r
		}

		if swapIdx == idx {
			return
		}

		h.arr[idx], h.arr[swapIdx] = h.arr[swapIdx], h.arr[idx]
		idx = swapIdx
	}
}
