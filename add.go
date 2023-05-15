package ptr

import "unsafe"

type Int interface {
	int | int32 | int64 | uintptr
}

func Add[T any, U Int](base *T, offset U) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(base)) + uintptr(offset))
}

func AddOffset[T any, U Int](base *T, elementSize, offset U) unsafe.Pointer {
	return Add(base, elementSize*offset)
}
