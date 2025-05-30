package sort

func MergeSort(nums []int) {
	buf := make([]int, len(nums))
	mergeSort(buf, nums)
}

func mergeSort(buf, arr []int) {
	if len(arr) <= 1 {
		return
	}

	al := arr[:len(arr)/2]
	ar := arr[len(arr)/2:]
	MergeSort(al)
	MergeSort(ar)

	merge(buf, al, ar)
	copy(arr, buf[:len(arr)])
}

func merge(buf, al, ar []int) {
	idxL := 0
	idxR := 0
	k := 0
	for idxL < len(al) && idxR < len(ar) {
		if al[idxL] < ar[idxR] {
			buf[k] = al[idxL]
			idxL++
		} else {
			buf[k] = ar[idxR]
			idxR++
		}
		k++
	}

	for idxL < len(al) {
		buf[k] = al[idxL]
		idxL++
		k++
	}

	for idxR < len(ar) {
		buf[k] = ar[idxR]
		idxR++
		k++
	}
}
