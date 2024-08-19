package kmp

import (
	"testing"
)

// TestPresence calls a case where substring match exists.
func TestKMPPresence(t *testing.T) {
	pattern, text := "mama", "ammamaa"
	if !knuth_morris_pratt(pattern, text) {
		t.Fatalf(`vertical("mama", "ammamaa") = false, want true`)
	}
}

// TestAbsence calls a case where substring match exists.
func TestKMPAbsence(t *testing.T) {
	pattern, text := "mava", "ammamaa"
	if knuth_morris_pratt(pattern, text) {
		t.Fatalf(`vertical("mava", "ammamaa") = true, want false`)
	}
}
