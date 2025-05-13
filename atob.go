package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseBoolSlice parses and returns the []bool slice represented by the string s.
func ParseBoolSlice(s string) ([]bool, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []bool
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseBool(strings.TrimSpace(v))
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, value)
	}

	return theSlice, nil
}

// Atob is equivalent to ParseBoolSlice(s).
func Atob(s string) ([]bool, error) {
	return ParseBoolSlice(s)
}
