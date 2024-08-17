package kmp

import (
	"testing"
)

// TestPresence calls a case where substring match exists.
func TestVPresence(t *testing.T) {
	pattern, text := "mama", "ammamaa"
	if !vertical(pattern, text) {
		t.Fatalf(`vertical("mama", "ammamaa") = false, want true`)
	}
}

// TestAbsence calls a case where substring match exists.
func TestVAbsence(t *testing.T) {
	pattern, text := "mava", "ammamaa"
	if vertical(pattern, text) {
		t.Fatalf(`vertical("mava", "ammamaa") = true, want false`)
	}
}
