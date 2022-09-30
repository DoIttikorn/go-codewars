package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIsPalindrome(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "IsPalindrome Suite")
}
