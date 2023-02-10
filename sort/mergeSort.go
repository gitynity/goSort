package sort

import (
	"sync"
)

func Merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func MergeSortMulti(s []int) []int {
	if len(s) <= 1 {
		return s
	}

	n := len(s) / 2

	wg := sync.WaitGroup{}
	wg.Add(2)

	var l []int
	var r []int

	go func() {
		l = MergeSortMulti(s[:n])
		wg.Done()
	}()

	go func() {
		r = MergeSortMulti(s[n:])
		wg.Done()
	}()

	wg.Wait()
	return Merge(l, r)
}
