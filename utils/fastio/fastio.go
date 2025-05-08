package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/psharaev/go_competitive/utils/slice"
)

func main() {
	in := NewFastReader()
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	n := in.NextInt()
	arr := in.NextLine()

	fmt.Fprintln(out, n)
	fmt.Fprintln(out, arr)
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

func (r *FastReader) NextMatrixInt(rows int, cols int) slice.Matrix[int] {
	res := make([][]int, rows)
	for row := range rows {
		res[row] = r.NextSliceInt(cols)
	}
	return slice.Matrix[int]{
		M:    res,
		Rows: rows,
		Cols: cols,
	}
}

func (r *FastReader) NextMatrixChars(rows int, cols int) slice.Matrix[byte] {
	res := make([][]byte, rows)
	for row := range rows {
		res[row] = r.NextSliceChars()
		if len(res[row]) != cols {
			panic("matrix rows out of range")
		}
	}
	return slice.Matrix[byte]{
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
