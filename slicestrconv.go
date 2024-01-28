package slicestrconv

import (
	"errors"
)

var (
	ErrInvalidSyntax error = errors.New("invalid slice syntax")
)

var (
	OpeningBracket string = "["
	ClosingBracket string = "]"
	Delimiter      string = ","
)
