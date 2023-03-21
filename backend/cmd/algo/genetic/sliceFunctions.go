package genetic

import (
	"errors"
	"fmt"
)

// insert inserts element a at index i in slice l and returns the resulting slice.
// If 0 <= i <= len(l) then err==nil, otherwise err is an error
func insert[A interface{}](l []A, a A, i int) (slice []A, err error) {
	if i < 0 || i > len(l) {
		return l, errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(l)] = [0,%d]", i, len(l)))
	}
	l = append(l, a)
	copy(l[i+1:], l[i:])
	l[i] = a
	return l, nil
}

// remove removes the element at index i from slice l and returns the resulting slice together with the element that was removed.
// If 0 <= i < len(l) then err==nil, otherwise err is an error
func remove[A interface{}](l []A, i int) (s []A, elem A, err error) {
	if i < 0 || i >= len(l) {
		var a A
		return l, a, errors.New(fmt.Sprintf("index out of bounds: i=%d not in range [0,len(l)) = [0,%d)", i, len(l)))
	}
	elem = l[i]
	return append(l[:i], l[i+1:]...), elem, nil
}

// indexOf returns the index of a in l. If a is not in l, -1 is returned
func indexOf[A comparable](l []A, a A) int {
	for i, v := range l {
		if v == a {
			return i
		}
	}
	return -1
}
