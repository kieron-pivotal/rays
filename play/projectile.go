package play

import "github.com/kieron-pivotal/rays/tuple"

type Env struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
}

func NewEnv(gravity, wind tuple.Tuple) *Env {
	env := Env{
		gravity: gravity,
		wind:    wind,
	}
	return &env
}

func (e *Env) FireProjectile(p, v tuple.Tuple) []tuple.Tuple {
	out := []tuple.Tuple{}
	for p.Y >= 0 {
		out = append(out, p)
		p, v = e.tick(p, v)
	}

	return out
}

func (e *Env) tick(p, v tuple.Tuple) (newP, newV tuple.Tuple) {
	newP = p.Add(v)
	newV = v.Add(e.gravity).Add(e.wind)
	return
}
