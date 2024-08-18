package kmp

import (
	"fmt"
	"strings"
)

func verticalList(pattern, text string) bool {
	var any func(func([]string) bool, [][]string) bool
	any = func(done func([]string) bool, accs [][]string) bool {
		fmt.Println(accs)
		if len(accs) == 0 {
			return false
		}
		acc, tail := accs[0], accs[1:]
		return done(acc) || any(done, tail)
	}

	var scanl func(func([]string, string) []string, []string, []string) [][]string
	scanl = func(step func([]string, string) []string, acc []string, slice []string) [][]string {
		if len(slice) == 0 {
			return [][]string{acc}
		}
		x, xt := slice[0], slice[1:]
		return append([][]string{acc}, scanl(step, step(acc, x), xt)...)
	}
	init := []string{pattern}

	check := func(candidate, x string) bool {
		if len(candidate) == 0 {
			return false
		}
		return strings.HasPrefix(candidate, x)
	}

	var step func([]string, string) []string
	step = func(acc []string, x string) []string {
		if len(acc) == 0 {
			return init
		}
		candidate := acc[0]
		if check(candidate, x) {
			return append([]string{candidate[1:]}, step(acc[1:], x)...)
		}
		return step(acc[1:], x)
	}

	done := func(acc []string) bool {
		if len(acc) == 0 {
			return false
		}
		return acc[0] == ""
	}
	return any(done, scanl(step, init, strings.Split(text, "")))
}
