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
)

var _ = Describe("Cube", func() {

	It("can draw a table", func() {

		w := world.New()

		f := shape.NewPlane()
		f.SetTransform(matrix.RotationX(math.Pi / 2))
		fMat := material.New()
		fPat := pattern.NewChecker(color.New(0, 0, 0), color.New(1, 1, 1))
		fPat.SetTransform(matrix.Translation(0, 0.0001, 0).Scale(15, 15, 15))
		fMat.SetPattern(&fPat)
		f.SetMaterial(fMat)
		w.AddObject(f)

		l := light.NewPoint(tuple.Point(-50, -100, 100), color.New(1, 1, 1))
		w.LightSource = &l

		tMat := material.New()
		tMat.Color = color.New(165.0/255.0, 42.0/255.0, 42.0/255.0)

		l1 := shape.NewCube()
		l1.SetTransform(matrix.Scaling(2, 2, 15).Translate(-49, -19, 15))
		l1.SetMaterial(tMat)
		w.AddObject(l1)

		l2 := shape.NewCube()
		l2.SetTransform(matrix.Scaling(2, 2, 15).Translate(-49, 19, 15))
		l2.SetMaterial(tMat)
		w.AddObject(l2)

		l3 := shape.NewCube()
		l3.SetTransform(matrix.Scaling(2, 2, 15).Translate(49, -19, 15))
		l3.SetMaterial(tMat)
		w.AddObject(l3)

		l4 := shape.NewCube()
		l4.SetTransform(matrix.Scaling(2, 2, 15).Translate(49, 19, 15))
		l4.SetMaterial(tMat)
		w.AddObject(l4)

		top := shape.NewCube()
		top.SetTransform(matrix.Scaling(52, 22, 1).Translate(0, 0, 30))
		top.SetMaterial(tMat)
		w.AddObject(top)

		box := shape.NewCube()
		box.SetTransform(matrix.Scaling(3, 3, 3).RotateZ(math.Pi/8).Translate(0, 0, 34))
		bMat := material.New()
		bMat.Color = color.New(0, 0, 1)
		bMat.Reflective = 0.9
		box.SetMaterial(bMat)
		w.AddObject(box)

		cam := camera.New(480, 320, math.Pi/4)
		cam.SetTransform ( matrix.ViewTransformation(
			tuple.Point(-100, -100, 70),
			tuple.Point(0, 0, 15),
			tuple.Vector(0, 0, 1),
		))

		canvas := cam.Render(w)

		out, err := os.Create("table.ppm")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		img := canvas.ToPPM()
		fmt.Fprint(out, img)

	})

})
