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

// IntInc return inclusive [min, max]
func (g *Generator) IntInc(min, max int) int {
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
	if minSize < 0 {
		panic("minSize must be >= 0")
	}
	if maxSize < minSize {
		panic("minSize must be <= maxSize")
	}

	n := g.IntInc(minSize, maxSize)
	return GenSlice[int](g, n,
		func() int {
			return g.IntInc(minVal, maxVal)
		},
	)
}

func GenSlice[T any](g *Generator, size int, valGen func() T) []T {
	if size < 0 {
		panic("size must be >= 0")
	}

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
