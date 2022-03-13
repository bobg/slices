// Package slices contains utility functions for working with slices.
// It adds the ability to index from the right end of a slice using negative integers.
package slices

// Get gets the idx'th element of s.
//
// If idx < 0 it counts from the end of s.
func Get[T any](s []T, idx int) T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx]
}

// Put puts a given value into the idx'th location in s.
//
// If idx < 0 it counts from the end of s.
//
// The input slice is modified.
func Put[T any](s []T, idx int, val T) {
	if idx < 0 {
		idx += len(s)
	}
	s[idx] = val
}

// Append is the same as Go's builtin append and is included for completeness.
func Append[T any](s []T, vals ...T) []T {
	return append(s, vals...)
}

// Insert inserts the given values at the idx'th location in s and returns the result.
// After the insert, the first new value has position idx.
// If idx < 0, it counts from the end of s.
//
// The input slice is modified.
//
// Example: Insert([x, y, z], 1, a, b, c) -> [x, a, b, c, y, z]
func Insert[T any](s []T, idx int, vals ...T) []T {
	if idx < 0 {
		idx += len(s)
	}
	return insert(s, idx, vals...)
}

func insert[T any](s []T, idx int, vals ...T) []T {
	// Make s long enough.
	s = append(s, vals...)

	// Make space in s at the right position.
	copy(s[idx+len(vals):], s[idx:])

	// Put values in the right spot.
	copy(s[idx:], vals)

	return s
}

// ReplaceN replaces the n values of s beginning at position idx with the given values.
// After the replace, the first new value has position idx.
// If idx < 0, it counts from the end of s.
//
// The input slice is modified.
func ReplaceN[T any](s []T, idx, n int, vals ...T) []T {
	if idx < 0 {
		idx += len(s)
	}
	return replaceN(s, idx, n, vals...)
}

func replaceN[T any](s []T, idx, n int, vals ...T) []T {
	if n > len(vals) {
		// Removing more items than inserting.
		s = removeN(s, idx, n-len(vals))
	} else if n < len(vals) {
		// Inserting more items than removing.
		delta := len(vals) - n
		s = insert(s, idx, vals[:delta]...)
		idx += delta
		vals = vals[delta:]
	}
	copy(s[idx:], vals)

	return s
}

// ReplaceTo replaces the values of s beginning at from and ending before to with the given values.
// After the replace, the first new value has position from.
//
// If from < 0 it counts from the end of s.
// If to < 0 it counts from the end of s.
// If to == 0, that means len(s).
//
// The input slice is modified.
func ReplaceTo[T any](s []T, from, to int, vals ...T) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return replaceN(s, from, to-from, vals...)
}

// RemoveN removes n items from s beginning at position idx and returns the result.
//
// If idx < 0 it counts from the end of s.
//
// The input slice is modified.
//
// Example: RemoveN([a, b, c, d], 1, 2) -> [a, d]
func RemoveN[T any](s []T, idx, n int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return removeN(s, idx, n)
}

func removeN[T any](s []T, idx, n int) []T {
	copy(s[idx:], s[idx+n:])
	newlen := len(s) - n
	return s[:newlen]
}

// RemoveTo removes items from s beginning at position from and ending before position to.
// It returns the result.
//
// If from < 0 it counts from the end of s.
// If to < 0 it counts from the end of s.
// If to == 0, that means len(s).
//
// The input slice is modified.
//
// Example: RemoveTo([a, b, c, d], 1, 3) -> [a, d]
func RemoveTo[T any](s []T, from, to int) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return removeN(s, from, to-from)
}

// Prefix returns s up to but not including position idx.
//
// If idx < 0 it counts from the end of s.
func Prefix[T any](s []T, idx int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[:idx]
}

// Suffix returns s excluding elements before position idx.
//
// If idx < 0 it counts from the end of s.
func Suffix[T any](s []T, idx int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx:]
}

// SliceN returns n elements of s beginning at position idx.
//
// If idx < 0 it counts from the end of s.
func SliceN[T any](s []T, idx, n int) []T {
	if idx < 0 {
		idx += len(s)
	}
	return s[idx : idx+n]
}

// SliceTo returns the elements of s beginning at position from and ending before position to.
//
// If from < 0 it counts from the end of s.
// If to < 0 it counts from the end of s.
// If to == 0, that means len(s).
func SliceTo[T any](s []T, from, to int) []T {
	if from < 0 {
		from += len(s)
	}
	if to < 0 {
		to += len(s)
	} else if to == 0 {
		to = len(s)
	}
	return s[from:to]
}
