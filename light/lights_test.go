package light_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/tuple"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lights", func() {

	Context("point light", func() {
		It("has position and intensity", func() {
			intensity := color.New(1, 1, 1)
			pos := tuple.Point(0, 0, 0)
			l := light.NewPoint(pos, intensity)
			Expect(l.Position).To(tuple.Equal(pos))
			Expect(l.Intensity).To(color.Equal(intensity))
		})
	})

})
