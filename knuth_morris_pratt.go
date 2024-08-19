package kmp

import (
	"strings"
)

func knuth_morris_pratt(pattern, text string) bool {

	var step func(*tree, string) *tree

	check := func(candidate *tree, x string) bool {
		if candidate == nil {
			return false
		}
		return strings.HasPrefix(candidate.top, x)
	}

	var make func([]string, *tree) *tree
	make = func(slice []string, partial *tree) *tree {
		if len(slice) == 0 {
			return &tree{top: "", next: func() *tree { return nil }, rest: partial}
		}
		n := func() *tree { return make(slice[1:], step(partial, slice[0])) }
		var r_star *tree
		if check(partial, slice[0]) {
			r_star = partial.rest
		} else {
			r_star = partial
		}
		return &tree{top: slice[0], next: n, rest: r_star}
	}

	var any func(func(*tree) bool, []*tree) bool
	any = func(done func(*tree) bool, accs []*tree) bool {
		if len(accs) == 0 {
			return false
		}
		acc, tail := accs[0], accs[1:]
		return done(acc) || any(done, tail)
	}

	var scanl func(func(*tree, string) *tree, *tree, []string) []*tree
	scanl = func(step func(*tree, string) *tree, acc *tree, slice []string) []*tree {
		if len(slice) == 0 {
			return []*tree{acc}
		}
		x, xt := slice[0], slice[1:]
		return append([]*tree{acc}, scanl(step, step(acc, x), xt)...)
	}

	init := make(strings.Split(pattern, ""), nil)

	step = func(acc *tree, x string) *tree {
		if acc == nil {
			return init
		}
		if check(acc, x) {
			return acc.next()
		}
		return step(acc.rest, x)
	}

	done := func(acc *tree) bool {
		if acc == nil {
			return false
		}
		return acc.top == ""
	}

	return any(done, scanl(step, init, strings.Split(text, "")))
}
