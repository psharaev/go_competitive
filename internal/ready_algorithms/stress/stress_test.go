package stress

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestStress(t *testing.T) {
	testCase(t, []int{1})

	for seed := range 1000 {
		t.Run(strconv.Itoa(seed), func(t *testing.T) {
			test(t, seed)
		})
	}
	//test(t, 50)
}

func test(t *testing.T, seed int) {
	rnd := rand.New(rand.NewSource(int64(seed)))

	arr := genSliceInt(rnd, 1, 100, 2, 50)
	testCase(t, arr)
}

func testCase(t *testing.T, arr []int) {
	fmt.Println(len(arr))
	fmt.Println(sliceToString(arr))

	require.Equal(t, len(arr), len(arr))
}

func genSliceInt(r *rand.Rand, minSize, maxSize, minValue, maxValueInc int) []int {
	n := genInt(r, minSize, maxSize)
	if n == 0 {
		if r.Intn(2) == 0 {
			return nil
		}
		return []int{}
	}
	a := make([]int, n)
	for i := range a {
		a[i] = genInt(r, minValue, maxValueInc)
	}
	return a
}

func genInt(r *rand.Rand, min, maxInc int) int {
	return r.Intn(maxInc-min+1) + min
}

func sliceToString[T any](slice []T) string {
	if len(slice) == 0 {
		return ""
	}
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%v", slice[0]))

	for _, item := range slice[1:] {
		builder.WriteString(fmt.Sprintf(" %v", item))
	}

	return builder.String()
}
