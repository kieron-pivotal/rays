package pattern_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/pattern/patternfakes"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	black = color.New(0, 0, 0)
	white = color.New(1, 1, 1)
)

var _ = Describe("Pattern", func() {

	Context("patterns", func() {
		var (
			fakePattern   *patternfakes.FakeActualPattern
			fakeInvGetter *patternfakes.FakeInvTransformGetter
			p             pattern.Pattern
		)

		BeforeEach(func() {
			fakePattern = new(patternfakes.FakeActualPattern)
			fakeInvGetter = new(patternfakes.FakeInvTransformGetter)
			p = pattern.New(fakePattern)
		})

		It("a new pattern has the identity transformation", func() {
			Expect(p.GetTransform()).To(matrix.Equal(matrix.Identity(4, 4)))
		})

		It("can be assigned a transformation", func() {
			t := matrix.Scaling(1, 2, 3)
			p.SetTransform(t)
			Expect(p.GetTransform()).To(matrix.Equal(t))
		})

		It("transforms the point when the object has a transformation", func() {
			t := matrix.Scaling(2, 2, 2).Inverse()
			fakeInvGetter.GetInverseTransformReturns(t)
			p.PatternAtShape(fakeInvGetter, tuple.Point(2, 3, 4))
			Expect(fakePattern.PatternAtCallCount()).To(Equal(1))
			op := fakePattern.PatternAtArgsForCall(0)
			Expect(op).To(tuple.Equal(tuple.Point(1, 1.5, 2)))
		})

		It("transforms the point when the pattern has a transformation", func() {
			t := matrix.Scaling(2, 2, 2)
			p.SetTransform(t)
			fakeInvGetter.GetInverseTransformReturns(matrix.Identity(4, 4))
			p.PatternAtShape(fakeInvGetter, tuple.Point(2, 3, 4))
			Expect(fakePattern.PatternAtCallCount()).To(Equal(1))
			op := fakePattern.PatternAtArgsForCall(0)
			Expect(op).To(tuple.Equal(tuple.Point(1, 1.5, 2)))
		})

		It("transforms the point when both obj and pattern have transforms", func() {
			t := matrix.Scaling(2, 2, 2).Inverse()
			fakeInvGetter.GetInverseTransformReturns(t)
			s := matrix.Translation(0.5, 1, 1.5)
			p.SetTransform(s)
			p.PatternAtShape(fakeInvGetter, tuple.Point(3, 4, 5))
			Expect(fakePattern.PatternAtCallCount()).To(Equal(1))
			op := fakePattern.PatternAtArgsForCall(0)
			Expect(op).To(tuple.Equal(tuple.Point(1, 1, 1)))
		})
	})
})
