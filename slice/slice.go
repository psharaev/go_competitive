package slice

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
