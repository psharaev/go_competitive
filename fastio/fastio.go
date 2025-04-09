package fastio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := nextInt(in)
	arr := nextArrayInt(in, n)

	fmt.Fprintln(out, arr)
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

func dumpArrayInt(out io.Writer, arr []int) {
	for _, item := range arr {
		fmt.Fprintf(out, "%d ", item)
	}
}
