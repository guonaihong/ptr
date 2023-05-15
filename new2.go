package ptr

func New2[T any](v T) **T {
	p1 := &v
	return &p1
}
