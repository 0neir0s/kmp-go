package kmp

type tree struct {
	top  string
	next func() *tree
	rest *tree
}
