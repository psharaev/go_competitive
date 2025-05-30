package binary_search

import "slices"

func BinarySearchInsert(nums []int, target int) int {
	search, _ := slices.BinarySearch(nums, target)
	return search
}

func BinarySearchLeft(nums []int, target int) int {
	// min i: arr[i] >= target
	// arr[l] < target && arr[r] >= target
	l := -1
	r := len(nums)
	for l+1 < r {
		m := l + (r-l)/2
		if nums[m] >= target {
			r = m
		} else {
			l = m
		}
	}
	return r
}

func BinarySearchRight(nums []int, target int) int {
	// max i: arr[i] <= target
	// arr[l] <= target && arr[r] > target
	l := -1
	r := len(nums)
	for l+1 < r {
		m := l + (r-l)/2
		if nums[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	return l
}
