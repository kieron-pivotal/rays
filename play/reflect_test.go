package play_test

import (
	"fmt"
	"log"
	"math"
	"os"

	"github.com/kieron-pivotal/rays/camera"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/material"
	"github.com/kieron-pivotal/rays/matrix"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Reflect", func() {

	FIt("does something", func() {
		w := world.New()
		l := light.NewPoint(tuple.Point(0, 10, 10), color.New(1, 1, 1))
		w.LightSource = &l

		plane := shape.NewPlane()
		mp := material.New()
		mp.Reflective = 1
		mp.Color = color.New(1, 0, 0)
		plane.SetMaterial(mp)
		w.AddObject(plane)

		sphere := shape.NewSphere()
		sphere.SetTransform(matrix.Scaling(2, 2, 2).Translate(0, 2, 0))
		ms := material.New()
		ms.Color = color.New(0, 1, 0)
		sphere.SetMaterial(ms)
		w.AddObject(sphere)

		camera := camera.New(500, 300, math.Pi/3)
		camera.Transform = matrix.ViewTransformation(
			tuple.Point(-15, 2, 5),
			tuple.Point(0, 0, 0),
			tuple.Vector(0, 1, 0),
		)

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
