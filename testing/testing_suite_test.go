package testing_test

import (
	testing2 "github.com/Doittikorn/golang-codewars/testing"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestTesting(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Testing Suite")
}

var _ = Describe("Person.IsChild()", func() {
	XContext("when the person is a child", func() {
		It("return true", func() {
			person := testing2.Person{Age: 10}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})
	})

	XContext("when the person is a man", func() {
		It("return true", func() {
			person := testing2.Person{Age: 20}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})

	XContext("when the person is not a child", func() {

		BeforeEach(func() {
			log.Printf("not a child")
		})

		AfterEach(func() {
			log.Printf("not a child")
		})
		// Pending context XIt
		// skip context XIt
		It("return true", func() {
			person := testing2.Person{Age: 20}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})

	DescribeTable("IsChild table test",
		func(age int, expectedResponse bool) {
			person := testing2.Person{Age: age}

			response := person.IsChild()

			Expect(response).To(Equal(expectedResponse))
		},
		Entry("when the person is a child", 10, true),
		Entry("when the person is not a child", 20, false),
	)
})
