package bitset

import "math/bits"

type Bitset struct {
	data []uint64
	size int
}

func NewBitset(size int) Bitset {
	return Bitset{data: make([]uint64, (size+63)>>6), size: size}
}

func (b *Bitset) Set(index int) {
	b.data[index>>6] |= 1 << (index & 63)
}

func (b *Bitset) Clear(index int) {
	b.data[index>>6] &^= 1 << (index & 63)
}

func (b *Bitset) Get(index int) bool {
	return (b.data[index>>6] & (1 << (index & 63))) != 0
}

// Toggle инвертирует бит в позиции index
func (b *Bitset) Toggle(index int) {
	b.data[index>>6] ^= 1 << (index & 63)
}

// Count возвращает количество установленных битов
func (b *Bitset) Count() int {
	count := 0
	for _, word := range b.data {
		count += bits.OnesCount64(word)
	}
	return count
}

func (b *Bitset) Intersect(other *Bitset) {
	minLen := min(len(b.data), len(other.data))
	for i := 0; i < minLen; i++ {
		b.data[i] &= other.data[i]
	}
	for i := minLen; i < len(b.data); i++ {
		b.data[i] = 0
	}
}

func (b *Bitset) Union(other *Bitset) {
	minLen := min(len(b.data), len(other.data))
	for i := 0; i < minLen; i++ {
		b.data[i] |= other.data[i]
	}
}

// SetAll устанавливает все биты в 1
func (b *Bitset) SetAll() {
	for i := range b.data {
		b.data[i] = ^uint64(0)
	}
	if remainder := b.size % 64; remainder != 0 {
		last := len(b.data) - 1
		b.data[last] &= (1 << remainder) - 1
	}
}

// ClearAll устанавливает все биты в 0
func (b *Bitset) ClearAll() {
	for i := range b.data {
		b.data[i] = 0
	}
}

func (b *Bitset) SetVal(index int, val bool) {
	if val {
		b.Set(index)
	} else {
		b.Clear(index)
	}
}
