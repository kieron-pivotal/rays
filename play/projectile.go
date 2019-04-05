package play

import "github.com/kieron-pivotal/rays/geometry"

type Env struct {
	gravity geometry.Tuple
	wind    geometry.Tuple
}

func NewEnv(gravity, wind geometry.Tuple) *Env {
	env := Env{
		gravity: gravity,
		wind:    wind,
	}
	return &env
}

func (e *Env) FireProjectile(p, v geometry.Tuple) []geometry.Tuple {
	out := []geometry.Tuple{}
	for p.Y >= 0 {
		out = append(out, p)
		p, v = e.tick(p, v)
	}

	return out
}

func (e *Env) tick(p, v geometry.Tuple) (newP, newV geometry.Tuple) {
	newP = p.Add(v)
	newV = v.Add(e.gravity).Add(e.wind)
	return
}
