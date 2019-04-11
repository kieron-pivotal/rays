package shape

type Intersections struct {
	list []Intersection
}

type Intersection struct {
	T float64
	S *Sphere
}

func NewIntersections() *Intersections {
	return &Intersections{}
}

func (i Intersections) Count() int {
	return len(i.list)
}

func (i Intersections) Get(idx int) Intersection {
	return i.list[idx]
}

func (i *Intersections) Add(t float64, s *Sphere) {
	i.list = append(i.list, Intersection{T: t, S: s})
}
