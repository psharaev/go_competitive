package stress

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestStress(t *testing.T) {
	r := rand.New(rand.NewSource(42))
	for {
		test(t, r.Int63())
	}
	//test(t, 50)
}

func test(t *testing.T, seed int64) {
	rnd := rand.New(rand.NewSource(seed))

	arr := genArray(rnd, 1, 100, 2, 50)
	fmt.Println(len(arr))
	fmt.Println(sliceToString(arr))
}

func genArray(r *rand.Rand, minSize, maxSize, minValue, maxValueInc int) []int {
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
