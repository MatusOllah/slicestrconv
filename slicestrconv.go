package slicestrconv

import (
	"errors"
	"strings"
)

var (
	ErrInvalidSyntax error = errors.New("invalid slice syntax")
)

var (
	OpeningBracket string = "["
	ClosingBracket string = "]"
	Delimiter      string = ","
)

func parseStringSlice(s string) ([]string, error) {
	s = strings.TrimSpace(s)

	if !strings.HasPrefix(s, OpeningBracket) || !strings.HasSuffix(s, ClosingBracket) {
		return nil, ErrInvalidSyntax
	}

	s = strings.TrimPrefix(s, OpeningBracket)
	s = strings.TrimSuffix(s, ClosingBracket)

	values := strings.Split(s, Delimiter)
	return values, nil
}
