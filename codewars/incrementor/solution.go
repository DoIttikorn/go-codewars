package incrementor

func incrementor2(n []int) []int {
	result := make([]int, len(n))
	for i, number := range n {
		result[i] = (i + 1 + number) % 10
	}
	return result
}
