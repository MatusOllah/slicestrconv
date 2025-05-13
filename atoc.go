package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseComplex64Slice parses and returns the []complex64 slice represented by the string s.
func ParseComplex64Slice(s string, base int) ([]complex64, error) {
	if !strings.HasPrefix(s, OpeningBracket) || !strings.HasSuffix(s, ClosingBracket) {
		return nil, ErrInvalidSyntax
	}

	s = strings.TrimPrefix(s, OpeningBracket)
	s = strings.TrimSuffix(s, ClosingBracket)

	values := strings.Split(s, Delimiter)

	var theSlice []complex64
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseComplex(strings.TrimSpace(v), base)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, complex64(value))
	}

	return theSlice, nil
}

// ParseComplex128Slice parses and returns the []complex128 slice represented by the string s.
func ParseComplex128Slice(s string, base int) ([]complex128, error) {
	if !strings.HasPrefix(s, OpeningBracket) || !strings.HasSuffix(s, ClosingBracket) {
		return nil, ErrInvalidSyntax
	}

	s = strings.TrimPrefix(s, OpeningBracket)
	s = strings.TrimSuffix(s, ClosingBracket)

	values := strings.Split(s, Delimiter)

	var theSlice []complex128
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseComplex(strings.TrimSpace(v), base)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, value)
	}

	return theSlice, nil
}
