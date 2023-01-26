package incrementor

func Incrementor(nums []int) []int {
	var result = make([]int, len(nums))
	for i := range nums {

		result[i] = nums[i] + i + 1
		if result[i] >= 10 {
			result[i] = result[i] % 10
		}
	}
	return result
}
