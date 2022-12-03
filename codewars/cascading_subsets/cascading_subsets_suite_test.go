package kata_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCascadingSubsets(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CascadingSubsets Suite")
}
