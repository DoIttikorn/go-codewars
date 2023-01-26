package incrementer

func Incrementer(nums []int) []int {
	for i := range nums {
		nums[i]++
	}
	return nums
}
