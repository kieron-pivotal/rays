package tuple_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTuple(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tuple Suite")
}
