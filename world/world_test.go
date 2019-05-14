package world_test

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/pattern/fakes"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("World", func() {
	It("can create a world", func() {
		w := world.New()
		Expect(w.Objects).To(BeEmpty())
		Expect(w.LightSource).To(BeNil())
	})

	It("can create the default world", func() {
		w := world.Default()
		lightSource := w.LightSource
		Expect(lightSource.Position).To(tuple.Equal(tuple.Point(-10, 10, -10)))
		Expect(lightSource.Intensity).To(color.Equal(color.New(1, 1, 1)))
		Expect(w.Objects).To(HaveLen(2))
		mat1 := w.Objects[0].Material()
		Expect(mat1.Color).To(color.Equal(color.New(0.8, 1.0, 0.6)))
		Expect(mat1.Diffuse).To(BeNumerically("~", 0.7))
		Expect(mat1.Specular).To(BeNumerically("~", 0.2))
		trans2 := w.Objects[1].GetTransform()
		Expect(trans2).To(matrix.Equal(matrix.Scaling(0.5, 0.5, 0.5)))
	})

	It("can find ray intersections with a world", func() {
		w := world.Default()
		r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
		xs := w.Intersections(r)
		Expect(xs.Count()).To(Equal(4))
		Expect(xs.Get(0).T).To(BeNumerically("~", 4))
		Expect(xs.Get(1).T).To(BeNumerically("~", 4.5))
		Expect(xs.Get(2).T).To(BeNumerically("~", 5.5))
		Expect(xs.Get(3).T).To(BeNumerically("~", 6))
	})

	It("can shade an intersection", func() {
		w := world.Default()
		r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
		s := w.Objects[0]
		ix := shape.NewIntersections()
		ix.Add(4, s)
		i := ix.Get(0)
		comps := i.PrepareComputations(r, ix)
		c := w.ShadeHit(comps)
		Expect(c).To(color.Equal(color.New(0.38066, 0.47583, 0.2855)))
	})

	It("can shade an intersection from the inside", func() {
		w := world.Default()
		l := light.NewPoint(tuple.Point(0, 0.25, 0), color.New(1, 1, 1))
		w.LightSource = &l
		r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
		s := w.Objects[1]
		ix := shape.NewIntersections()
		ix.Add(0.5, s)
		i := ix.Get(0)
		comps := i.PrepareComputations(r, ix)
		c := w.ShadeHit(comps)
		Expect(c).To(color.Equal(color.New(0.90498, 0.90498, 0.90498)))
	})

	Context("color for a ray", func() {
		It("gives black for a missing ray", func() {
			w := world.Default()
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 1, 0))
			c := w.ColorAt(r)
			Expect(c).To(color.Equal(color.New(0, 0, 0)))
		})

		It("gets an intersecting ray color correct", func() {
			w := world.Default()
			r := ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			c := w.ColorAt(r)
			Expect(c).To(color.Equal(color.New(0.38066, 0.47583, 0.2855)))
		})

		It("gets an inside intersection correct", func() {
			w := world.Default()
			outer := w.Objects[0]
			m := outer.Material()
			m.Ambient = 1
			outer.SetMaterial(m)
			inner := w.Objects[1]
			mInner := inner.Material()
			mInner.Ambient = 1
			inner.SetMaterial(mInner)
			r := ray.New(tuple.Point(0, 0, 0.75), tuple.Vector(0, 0, -1))
			c := w.ColorAt(r)
			Expect(c).To(color.Equal(inner.Material().Color))
		})

		It("gets a shadow color correct", func() {
			w := world.New()
			lightSource := light.NewPoint(tuple.Point(0, 0, -10), color.New(1, 1, 1))
			w.LightSource = &lightSource
			s1 := shape.NewSphere()
			w.AddObject(s1)
			s2 := shape.NewSphere()
			s2.SetTransform(matrix.Translation(0, 0, 10))
			w.AddObject(s2)
			r := ray.New(tuple.Point(0, 0, 5), tuple.Vector(0, 0, 1))
			ix := shape.NewIntersections()
			ix.Add(4, s2)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			c := w.ShadeHit(comps)
			Expect(c).To(color.Equal(color.New(0.1, 0.1, 0.1)))
		})
	})

	DescribeTable("in shadow?", func(point tuple.Tuple, inShadow bool) {
		w := world.Default()
		Expect(w.InShadow(point)).To(Equal(inShadow))
	},

		Entry("nothing colinear with sphere and light", tuple.Point(0, 10, 0), false),
		Entry("sphere between light and point", tuple.Point(10, -10, 10), true),
		Entry("sphere behind the light", tuple.Point(-20, 20, -20), false),
		Entry("object behind the point", tuple.Point(-2, 2, -2), false),
	)

	Context("reflection", func() {
		It("returns black when a ray reflects from a non-reflective surface", func() {
			w := world.Default()
			r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
			s := w.Objects[1]
			m := s.Material()
			m.Ambient = 1
			s.SetMaterial(m)
			ix := shape.NewIntersections()
			ix.Add(1, s)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			c := w.ReflectedColor(comps, 1)
			Expect(c).To(color.Equal(color.New(0, 0, 0)))
		})

		It("returns a reflected color from a reflective surface", func() {
			r2 := math.Sqrt(2)

			w := world.Default()
			s := shape.NewPlane()
			s.SetTransform(matrix.Translation(0, -1, 0))
			m := s.Material()
			m.Reflective = 0.5
			s.SetMaterial(m)
			w.AddObject(s)
			r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -r2/2, r2/2))
			ix := shape.NewIntersections()
			ix.Add(r2, s)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			c := w.ReflectedColor(comps, 1)
			Expect(c).To(color.Equal(color.New(0.19033, 0.23791, 0.14274)))
		})

		It("returns a shadeHit incorporating color from a reflective surface", func() {
			r2 := math.Sqrt(2)

			w := world.Default()
			s := shape.NewPlane()
			s.SetTransform(matrix.Translation(0, -1, 0))
			m := s.Material()
			m.Reflective = 0.5
			s.SetMaterial(m)
			w.AddObject(s)
			r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -r2/2, r2/2))
			ix := shape.NewIntersections()
			ix.Add(r2, s)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			c := w.ShadeHit(comps)
			Expect(c).To(color.Equal(color.New(0.87675, 0.92434, 0.82917)))
		})

		It("avoids infinite recursion", func() {
			w := world.New()
			l := light.NewPoint(tuple.Point(0, 0, 0), color.New(1, 1, 1))
			w.LightSource = &l

			lower := shape.NewPlane()
			lower.SetTransform(matrix.Translation(0, -1, 0))
			ml := lower.Material()
			ml.Reflective = 1
			lower.SetMaterial(ml)
			w.AddObject(lower)

			upper := shape.NewPlane()
			upper.SetTransform(matrix.Translation(0, 1, 0))
			mu := upper.Material()
			mu.Reflective = 1
			upper.SetMaterial(mu)
			w.AddObject(upper)

			r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 1, 0))
			Expect(func() { w.ColorAt(r) }).ToNot(Panic())
		})

		It("respects the remaining parameter", func() {
			r2 := math.Sqrt(2)

			w := world.Default()
			s := shape.NewPlane()
			s.SetTransform(matrix.Translation(0, -1, 0))
			m := s.Material()
			m.Reflective = 0.5
			s.SetMaterial(m)
			w.AddObject(s)
			r := ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -r2/2, r2/2))
			ix := shape.NewIntersections()
			ix.Add(r2, s)
			i := ix.Get(0)
			comps := i.PrepareComputations(r, ix)
			c := w.ReflectedColor(comps, 0)
			Expect(c).To(color.Equal(color.New(0, 0, 0)))
		})
	})

	Context("refraction", func() {

		var (
			w     *world.World
			s     *shape.Object
			r     ray.Ray
			ix    *shape.Intersections
			i     *shape.Intersection
			comps shape.Computations
		)

		BeforeEach(func() {
			w = world.Default()
			s = w.Objects[0]
			r = ray.New(tuple.Point(0, 0, -5), tuple.Vector(0, 0, 1))
			ix = shape.NewIntersections()
			ix.Add(4, s)
			ix.Add(6, s)
			i = ix.Get(0)
		})

		It("returns black for an opaque surface", func() {
			comps = i.PrepareComputations(r, ix)
			c := w.RefractedColor(comps, 5)
			Expect(c).To(color.Equal(color.New(0, 0, 0)))
		})

		When("recursion has run out", func() {
			It("returns black", func() {
				m := material.New()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				s.SetMaterial(m)

				comps = i.PrepareComputations(r, ix)
				c := w.RefractedColor(comps, 0)
				Expect(c).To(color.Equal(color.New(0, 0, 0)))
			})
		})

		When("there is total internal reflection", func() {
			It("returns black", func() {
				r2 := math.Sqrt(2)
				r = ray.New(tuple.Point(0, 0, r2/2), tuple.Vector(0, 1, 0))
				ix = shape.NewIntersections()
				ix.Add(-r2/2, s)
				ix.Add(r2/2, s)
				i = ix.Get(1)
				m := material.New()
				m.Transparency = 1.0
				m.RefractiveIndex = 1.5
				s.SetMaterial(m)
				comps = i.PrepareComputations(r, ix)
				c := w.RefractedColor(comps, 5)
				Expect(c).To(color.Equal(color.New(0, 0, 0)))
			})
		})

		When("there is real refraction", func() {
			It("sends the correct new ray off", func() {
				ma := material.New()
				ma.Ambient = 1.0
				fakePattern := new(fakes.FakeActualPattern)
				pattern := pattern.New(fakePattern)
				ma.SetPattern(&pattern)
				w.Objects[0].SetMaterial(ma)

				mb := material.New()
				mb.Transparency = 1.0
				mb.RefractiveIndex = 1.5
				w.Objects[1].SetMaterial(mb)

				r = ray.New(tuple.Point(0, 0, 0.1), tuple.Vector(0, 1, 0))
				ix = shape.NewIntersections()
				ix.Add(-0.9899, w.Objects[0])
				ix.Add(-0.4899, w.Objects[1])
				ix.Add(0.4899, w.Objects[1])
				ix.Add(0.9899, w.Objects[0])

				comps := ix.Get(2).PrepareComputations(r, ix)
				w.RefractedColor(comps, 5)

				Expect(fakePattern.PatternAtCallCount()).To(Equal(1))
				p := fakePattern.PatternAtArgsForCall(0)
				Expect(p).To(tuple.Equal(tuple.Point(0, 0.99888, 0.04721)))
			})
		})

		It("is handled in ShadeHit", func() {
			floor := shape.NewPlane()
			fm := material.New()
			fm.Transparency = 0.5
			fm.RefractiveIndex = 1.5
			floor.SetMaterial(fm)
			floor.SetTransform(matrix.Translation(0, -1, 0))
			w.AddObject(floor)

			ball := shape.NewSphere()
			bm := material.New()
			bm.Color = color.New(1, 0, 0)
			bm.Ambient = 0.5
			ball.SetMaterial(bm)
			ball.SetTransform(matrix.Translation(0, -3.5, -0.5))
			w.AddObject(ball)

			r2 := math.Sqrt(2)
			r = ray.New(tuple.Point(0, 0, -3), tuple.Vector(0, -r2/2, r2/2))
			ix = shape.NewIntersections()
			ix.Add(r2, floor)

			comps := ix.Get(0).PrepareComputations(r, ix)
			c := w.ShadeHit(comps, 5)
			Expect(c).To(color.Equal(color.New(0.93642, 0.68642, 0.68642)))
		})
	})

})
