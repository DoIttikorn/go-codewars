package kata_test

import (
	"math/rand"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/Doittikorn/golang-codewars/kata/cascading_subsets"
)

func referenceSolution(arr []int, n int) (o [][]int) {
	for i := 0; i < len(arr)-n+1; i++ {
		var a []int
		for j := i; j < i+n; j++ {
			a = append(a, arr[j])
		}
		o = append(o, a)
	}
	return
}

var _ = Describe("CascadingSubsets", func() {
	It("Sample tests", func() {
		arr := []int{3, 5, 8, 13}
		Expect(CascadingSubsets(arr, 1)).To(Equal([][]int{{3}, {5}, {8}, {13}}))
		Expect(CascadingSubsets(arr, 2)).To(Equal([][]int{{3, 5}, {5, 8}, {8, 13}}))
		Expect(CascadingSubsets(arr, 3)).To(Equal([][]int{{3, 5, 8}, {5, 8, 13}}))
		Expect(CascadingSubsets([]int{}, 3)).To(BeEmpty())
	})

	It("Random tests", func() {
		rand.Seed(time.Now().UTC().UnixNano())
		for i := 0; i < 100; i++ {
			l, n := rand.Intn(11), 1+rand.Intn(10)
			arr := make([]int, l)
			for j := 0; j < l; j++ {
				arr[j] = rand.Intn(21)
			}
			expected, actual := referenceSolution(arr, n), CascadingSubsets(arr, n)
			if len(expected) > 0 {
				Expect(actual).To(Equal(expected))
			} else {
				Expect(actual).To(BeEmpty())
			}
		}
	})
})
