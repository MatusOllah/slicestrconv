package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseComplex64Slice parses and returns the []complex64 slice represented by the string s.
func ParseComplex64Slice(s string, base int) ([]complex64, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

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
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

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

// Atoc is equivalent to ParseComplex128Slice(s, 10).
func Atoc(s string) ([]complex128, error) {
	return ParseComplex128Slice(s, 10)
}
