package material_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMaterial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Material Suite")
}
