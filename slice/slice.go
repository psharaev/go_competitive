package slice

func FillSlice[T any](arr []T, val T) {
	if len(arr) == 0 {
		return
	}
	arr[0] = val

	for i := 1; i < len(arr); i *= 2 {
		copy(arr[i:], arr[:i])
	}
}

func FilledSlice[T any](n int, val T) []T {
	res := make([]T, n)

	if n == 0 {
		return res
	}
	res[0] = val

	for i := 1; i < len(res); i *= 2 {
		copy(res[i:], res[:i])
	}
	return res
}

func FilledMatrix[T any](rows, cols int, val T) [][]T {
	res := make([][]T, rows)
	for i := range res {
		res[i] = FilledSlice(cols, val)
	}
	return res
}

func NewMatrix[T any](rows, cols int) [][]T {
	res := make([][]T, rows)
	for i := range res {
		res[i] = make([]T, cols)
	}
	return res
}

func Last[T any](arr []T) T {
	return arr[len(arr)-1]
}

func Sum(arr []int) int {
	res := 0
	for _, v := range arr {
		res += v
	}
	return res
}

func SumMatrix(arr [][]int) int {
	res := 0
	for _, v := range arr {
		res += Sum(v)
	}
	return res
}
