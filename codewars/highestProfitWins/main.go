package highestprofitwins

func MinMax(arr []int) [2]int {
	var min, max int
	for i, v := range arr {
		if i == 0 {
			min = v
			max = v
		}
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return [2]int{min, max}
}

// https://www.codewars.com/kata/559590633066759614000063/train/go
