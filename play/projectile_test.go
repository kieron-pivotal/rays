package play_test

import (
	"fmt"
	"math"
	"os"

	"github.com/kieron-pivotal/rays/canvas"
	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/play"
	"github.com/kieron-pivotal/rays/tuple"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Projectile", func() {

	var (
		gravity tuple.Tuple
		wind    tuple.Tuple
	)

	BeforeEach(func() {
		gravity = tuple.Vector(0, -0.1, 0)
		wind = tuple.Vector(-0.01, 0, 0)
	})

	It("can fire a projectile", func() {
		env := play.NewEnv(gravity, wind)
		Expect(env).ToNot(BeNil())

		startPosition := tuple.Point(0, 1, 0)
		initialVelocity := tuple.Vector(1, 1, 0).Normalize()
		speedFactor := 10.0
		trace := env.FireProjectile(startPosition, initialVelocity.Multiply(speedFactor))

		fmt.Printf("len(trace) = %+v\n", len(trace))

		for _, p := range trace {
			fmt.Printf("p = %+v\n", p)
		}
	})

	It("can plot the project to PPM", func() {
		env := play.NewEnv(gravity, wind)
		Expect(env).ToNot(BeNil())

		startPosition := tuple.Point(0, 1, 0)
		initialVelocity := tuple.Vector(1, 1.8, 0).Normalize()
		speedFactor := 12.0
		trace := env.FireProjectile(startPosition, initialVelocity.Multiply(speedFactor))

		maxX, maxY := 0.0, 0.0
		for _, p := range trace {
			if p.X > maxX {
				maxX = p.X
			}
			if p.Y > maxY {
				maxY = p.Y
			}
		}

		width := 900
		height := 500
		xFactor := float64(width-1) / maxX
		yFactor := float64(height-1) / maxY
		factor := math.Min(xFactor, yFactor)

		canvas := canvas.New(width, height)
		red := color.New(1, 0, 0)

		for _, p := range trace {
			x := int(math.Round(p.X * factor))
			y := int(math.Round(p.Y * factor))
			canvas.SetPixel(x, height-y-1, red)
		}

		file, err := os.Create("projectile.ppm")
		Expect(err).NotTo(HaveOccurred())

		file.WriteString(canvas.ToPPM())
		file.Close()
	})
})
