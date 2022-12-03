package convertstringnumber_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStringToNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringToNumber Suite")
}
