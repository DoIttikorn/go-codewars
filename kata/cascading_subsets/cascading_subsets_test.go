package kata_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/Doittikorn/golang-codewars/kata/cascading_subsets"
)

var _ = Describe("CascadingSubsets", func() {
	It("Sample tests", func() {
		arr := []int{3, 5, 8, 13}
		Expect(CascadingSubsets(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
		Expect(CascadingSubsets(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
		Expect(CascadingSubsets(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
		Expect(CascadingSubsets([]int{}, 3)).To(BeEmpty())
	})
})
