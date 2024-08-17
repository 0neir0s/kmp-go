package kmp

import (
	"testing"
)

// TestPresence calls a case where substring match exists.
func TestVLPresence(t *testing.T) {
	pattern, text := "mama", "ammamaa"
	if !verticalList(pattern, text) {
		t.Fatalf(`vertical("mama", "ammamaa") = false, want true`)
	}
}

// TestAbsence calls a case where substring match exists.
func TestVLAbsence(t *testing.T) {
	pattern, text := "mava", "ammamaa"
	if verticalList(pattern, text) {
		t.Fatalf(`vertical("mava", "ammamaa") = true, want false`)
	}
}
