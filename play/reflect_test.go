package play_test

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/kieron-pivotal/rays/camera"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Reflect", func() {

	It("does something", func() {

		f, err := os.Create("cpu.pprof")
		Expect(err).NotTo(HaveOccurred())
		err = pprof.StartCPUProfile(f)
		Expect(err).NotTo(HaveOccurred())

		defer pprof.StopCPUProfile()

		w := world.New()
		l := light.NewPoint(tuple.Point(0, 10, 10), color.New(1, 1, 1))
		w.LightSource = &l

		plane := shape.NewPlane()
		mp := material.New()
		mp.Reflective = 0.3
		mpp := pattern.NewChecker(color.New(0.6, 0.5, 0.5), color.New(0, 0, 0))
		mpp.SetTransform(matrix.Scaling(1.5, 1.5, 1.5).Translate(0, 0.0001, 0))
		mp.SetPattern(&mpp)
		mp.Specular = 0
		plane.SetMaterial(mp)
		w.AddObject(plane)

		sphere1 := shape.NewSphere()
		sphere1.SetTransform(matrix.Scaling(2, 2, 2).Translate(0, 2, 0))
		ms := material.New()
		ms.Reflective = 0.5
		ms.Transparency = 0.9
		ms.RefractiveIndex = 1.5
		ms.Ambient = 0.3
		ms.Color = color.New(0, 0.2, 0)
		sphere1.SetMaterial(ms)
		w.AddObject(sphere1)

		sphere2 := shape.NewSphere()
		sphere2.SetTransform(matrix.Translation(-4, 1, -2))
		ms2 := material.New()
		ms2.Reflective = 0.3
		ms2.Color = color.New(0, 0, 1)
		sphere2.SetMaterial(ms2)
		w.AddObject(sphere2)

		camera := camera.New(600, 400, math.Pi/3)
		camera.SetTransform(matrix.ViewTransformation(
			tuple.Point(-12, 2.8, 0),
			tuple.Point(0, 0, 0),
			tuple.Vector(0, 1, 0),
		))

		canvas := camera.Render(w)

		out, err := os.Create("reflect.ppm")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		img := canvas.ToPPM()
		fmt.Fprint(out, img)

	})

})
