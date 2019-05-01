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

var _ = Describe("WorldOne", func() {
	FIt("can draw the scene", func() {

		w := world.New()

		p := pattern.NewStripe(color.New(0.8, 0.8, 0.8), color.New(0.1, 0.1, 0.1))
		p.SetTransform(matrix.Scaling(0.25, 0.25, 0.25))
		p2 := pattern.NewGradient(color.New(0.8, 0.2, 0.1), color.New(0.1, 0.1, 0.1))

		floor := shape.NewSphere()
		floor.SetTransform(matrix.Scaling(10, 0.01, 10))
		wallMaterial := material.New()
		wallMaterial.Color = color.New(1, 0.9, 0.9)
		wallMaterial.Specular = 0
		wallMaterial.SetPattern(&p)
		floor.SetMaterial(wallMaterial)
		w.AddObject(floor)

		leftWall := shape.NewSphere()
		leftWall.SetTransform(matrix.Identity(4, 4).Scale(10, 0.01, 10).
			RotateX(math.Pi/2).
			RotateY(-math.Pi/4).
			Translate(0, 0, 5))
		leftWall.SetMaterial(wallMaterial)
		w.AddObject(leftWall)

		rightWall := shape.NewSphere()
		rightWall.SetTransform(matrix.Identity(4, 4).Scale(10, 0.01, 10).
			RotateX(math.Pi/2).
			RotateY(math.Pi/4).
			RotateZ(math.Pi/8).
			Translate(0, 0, 5))
		rightWall.SetMaterial(wallMaterial)
		w.AddObject(rightWall)

		middle := shape.NewSphere()
		middle.SetTransform(matrix.Translation(-0.5, 1, 0.5))
		middleMaterial := material.New()
		middleMaterial.Color = color.New(0.1, 1, 0.5)
		middleMaterial.Diffuse = 0.7
		middleMaterial.Specular = 0.3
		middleMaterial.SetPattern(&p2)
		middle.SetMaterial(middleMaterial)
		w.AddObject(middle)

		left := shape.NewSphere()
		left.SetTransform(matrix.Identity(4, 4).Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75))
		leftMaterial := material.New()
		leftMaterial.Color = color.New(1, 0.8, 0.1)
		leftMaterial.Diffuse = 0.7
		leftMaterial.Specular = 0.3
		leftMaterial.SetPattern(&p)
		left.SetMaterial(leftMaterial)
		w.AddObject(left)

		right := shape.NewSphere()
		right.SetTransform(matrix.Identity(4, 4).RotateZ(math.Pi/2).Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5))
		rightMaterial := material.New()
		rightMaterial.Color = color.New(0.5, 1, 0.1)
		rightMaterial.Diffuse = 0.7
		rightMaterial.Specular = 0.3
		rightMaterial.SetPattern(&p)
		right.SetMaterial(rightMaterial)
		w.AddObject(right)

		lightSource := light.NewPoint(tuple.Point(-10, 10, -10), color.New(1, 1, 1))
		w.LightSource = &lightSource

		camera := camera.New(300, 180, math.Pi/3)
		camera.Transform = matrix.ViewTransformation(
			tuple.Point(0, 1.5, -5),
			tuple.Point(0, 1, 0),
			tuple.Vector(0, 1, 0),
		)

		canvas := camera.Render(w)

		out, err := os.Create("world_one.ppm")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		img := canvas.ToPPM()
		fmt.Fprint(out, img)
	})

	It("can draw the scene with a plane", func() {

		w := world.New()

		floor := shape.NewPlane()
		// floor.SetTransform(matrix.Scaling(10, 0.01, 10))
		wallMaterial := material.New()
		wallMaterial.Color = color.New(1, 0.9, 0.9)
		wallMaterial.Specular = 0
		floor.SetMaterial(wallMaterial)
		w.AddObject(floor)

		for i := 0; i < 6; i++ {
			leftWall := shape.NewPlane()
			leftWall.SetTransform(
				matrix.Identity(4, 4).
					RotateX(math.Pi/2).
					Translate(0, 0, 4).
					RotateY(float64(i) * math.Pi / 3))
			leftWall.SetMaterial(wallMaterial)
			w.AddObject(leftWall)
		}
		//
		// rightWall := shape.NewSphere()
		// rightWall.SetTransform(matrix.Identity(4, 4).Scale(10, 0.01, 10).
		// 	RotateX(math.Pi/2).
		// 	RotateY(math.Pi/4).
		// 	Translate(0, 0, 5))
		// rightWall.SetMaterial(wallMaterial)
		// w.AddObject(rightWall)

		middle := shape.NewSphere()
		middle.SetTransform(matrix.Translation(-0.5, 1, 0.5))
		middleMaterial := material.New()
		middleMaterial.Color = color.New(0.1, 1, 0.5)
		middleMaterial.Diffuse = 0.7
		middleMaterial.Specular = 0.3
		middle.SetMaterial(middleMaterial)
		w.AddObject(middle)

		left := shape.NewSphere()
		left.SetTransform(matrix.Identity(4, 4).Scale(0.33, 0.33, 0.33).Translate(-1.5, 0.33, -0.75))
		leftMaterial := material.New()
		leftMaterial.Color = color.New(1, 0.8, 0.1)
		leftMaterial.Diffuse = 0.7
		leftMaterial.Specular = 0.3
		left.SetMaterial(leftMaterial)
		w.AddObject(left)

		right := shape.NewSphere()
		right.SetTransform(matrix.Identity(4, 4).Scale(0.5, 0.5, 0.5).Translate(1.5, 0.5, -0.5))
		rightMaterial := material.New()
		rightMaterial.Color = color.New(0.5, 1, 0.1)
		rightMaterial.Diffuse = 0.7
		rightMaterial.Specular = 0.3
		right.SetMaterial(rightMaterial)
		w.AddObject(right)

		lightSource := light.NewPoint(tuple.Point(0, 10, 0), color.New(1, 1, 1))
		w.LightSource = &lightSource

		camera := camera.New(150, 90, math.Pi/2)
		camera.Transform = matrix.ViewTransformation(
			tuple.Point(-3, 4, 0),
			tuple.Point(0, 0, 0),
			tuple.Vector(0, 0, 1),
		)

		canvas := camera.Render(w)

		out, err := os.Create("world_two.ppm")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()

		img := canvas.ToPPM()
		fmt.Fprint(out, img)
	})
})
