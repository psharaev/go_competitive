package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	in := NewFastReader()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	countTests := in.NextInt()
	for range countTests {
		n := in.NextInt()
		arr := in.NextSliceInt(n)

		fmt.Fprintln(out, n)
		fmt.Fprintln(out, arr)
	}
}

type FastReader struct {
	Reader bufio.Reader
}

func NewFastReader() FastReader {
	return FastReader{Reader: *bufio.NewReader(os.Stdin)}
}

func (r *FastReader) NextWord() string {
	sb := strings.Builder{}

	foundWord := false
	for {
		ch, _, err := r.Reader.ReadRune()
		if err != nil {
			panic(err)
		}

		if unicode.IsSpace(ch) {
			if foundWord {
				return sb.String()
			}
			continue
		}

		foundWord = true
		sb.WriteRune(ch)
	}
}

func (r *FastReader) NextWordChecked() (string, bool) {
	sb := strings.Builder{}

	foundWord := false
	for {
		ch, _, err := r.Reader.ReadRune()
		if err != nil {
			return "", false
		}

		if unicode.IsSpace(ch) {
			if foundWord {
				return sb.String(), true
			}
			continue
		}

		foundWord = true
		sb.WriteRune(ch)
	}
}

func (r *FastReader) NextInt() int {
	num, err := strconv.Atoi(r.NextWord())
	if err != nil {
		panic(err)
	}
	return num
}

func (r *FastReader) NextLine() string {
	str, err := r.Reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return strings.TrimRight(str, "\r\n")
}

func (r *FastReader) NextSliceInt(size int) []int {
	res := make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = r.NextInt()
	}
	return res
}

func (r *FastReader) NextSliceChars() []byte {
	return []byte(r.NextWord())
}

func (r *FastReader) NextSliceWord(size int) []string {
	res := make([]string, size)
	for i := 0; i < size; i++ {
		res[i] = r.NextWord()
	}
	return res
}

func (r *FastReader) NextMatrixInt(rows int, cols int) Matrix[int] {
	res := make([][]int, rows)
	for row := range rows {
		res[row] = r.NextSliceInt(cols)
	}
	return Matrix[int]{
		M:    res,
		Rows: rows,
		Cols: cols,
	}
}

func (r *FastReader) NextMatrixChars(rows int, cols int) Matrix[byte] {
	res := make([][]byte, rows)
	for row := range rows {
		res[row] = r.NextSliceChars()
		if len(res[row]) != cols {
			panic("matrix rows out of range")
		}
	}
	return Matrix[byte]{
		M:    res,
		Rows: rows,
		Cols: cols,
	}
}

func (r *FastReader) NextFloat64() float64 {
	num, err := strconv.ParseFloat(r.NextWord(), 64)
	if err != nil {
		panic(err)
	}
	return num
}

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

func (d *Deque[T]) PopBack() T {
	if d.Empty() {
		panic("PopBack: deque is empty")
	}
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

func DumpSlice[T any](arr []T) {
	for _, item := range arr {
		fmt.Printf("%v ", item)
	}
	fmt.Println()
}
