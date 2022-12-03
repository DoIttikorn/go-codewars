package convertstringnumber_test

import (
	. "github.com/Doittikorn/golang-codewars/codewars/convert_string_number"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringToNumber", func() {
	It("tests basic strings", func() {
		Expect(StringToNumber("1234")).To(Equal(1234))
		Expect(StringToNumber("605")).To(Equal(605))
		Expect(StringToNumber("1405")).To(Equal(1405))
		Expect(StringToNumber("-7")).To(Equal(-7))
	})
})
