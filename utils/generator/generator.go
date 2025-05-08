package generator

import "math/rand"

type Generator struct {
	r rand.Rand
}

func NewGenerator(seed int) *Generator {
	return &Generator{
		r: *rand.New(rand.NewSource(int64(seed))),
	}
}

// Int return inclusive [min, max]
func (g *Generator) Int(min, max int) int {
	return g.r.Intn(max-min+1) + min
}

// IntExc return exclusive [min, max)
func (g *Generator) IntExc(min, max int) int {
	return g.r.Intn(max-min) + min
}

// Bool returns true with probability p (clamped to [0.0,1.0]), false otherwise.
func (g *Generator) Bool(p float64) bool {
	if p < 0 || p > 1 {
		panic("p must be between 0.0 and 1.0")
	}

	return g.r.Float64() < p
}

// SliceInt n = [minSize, maxSize] val = [minVal, maxVal]
func (g *Generator) SliceInt(minSize, maxSize, minVal, maxVal int) []int {
	return Slice[int](g, minSize, maxSize,
		func() int {
			return g.Int(minVal, maxVal)
		},
	)
}

// SliceBool n = [minSize, maxSize] val = gen.Bool(p)
func (g *Generator) SliceBool(minSize, maxSize int, p float64) []bool {
	return Slice[bool](g, minSize, maxSize,
		func() bool {
			return g.Bool(p)
		},
	)
}

// Slice n = [minSize, maxSize] val = valGen()
func Slice[T any](g *Generator, minSize, maxSize int, valGen func() T) []T {
	if minSize < 0 {
		panic("minSize must be >= 0")
	}
	if maxSize < minSize {
		panic("minSize must be <= maxSize")
	}

	size := g.Int(minSize, maxSize)

	if size == 0 {
		if g.Bool(0.5) {
			return nil
		}
		return []T{}
	}

	slice := make([]T, size)

	for i := 0; i < size; i++ {
		slice[i] = valGen()
	}

	return slice
}

// Segment
// l <= r
// l = [0, len(arr))
// r = [l, len(arr))
func Segment[T any](g *Generator, arr []T) (l int, r int) {
	l = g.IntExc(0, len(arr))
	r = g.IntExc(l, len(arr))
	return l, r
}

// Pos idx = [0, len(arr))
func Pos[T any](g *Generator, arr []T) (idx int) {
	return g.IntExc(0, len(arr))
}
