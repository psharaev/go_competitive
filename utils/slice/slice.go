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

func New2dSlice[T any](rows, cols int) [][]T {
	res := make([][]T, rows)
	for i := range res {
		res[i] = make([]T, cols)
	}
	return res
}

func NewMatrix[T any](rows, cols int) Matrix[T] {
	return Matrix[T]{M: New2dSlice[T](rows, cols), Rows: rows, Cols: cols}
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

func (m *Matrix[T]) Clone() Matrix[T] {
	res := New2dSlice[T](m.Rows, m.Cols)
	for row := range m.Rows {
		for col := range m.Cols {
			res[row][col] = m.M[row][col]
		}
	}
	return Matrix[T]{M: res, Rows: m.Rows, Cols: m.Cols}
}

func (m *Matrix[T]) Transpose() Matrix[T] {
	res := New2dSlice[T](m.Cols, m.Rows)

	for row := 0; row < m.Rows; row++ {
		for col := 0; col < m.Cols; col++ {
			res[col][row] = m.M[row][col]
		}
	}

	return Matrix[T]{M: res, Rows: m.Cols, Cols: m.Rows}
}

func (m *Matrix[T]) RotateClockwise(count int) Matrix[T] {
	// Нормализация count в диапазон [0, 3]
	count = ((count % 4) + 4) % 4

	switch count % 4 {
	case 0:
		return m.Clone()
	case 1:
		return m.RotateClockwise90()
	case 2:
		return m.RotateClockwise180()
	case 3:
		return m.RotateClockwise270()
	default:
		panic("unreachable")
	}
}

// RotateClockwise90 Поворот на 90 градусов по часовой стрелке
func (m *Matrix[T]) RotateClockwise90() Matrix[T] {
	res := New2dSlice[T](m.Cols, m.Rows)
	for i := 0; i < m.Cols; i++ {
		for j := 0; j < m.Rows; j++ {
			res[i][j] = m.M[m.Rows-1-j][i]
		}
	}
	return Matrix[T]{M: res, Rows: m.Cols, Cols: m.Rows}
}

// RotateClockwise180 Поворот на 180 градусов по часовой стрелке
func (m *Matrix[T]) RotateClockwise180() Matrix[T] {
	res := New2dSlice[T](m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			res[i][j] = m.M[m.Rows-1-i][m.Cols-1-j]
		}
	}
	return Matrix[T]{M: res, Rows: m.Rows, Cols: m.Cols}
}

// RotateClockwise270 Поворот на 270 градусов по часовой стрелке
func (m *Matrix[T]) RotateClockwise270() Matrix[T] {
	res := New2dSlice[T](m.Cols, m.Rows)
	for i := 0; i < m.Cols; i++ {
		for j := 0; j < m.Rows; j++ {
			res[i][j] = m.M[j][m.Cols-1-i]
		}
	}
	return Matrix[T]{M: res, Rows: m.Cols, Cols: m.Rows}
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
	sb.WriteString(fmt.Sprintf("%v", arr[0]))
	for _, item := range arr[1:] {
		sb.WriteString(sep)
		sb.WriteString(fmt.Sprintf("%v", item))
	}
	sb.WriteString(suffix)
	return sb.String()
}

func SliceCopy[T any](arr []T) []T {
	return append([]T(nil), arr...)
}

func DumpSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
