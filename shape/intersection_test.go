package shape_test

import (
	"math"

	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intersection", func() {

	var (
		s  *shape.Object
		ix *shape.Intersections
	)

	BeforeEach(func() {
		s = shape.NewSphere()
		ix = shape.NewIntersections()
	})

	It("gets lower of 1 and 2", func() {
		ix.Add(1, s)
		ix.Add(2, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 1))
	})

	It("gets higher of 1 and -1", func() {
		ix.Add(-1, s)
		ix.Add(1, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 1))
	})

	It("returns nil for no hit", func() {
		ix.Add(-1, s)
		ix.Add(-2, s)
		i := ix.Hit()
		Expect(i).To(BeNil())
	})

	It("returns nil for no hit", func() {
		ix.Add(5, s)
		ix.Add(7, s)
		ix.Add(-3, s)
		ix.Add(2, s)
		i := ix.Hit()
		Expect(i.T).To(BeNumerically("~", 2))
	})

	Context("PrepareComputations", func() {
		It("can prepare common details", func() {
			o := shape.NewSphere()
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			ix := shape.NewIntersections()
			ix.Add(4, o)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			Expect(comps.T).To(BeNumerically("~", 4))
			Expect(comps.Object).To(Equal(o))
			Expect(comps.Point).To(tuple.Equal(tuple.Point(0, 0, -1)))
			Expect(comps.EyeV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
			Expect(comps.NormalV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
		})

		It("can identify an outside hit", func() {
			o := shape.NewSphere()
			ix := shape.NewIntersections()
			ix.Add(4, o)
			i := ix.Get(0)
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			comps := i.PrepareComputations(r, ix)
			Expect(comps.Inside).To(BeFalse())
		})

		It("can identify an inside hit", func() {
			o := shape.NewSphere()
			ix := shape.NewIntersections()
			ix.Add(1, o)
			i := ix.Get(0)
			r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
			comps := i.PrepareComputations(r, ix)
			Expect(comps.Point).To(tuple.Equal(tuple.Point(0, 0, 1)))
			Expect(comps.EyeV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
			Expect(comps.Inside).To(BeTrue())
			Expect(comps.NormalV).To(tuple.Equal(tuple.Vector(0, 0, -1)))
		})
	})

	It("can calculate the over point", func() {
		r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
		s := shape.NewSphere()
		s.SetTransform(matrix.Translation(0, 0, 1))
		ix := shape.NewIntersections()
		ix.Add(5, s)
		i := ix.Get(0)
		comps := i.PrepareComputations(r, ix)
		Expect(comps.OverPoint.Z).To(BeNumerically("<", -tuple.EPSILON/2))
		Expect(comps.Point.Z).To(BeNumerically(">", comps.OverPoint.Z))
	})

	It("can calculate the reflective vector", func() {
		r2 := math.Sqrt(2)
		r := ray.New(tuple.Point(0, 1, -1), tuple.Vector(0, -r2/2, r2/2))
		s := shape.NewPlane()
		ix := shape.NewIntersections()
		ix.Add(r2, s)
		i := ix.Get(0)
		comps := i.PrepareComputations(r, ix)
		Expect(comps.ReflectV).To(tuple.Equal(tuple.Vector(0, r2/2, r2/2)))
	})

	Context("refractive index each side of an intersection", func() {
		It("can get n1 and n2 correct", func() {
			a := shape.NewGlassSphere()
			a.SetTransform(matrix.Scaling(2, 2, 2))
			ma := material.New()
			ma.RefractiveIndex = 1.5
			a.SetMaterial(ma)

			b := shape.NewGlassSphere()
			b.SetTransform(matrix.Translation(0, 0, -0.25))
			mb := material.New()
			mb.RefractiveIndex = 2.0
			b.SetMaterial(mb)

			c := shape.NewGlassSphere()
			c.SetTransform(matrix.Translation(0, 0, 0.25))
			mc := material.New()
			mc.RefractiveIndex = 2.5
			c.SetMaterial(mc)

			r := ray.New(tuple.Point(0, 0, -4), tuple.Vector(0, 0, 1))
			xs := shape.NewIntersections()
			xs.Add(2, a)
			xs.Add(2.75, b)
			xs.Add(3.25, c)
			xs.Add(4.75, b)
			xs.Add(5.25, c)
			xs.Add(6, a)

			expectations := [][2]float64{
				{1.0, 1.5},
				{1.5, 2.0},
				{2.0, 2.5},
				{2.5, 2.5},
				{2.5, 1.5},
				{1.5, 1.0},
			}

			for i, expt := range expectations {
				comps := xs.Get(i).PrepareComputations(r, xs)
				Expect(comps.N1).To(Equal(expt[0]))
				Expect(comps.N2).To(Equal(expt[1]))
			}

		})
	})

	It("can calculate the under_point", func() {
		r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
		s := shape.NewGlassSphere()
		s.SetTransform(matrix.Translation(0, 0, 1))
		xs := shape.NewIntersections()
		xs.Add(5, s)
		i := xs.Get(0)
		comps := i.PrepareComputations(r, xs)
		Expect(comps.UnderPoint.Z).To(BeNumerically(">", tuple.EPSILON/2))
		Expect(comps.Point.Z).To(BeNumerically("<", comps.UnderPoint.Z))
	})
})
