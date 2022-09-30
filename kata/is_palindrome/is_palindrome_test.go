package kata_test

import (
	. "github.com/Doittikorn/golang-codewars/kata/is_palindrome"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("IsPalindrome", func() {
	It("tests basic strings", func() {
		Expect(IsPalindrome("a")).To(Equal(true))
		Expect(IsPalindrome("aba")).To(Equal(true))
		Expect(IsPalindrome("Abba")).To(Equal(true))
		Expect(IsPalindrome("hello")).To(Equal(false))
	})
})
