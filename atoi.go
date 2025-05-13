package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseIntSlice parses and returns the []int slice represented by the string s.
func ParseIntSlice(s string, base int) ([]int, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []int
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseInt(strings.TrimSpace(v), base, 64)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, int(value))
	}

	return theSlice, nil
}

// ParseInt8Slice parses and returns the []int8 slice represented by the string s.
func ParseInt8Slice(s string, base int) ([]int8, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []int8
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseInt(strings.TrimSpace(v), base, 8)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, int8(value))
	}

	return theSlice, nil
}

// ParseInt16Slice parses and returns the []int16 slice represented by the string s.
func ParseInt16Slice(s string, base int) ([]int16, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []int16
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseInt(strings.TrimSpace(v), base, 16)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, int16(value))
	}

	return theSlice, nil
}

// ParseInt32Slice parses and returns the []int32 slice represented by the string s.
func ParseInt32Slice(s string, base int) ([]int32, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []int32
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseInt(strings.TrimSpace(v), base, 32)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, int32(value))
	}

	return theSlice, nil
}

// ParseInt64Slice parses and returns the []int64 slice represented by the string s.
func ParseInt64Slice(s string, base int) ([]int64, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []int64
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseInt(strings.TrimSpace(v), base, 64)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, value)
	}

	return theSlice, nil
}

// ParseUintSlice parses and returns the []uint slice represented by the string s.
func ParseUintSlice(s string, base int) ([]uint, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []uint
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseUint(strings.TrimSpace(v), base, 64)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, uint(value))
	}

	return theSlice, nil
}

// ParseUint8Slice parses and returns the []uint8 slice represented by the string s.
func ParseUint8Slice(s string, base int) ([]uint8, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []uint8
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseUint(strings.TrimSpace(v), base, 8)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, uint8(value))
	}

	return theSlice, nil
}

// ParseUint16Slice parses and returns the []uint16 slice represented by the string s.
func ParseUint16Slice(s string, base int) ([]uint16, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []uint16
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseUint(strings.TrimSpace(v), base, 16)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, uint16(value))
	}

	return theSlice, nil
}

// ParseUint32Slice parses and returns the []uint32 slice represented by the string s.
func ParseUint32Slice(s string, base int) ([]uint32, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []uint32
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseUint(strings.TrimSpace(v), base, 32)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, uint32(value))
	}

	return theSlice, nil
}

// ParseUint64Slice parses and returns the []uint64 slice represented by the string s.
func ParseUint64Slice(s string, base int) ([]uint64, error) {
	values, err := parseStringSlice(s)
	if err != nil {
		return nil, err
	}

	var theSlice []uint64
	for _, v := range values {
		if v == "" {
			continue
		}

		value, err := strconv.ParseUint(strings.TrimSpace(v), base, 64)
		if err != nil {
			return nil, err
		}
		theSlice = append(theSlice, value)
	}

	return theSlice, nil
}

// Atoi is equivalent to ParseIntSlice(s, 10).
func Atoi(s string) ([]int, error) {
	return ParseIntSlice(s, 10)
}
