package ptr

import "unsafe"

func Sub[T any, U Int](base *T, offset U) unsafe.Pointer {
	return Add(base, offset)
}
