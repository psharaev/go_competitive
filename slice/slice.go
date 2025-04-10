package slice

import (
	"fmt"
	"slices"
)

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

func RemoveItem[T any](arr []T, idx int) []T {
	if idx == 0 {
		return arr[1:]
	}
	if idx == len(arr)-1 {
		return arr[:idx]
	}

	return slices.Concat(arr[:idx], arr[idx+1:])
}

func DumpSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}

func dumpMatrix[T any](arr [][]T) {
	for _, row := range arr {
		DumpSlice(row)
	}
}
