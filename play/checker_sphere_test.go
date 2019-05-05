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
	"github.com/kieron-pivotal/rays/pattern"
	"github.com/kieron-pivotal/rays/shape"
	"github.com/kieron-pivotal/rays/tuple"
	"github.com/kieron-pivotal/rays/world"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("CheckerSphere", func() {

	FIt("draws a checkered sphere", func() {
		w := world.New()

		s := shape.NewSphere()
		s.SetTransform(matrix.Scaling(5, 5, 5))

		p := pattern.NewChecker(color.New(1, 1, 1), color.New(0, 0, 0))
		p.SetTransform(matrix.Scaling(0.3, 0.3, 0.3).RotateX(-math.Pi / 8))

		m := material.New()
		m.Ambient = 0.2
		m.SetPattern(&p)

		s.SetMaterial(m)

		w.AddObject(s)

		l := light.NewPoint(tuple.Point(10, 10, 10), color.New(1, 1, 1))
		w.LightSource = &l

		camera := camera.New(300, 200, math.Pi/4)
		camera.Transform = matrix.ViewTransformation(
			tuple.Point(0, 20, 0),
			tuple.Point(0, 0, 0),
			tuple.Vector(1, 0, 0),
		)

		canvas := camera.Render(w)

		out, err := os.Create("checker_sphere.ppm")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		img := canvas.ToPPM()
		fmt.Fprint(out, img)
	})

})
