package kmp

import (
	"testing"
)

// TestPresence calls a case where substring match exists.
func TestMPPresence(t *testing.T) {
	pattern, text := "mama", "ammamaa"
	if !morris_pratt(pattern, text) {
		t.Fatalf(`vertical("mama", "ammamaa") = false, want true`)
	}
}

// TestAbsence calls a case where substring match exists.
func TestMPAbsence(t *testing.T) {
	pattern, text := "mava", "ammamaa"
	if morris_pratt(pattern, text) {
		t.Fatalf(`vertical("mava", "ammamaa") = true, want false`)
	}
}
