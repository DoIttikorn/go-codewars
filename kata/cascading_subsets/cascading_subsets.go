package kata

func CascadingSubsets(arr []int, n int) (result [][]int) {

	for index := range arr {
		if index+n > len(arr) {
			break
		}
		result = append(result, arr[index:index+n])

	}
	return
}

/*
var _ = Describe("CascadingSubsets", func() {
	It("Sample tests", func() {
		arr := []int{3, 5, 8, 13}
		Expect(CascadingSubsets(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
		Expect(CascadingSubsets(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
		Expect(CascadingSubsets(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
		Expect(CascadingSubsets([]int{}, 3)).To(BeEmpty())
	})
})
*/
