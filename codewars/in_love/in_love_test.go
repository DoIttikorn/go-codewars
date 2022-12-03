package convertstringnumber_test

import (
	. "github.com/Doittikorn/golang-codewars/codewars/in_love"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoveFunc", func() {
	It("Basic Tests", func() {
		Expect(LoveFunc(1, 4)).To(Equal(true))
		Expect(LoveFunc(2, 2)).To(Equal(false))
		Expect(LoveFunc(0, 1)).To(Equal(true))
		Expect(LoveFunc(0, 0)).To(Equal(false))
	})
})
