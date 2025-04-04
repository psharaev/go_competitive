package fastio

import (
	"bufio"
	"fmt"
)

//func main() {
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	defer out.Flush()
//
//	n := nextInt(in)
//	arr := nextArrayInt(in, n)
//
//	fmt.Fprintln(out, arr)
//}

func next(r *bufio.Reader) string {
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
