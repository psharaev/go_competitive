package slice

import (
	"fmt"
	"slices"
	"strings"
)

type Matrix[T any] struct {
	M    [][]T
	Rows int
	Cols int
}

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

func FilledMatrix[T any](rows, cols int, val T) Matrix[T] {
	res := make([][]T, rows)
	for i := range res {
		res[i] = FilledSlice(cols, val)
	}
	return Matrix[T]{
		M:    res,
		Rows: rows,
		Cols: cols,
	}
}

func NewMatrix[T any](rows, cols int) Matrix[T] {
	res := make([][]T, rows)
	for i := range res {
		res[i] = make([]T, cols)
	}
	return Matrix[T]{
		M:    res,
		Rows: rows,
		Cols: cols,
	}
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

func RemoveItem[T any](arr []T, idx int) []T {
	if idx == 0 {
		return arr[1:]
	}
	if idx == len(arr)-1 {
		return arr[:idx]
	}

	return slices.Concat(arr[:idx], arr[idx+1:])
}

func (m *Matrix[T]) Join(sep string) string {
	sb := strings.Builder{}
	for _, row := range m.M {
		sb.WriteString(JoinSlice(row, sep, "", "\n"))
	}
	return sb.String()
}

func JoinSlice[T any](arr []T, sep, prefix, suffix string) string {
	if len(arr) == 0 {
		return prefix + suffix
	}
	sb := strings.Builder{}
	sb.WriteString(prefix)
	sb.WriteString(fmt.Sprintf("%v", arr))
	for _, item := range arr[1:] {
		sb.WriteString(sep)
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	sb.WriteString(suffix)
	return sb.String()
}

func DumpSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
