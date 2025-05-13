package slicestrconv

import (
	"strconv"
	"strings"
)

// ParseBoolSlice parses and returns the []bool slice represented by the string s.
func ParseBoolSlice(s string) ([]bool, error) {
	if !strings.HasPrefix(s, OpeningBracket) || !strings.HasSuffix(s, ClosingBracket) {
		return nil, ErrInvalidSyntax
	}

	s = strings.TrimPrefix(s, OpeningBracket)
	s = strings.TrimSuffix(s, ClosingBracket)

	values := strings.Split(s, Delimiter)

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
