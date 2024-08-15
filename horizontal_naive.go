package kmp

import (
	"fmt"
	"strings"
)

func scanl(f func(string, string) string, acc string, slice []string) []string {
	if len(slice) == 0 {
		return []string{acc}
	}
	x, xt := slice[0], slice[1:]
	return append([]string{acc}, scanl(f, f(acc, x), xt)...)
}

func any(f func(string) bool, slice []string) bool {
	fmt.Println(slice)
	if len(slice) == 0 {
		return false
	}
	return f(slice[0]) || any(f, slice[1:])
}

func horizontal(pattern, text string) bool {
	step := func(acc string, _ string) string {
		return acc[1:]
	}
	done := func(acc string) bool {
		return strings.HasPrefix(acc, pattern)
	}
	return any(done, scanl(step, text, strings.Split(text, "")))
}
