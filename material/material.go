package material

import (
	"math"

	"github.com/kieron-pivotal/rays/color"
	"github.com/kieron-pivotal/rays/light"
	"github.com/kieron-pivotal/rays/tuple"
)

type Material struct {
	Color     color.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func Default() Material {
	return Material{
		Color:     color.New(1, 1, 1),
		Ambient:   0.1,
		Diffuse:   0.9,
		Specular:  0.9,
		Shininess: 200,
	}
}

func (m Material) Lighting(l light.Point, pos, eye, normal tuple.Tuple) color.Color {

	black := color.New(0, 0, 0)
	var ambient, diffuse, specular color.Color

	effectiveColor := m.Color.ColorMultiply(l.Intensity)
	ambient = effectiveColor.Multiply(m.Ambient)

	lightV := l.Position.Subtract(pos).Normalize()
	lightDotNormal := lightV.Dot(normal)
	if lightDotNormal < 0 {
		diffuse = black
		specular = black
	} else {
		diffuse = effectiveColor.Multiply(m.Diffuse).Multiply(lightDotNormal)

		reflectV := lightV.Multiply(-1).Reflect(normal)
		reflectDotEye := reflectV.Dot(eye)

		if reflectDotEye <= 0 {
			specular = black
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = l.Intensity.Multiply(m.Specular).Multiply(factor)
		}
	}

	return ambient.Add(diffuse).Add(specular)
}
