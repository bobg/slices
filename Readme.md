# Slices - generic utility functions for operating on slices

[![Go Reference](https://pkg.go.dev/badge/github.com/bobg/slices.svg)](https://pkg.go.dev/github.com/bobg/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/bobg/slices)](https://goreportcard.com/report/github.com/bobg/slices)
[![Tests](https://github.com/bobg/slices/actions/workflows/go.yml/badge.svg)](https://github.com/bobg/slices/actions/workflows/go.yml)
[![Coverage Status](https://coveralls.io/repos/github/bobg/slices/badge.svg?branch=master)](https://coveralls.io/github/bobg/slices?branch=master)

This is slices,
a collection of generic utility functions for operating on Go slices.

This collection is useful primarily as a way to encapsulate the Go idioms for inserting and removing slice elements,
which can be hard to remember.

Secondarily it introduces negative-integer indexing.
`Get(s, -1)` on a slice `s` is the same as `s[len(s)-1]`.
