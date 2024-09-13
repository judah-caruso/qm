package qm_test

import (
	"github.com/judah-caruso/qm"
	"testing"
	"unsafe"
)

func TestPtrCast(t *testing.T) {
	f := qm.Pi
	u := *(*uint32)(unsafe.Pointer(&f))
	f2 := *(*float32)(unsafe.Pointer(&u))
	Expect(t, f2 == f)
}

func Expectf(t *testing.T, cond bool, format string, args ...any) {
	t.Helper()
	if !cond {
		t.Fatalf(format, args...)
	}
}

func Expect(t *testing.T, cond bool) {
	t.Helper()
	Expectf(t, cond, "call to Expect failed")
}
