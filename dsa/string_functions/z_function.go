package string_functions

func ZFunctionString(s string) []int {
	runes := []rune(s)
	arr := make([]int, len(runes))
	for idx, r := range runes {
		arr[idx] = int(r)
	}
	return ZFunction(arr)
}

func ZFunction[T comparable](arr []T) []int {
	if len(arr) == 0 {
		return []int{}
	}

	z := make([]int, len(arr))
	z[0] = len(arr)
	n := len(arr)
	l, r := 0, 0

	for i := 1; i < len(arr); i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}

		for i+z[i] < n && arr[z[i]] == arr[i+z[i]] {
			z[i]++
		}

		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}
