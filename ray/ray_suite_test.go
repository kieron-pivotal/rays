package ray_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRay(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ray Suite")
}
