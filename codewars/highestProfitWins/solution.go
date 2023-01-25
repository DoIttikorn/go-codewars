package highestprofitwins

import "sort"

func MinMax2(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

func MinMax3(a []int) [2]int {
	min, max := a[0], a[0]
	for _, v := range a[1:] {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return [2]int{min, max}
}
