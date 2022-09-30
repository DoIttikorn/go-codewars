package kata

import "fmt"

func CascadingSubsets(arr []int, n int) [][]int {
	var result [][]int
	var loop int = len(arr) % n
	for i := 0; i <= loop; i++ {
		fmt.Println(i)
		// var temp int
		// if i > len(arr) {
		// 	temp = len(arr)
		// } else {
		// 	temp = i + n
		// }
		result = append(result, arr[i:i+n])
		fmt.Println(result)
	}
	return result
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
