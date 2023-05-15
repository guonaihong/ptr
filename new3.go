package ptr

func New3[T any](v T) ***T {
	p1 := &v
	p2 := &p1
	return &p2
}
