package geometry_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGeometry(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Geometry Suite")
}
