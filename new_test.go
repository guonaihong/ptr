package ptr

import "testing"

func Test_New(t *testing.T) {
	p := New(3)
	if *p != 3 {
		t.Fatal("*p != 3")
	}
}

func Test_New2(t *testing.T) {
	p := New2(3)
	if **p != 3 {
		t.Fatal("**p != 3")
	}
}

func Test_New3(t *testing.T) {
	p := New3(3)
	if ***p != 3 {
		t.Fatal("**p != 3")
	}
}
