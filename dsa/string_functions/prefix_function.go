package string_functions

func PrefixFunctionString(s string) []int {
	runes := []rune(s)
	arr := make([]int, len(runes))
	for idx, r := range runes {
		arr[idx] = int(r)
	}
	return PrefixFunction(arr)
}

func PrefixFunction[T comparable](arr []T) []int {
	res := make([]int, len(arr)+1)
	res[0] = -1

	for i := 1; i <= len(arr); i++ {
		k := res[i-1]
		for k != -1 && arr[k] != arr[i-1] {
			k = res[k]
		}
		res[i] = k + 1
	}

	return res
}
