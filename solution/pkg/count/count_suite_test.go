package count_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCount(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Count Suite")
}
