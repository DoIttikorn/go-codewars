package multipleofindex_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMultipleOfIndex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsPalindrome Suite")
}
