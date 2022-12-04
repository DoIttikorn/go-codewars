package convertstringarray_test

import (
	. "github.com/Doittikorn/golang-codewars/codewars/convert_string_array"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tests", func() {
	It("Sample tests", func() {
		Expect(StringToArray("Robin Singh")).To(Equal([]string{"Robin", "Singh"}))
		Expect(StringToArray("CodeWars")).To(Equal([]string{"CodeWars"}))
		Expect(StringToArray("I love arrays they are my favorite")).To(Equal([]string{"I", "love", "arrays", "they", "are", "my", "favorite"}))
		Expect(StringToArray("1 2 3")).To(Equal([]string{"1", "2", "3"}))
	})
})
