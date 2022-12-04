package multipleofindex

func MultipleOfIndex(ints []int) []int {
	// var result []int
	result := make([]int, 0)
	for i := 1; i < len(ints); i++ {
		if ints[i]%i == 0 {
			result = append(result, ints[i])
		}
	}
	return result
}
