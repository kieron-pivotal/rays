package world_test

import (
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/ray"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
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
		i := shape.Intersection{
			T:      4,
			Object: s,
		}
		comps := i.PrepareComputations(r)
		c := w.ShadeHit(comps)
		Expect(c).To(color.Equal(color.New(0.38066, 0.47583, 0.2855)))
	})

	It("can shade an intersection from the inside", func() {
		w := world.Default()
		l := light.NewPoint(tuple.Point(0, 0.25, 0), color.New(1, 1, 1))
		w.LightSource = &l
		r := ray.New(tuple.Point(0, 0, 0), tuple.Vector(0, 0, 1))
		s := w.Objects[1]
		i := shape.Intersection{
			T:      0.5,
			Object: s,
		}
		comps := i.PrepareComputations(r)
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
	})

})
