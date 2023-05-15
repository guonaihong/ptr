package ptr

import (
	"testing"
	"unsafe"
)

func Test_Add(t *testing.T) {
	s := []int{1, 2}
	p := Add(&s[0], unsafe.Sizeof(s[0]))
	if *(*uintptr)(p) != 2 {
		t.Fatal("*p != 3")
	}
}

func Test_AddOffset(t *testing.T) {
	s := []int{1, 2}
	p := AddOffset(&s[0], unsafe.Sizeof(s[0]), 1)
	if *(*uintptr)(p) != 2 {
		t.Fatal("*p != 3")
	}
}
