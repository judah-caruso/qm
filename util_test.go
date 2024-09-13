package qm_test

import (
	"testing"
)

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
