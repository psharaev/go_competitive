package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := nextInt(in)
	fmt.Fprintf(out, "%d\n", n)
}

func nextWord(r *bufio.Reader) string {
	var s string
	fscan, err := fmt.Fscan(r, &s)
	if err != nil {
		panic(err)
	}
	if fscan != 1 {
		panic("not found string")
	}
	return s
}

func nextWordChecked(r *bufio.Reader) (string, bool) {
	var s string
	n, err := fmt.Fscan(r, &s)
	if n == 0 || err != nil {
		return "", false
	}
	return s, true
}

func nextInt(in *bufio.Reader) int {
	var t int
	fscan, err := fmt.Fscan(in, &t)
	if err != nil {
		panic(err)
	}
	if fscan != 1 {
		panic("not found int")
	}
	return t
}

func nextArrayInt(in *bufio.Reader, n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = nextInt(in)
	}
	return res
}

func nextMatrixInt(in *bufio.Reader, rows int, cols int) [][]int {
	res := make([][]int, rows)
	for row := range rows {
		res[row] = nextArrayInt(in, cols)
	}
	return res
}

func nextFloat64(in *bufio.Reader) float64 {
	var t float64
	fscan, err := fmt.Fscan(in, &t)
	if err != nil {
		panic(err)
	}
	if fscan != 1 {
		panic("not found int")
	}
	return t
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

	res := make([]T, 0, len(arr)-1)
	res = append(res, arr[:idx]...)
	res = append(res, arr[idx+1:]...)
	return res
}
