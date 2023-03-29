package genetic

import (
	"fmt"
)

// insert inserts element a at index i in slice l and returns the resulting slice.
// If i < 0 or i > len(l), then err is an error and s=l.
func insert[A any](l []A, a A, i int) (s []A, err error) {
	if i < 0 || i > len(l) {
		return l, fmt.Errorf("index out of bounds: i=%d is not within the valid range [0,len(l)] = [0,%d]", i, len(l))
	}
	l = append(l, a)
	copy(l[i+1:], l[i:])
	l[i] = a
	return l, nil
}

// remove removes the element at index i from slice l and returns the resulting slice together
// with the element that was removed.
// If i < 0 or i >= len(l), then err is an error, s=l, and elem is the zero value.
func remove[A any](l []A, i int) (s []A, elem A, err error) {
	if i < 0 || i >= len(l) {
		var a A
		return l, a, fmt.Errorf("index out of bounds: i=%d is not within the valid range [0,len(l)) = [0,%d)", i, len(l))
	}
	elem = l[i]
	return append(l[:i], l[i+1:]...), elem, nil
}

// indexOf returns the index of the first occurrence of a in slice l, or -1 if a is not found.
func indexOf[A comparable](l []A, a A) int {
	for i, v := range l {
		if v == a {
			return i
		}
	}
	return -1
}
