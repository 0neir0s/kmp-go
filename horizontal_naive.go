package kmp

import (
	"strings"
)

func horizontal(pattern, text string) bool {

	var any func(func(string) bool, []string) bool
	any = func(done func(string) bool, accs []string) bool {
		if len(accs) == 0 {
			return false
		}
		acc, tail := accs[0], accs[1:]
		return done(acc) || any(done, tail)
	}
	var scanl func(func(string, string) string, string, []string) []string
	scanl = func(step func(string, string) string, acc string, slice []string) []string {
		if len(slice) == 0 {
			return []string{acc}
		}
		x, xt := slice[0], slice[1:]
		return append([]string{acc}, scanl(step, step(acc, x), xt)...)
	}
	init := text
	step := func(acc string, _ string) string {
		return acc[1:]
	}
	done := func(acc string) bool {
		return strings.HasPrefix(acc, pattern)
	}
	return any(done, scanl(step, init, strings.Split(text, "")))
}
