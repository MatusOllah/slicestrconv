package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseFloat32Slice parses and returns the []float32 slice represented by the string s.
func ParseFloat32Slice(s string, base int) ([]float32, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []float32
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseFloat(strings.TrimSpace(v), base)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, float32(value))
	}

	return theSlice, nil
}

// ParseFloat64Slice parses and returns the []float64 slice represented by the string s.
func ParseFloat64Slice(s string, base int) ([]float64, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []float64
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseFloat(strings.TrimSpace(v), base)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, value)
	}

	return theSlice, nil
}
