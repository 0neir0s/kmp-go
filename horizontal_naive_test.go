package kmp

import (
	"testing"
)

// TestPresence calls a case where substring match exists.
func TestPresence(t *testing.T) {
	pattern, text := "mama", "ammamaa"
	if !horizontal(pattern, text) {
		t.Fatalf(`horizontal("mama", "ammamaa") = false, want true`)
	}
}

// TestAbsence calls a case where substring match exists.
func TestAbsence(t *testing.T) {
	pattern, text := "mava", "ammamaa"
	if horizontal(pattern, text) {
		t.Fatalf(`horizontal("mava", "ammamaa") = true, want false`)
	}
}
