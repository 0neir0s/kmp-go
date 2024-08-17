package kmp

import (
	"strings"
)

func union(s1, s2 map[string]bool) map[string]bool {
	s_union := map[string]bool{}
	for k := range s1 {
		s_union[k] = true
	}
	for k := range s2 {
		s_union[k] = true
	}
	return s_union
}

func verticalSet(pattern, text string) bool {
	var any func(func(map[string]bool) bool, []map[string]bool) bool
	any = func(done func(map[string]bool) bool, accs []map[string]bool) bool {
		if len(accs) == 0 {
			return false
		}
		acc, tail := accs[0], accs[1:]
		return done(acc) || any(done, tail)
	}

	var scanl func(func(map[string]bool, string) map[string]bool, map[string]bool, []string) []map[string]bool
	scanl = func(step func(map[string]bool, string) map[string]bool, acc map[string]bool, slice []string) []map[string]bool {
		if len(slice) == 0 {
			return []map[string]bool{acc}
		}
		x, xt := slice[0], slice[1:]
		return append([]map[string]bool{acc}, scanl(step, step(acc, x), xt)...)
	}
	init := map[string]bool{pattern: true}
	step := func(acc map[string]bool, x string) map[string]bool {
		candidates := make(map[string]bool)
		for c := range acc {
			if strings.HasPrefix(c, x) {
				candidates[c[1:]] = true
			}
		}
		return union(init, candidates)
	}
	done := func(acc map[string]bool) bool {
		_, ok := acc[""]
		return ok
	}
	return any(done, scanl(step, init, strings.Split(text, "")))
}
