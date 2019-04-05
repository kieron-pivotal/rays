package play_test

import (
	"fmt"

	"github.com/kieron-pivotal/rays/geometry"
	"github.com/kieron-pivotal/rays/play"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Projectile", func() {

	It("can fire a projectile", func() {
		gravity := geometry.Vector(0, -0.1, 0)
		wind := geometry.Vector(-0.01, 0, 0)
		env := play.NewEnv(gravity, wind)
		Expect(env).ToNot(BeNil())

		startPosition := geometry.Point(0, 1, 0)
		initialVelocity := geometry.Vector(1, 1, 0).Normalize()
		speedFactor := 10.0
		trace := env.FireProjectile(startPosition, initialVelocity.Multiply(speedFactor))

		fmt.Printf("len(trace) = %+v\n", len(trace))

		for _, p := range trace {
			fmt.Printf("p = %+v\n", p)
		}
	})

})
