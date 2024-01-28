package slicestrconv_test

import (
	"testing"

	. "github.com/MatusOllah/slicestrconv"
)

func TestParseBoolSlice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []bool
		expectedError  error
	}{
		{"[true]", []bool{true}, nil},
		{"[false]", []bool{false}, nil},
		{"[true, false]", []bool{true, false}, nil},
		{"[false, true]", []bool{false, true}, nil},
		{"[true, false, true]", []bool{true, false, true}, nil},
		{"[false, true, false]", []bool{false, true, false}, nil},
		{"[true,false]", []bool{true, false}, nil},
		{"[false,true]", []bool{false, true}, nil},
		{"[true,false,true]", []bool{true, false, true}, nil},
		{"[false,true,false]", []bool{false, true, false}, nil},
		{"[true,    false]", []bool{true, false}, nil},
		{"[false,    true]", []bool{false, true}, nil},
		{"[true,    false,    true]", []bool{true, false, true}, nil},
		{"[false,    true,    false]", []bool{false, true, false}, nil},
		{"true]", nil, ErrInvalidSyntax},
		{"false]", nil, ErrInvalidSyntax},
		{"true, false]", nil, ErrInvalidSyntax},
		{"false, true]", nil, ErrInvalidSyntax},
		{"true, false, true]", nil, ErrInvalidSyntax},
		{"false, true, false]", nil, ErrInvalidSyntax},
		{"[true", nil, ErrInvalidSyntax},
		{"[false", nil, ErrInvalidSyntax},
		{"[true, false", nil, ErrInvalidSyntax},
		{"[false, true", nil, ErrInvalidSyntax},
		{"[true, false, true", nil, ErrInvalidSyntax},
		{"[false, true, false", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseBoolSlice(tc.s)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseIntSlice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []int
		expectedError  error
	}{
		{"[69]", []int{69}, nil},
		{"[-69]", []int{-69}, nil},
		{"[1, 2]", []int{1, 2}, nil},
		{"[1, 2, 3]", []int{1, 2, 3}, nil},
		{"[1,2]", []int{1, 2}, nil},
		{"[1,2,3]", []int{1, 2, 3}, nil},
		{"[1,    2]", []int{1, 2}, nil},
		{"[1,    2,    3]", []int{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseIntSlice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseInt8Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []int8
		expectedError  error
	}{
		{"[69]", []int8{69}, nil},
		{"[-69]", []int8{-69}, nil},
		{"[1, 2]", []int8{1, 2}, nil},
		{"[1, 2, 3]", []int8{1, 2, 3}, nil},
		{"[1,2]", []int8{1, 2}, nil},
		{"[1,2,3]", []int8{1, 2, 3}, nil},
		{"[1,    2]", []int8{1, 2}, nil},
		{"[1,    2,    3]", []int8{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseInt8Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseInt16Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []int16
		expectedError  error
	}{
		{"[69]", []int16{69}, nil},
		{"[-69]", []int16{-69}, nil},
		{"[1, 2]", []int16{1, 2}, nil},
		{"[1, 2, 3]", []int16{1, 2, 3}, nil},
		{"[1,2]", []int16{1, 2}, nil},
		{"[1,2,3]", []int16{1, 2, 3}, nil},
		{"[1,    2]", []int16{1, 2}, nil},
		{"[1,    2,    3]", []int16{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseInt16Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseInt32Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []int32
		expectedError  error
	}{
		{"[69]", []int32{69}, nil},
		{"[-69]", []int32{-69}, nil},
		{"[1, 2]", []int32{1, 2}, nil},
		{"[1, 2, 3]", []int32{1, 2, 3}, nil},
		{"[1,2]", []int32{1, 2}, nil},
		{"[1,2,3]", []int32{1, 2, 3}, nil},
		{"[1,    2]", []int32{1, 2}, nil},
		{"[1,    2,    3]", []int32{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseInt32Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseInt64Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []int64
		expectedError  error
	}{
		{"[69]", []int64{69}, nil},
		{"[-69]", []int64{-69}, nil},
		{"[1, 2]", []int64{1, 2}, nil},
		{"[1, 2, 3]", []int64{1, 2, 3}, nil},
		{"[1,2]", []int64{1, 2}, nil},
		{"[1,2,3]", []int64{1, 2, 3}, nil},
		{"[1,    2]", []int64{1, 2}, nil},
		{"[1,    2,    3]", []int64{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseInt64Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseUintSlice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []uint
		expectedError  error
	}{
		{"[69]", []uint{69}, nil},
		{"[1, 2]", []uint{1, 2}, nil},
		{"[1, 2, 3]", []uint{1, 2, 3}, nil},
		{"[1,2]", []uint{1, 2}, nil},
		{"[1,2,3]", []uint{1, 2, 3}, nil},
		{"[1,    2]", []uint{1, 2}, nil},
		{"[1,    2,    3]", []uint{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseUintSlice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseUint8Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []uint8
		expectedError  error
	}{
		{"[69]", []uint8{69}, nil},
		{"[1, 2]", []uint8{1, 2}, nil},
		{"[1, 2, 3]", []uint8{1, 2, 3}, nil},
		{"[1,2]", []uint8{1, 2}, nil},
		{"[1,2,3]", []uint8{1, 2, 3}, nil},
		{"[1,    2]", []uint8{1, 2}, nil},
		{"[1,    2,    3]", []uint8{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseUint8Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseUint16Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []uint16
		expectedError  error
	}{
		{"[69]", []uint16{69}, nil},
		{"[1, 2]", []uint16{1, 2}, nil},
		{"[1, 2, 3]", []uint16{1, 2, 3}, nil},
		{"[1,2]", []uint16{1, 2}, nil},
		{"[1,2,3]", []uint16{1, 2, 3}, nil},
		{"[1,    2]", []uint16{1, 2}, nil},
		{"[1,    2,    3]", []uint16{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseUint16Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseUint32Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []uint32
		expectedError  error
	}{
		{"[69]", []uint32{69}, nil},
		{"[1, 2]", []uint32{1, 2}, nil},
		{"[1, 2, 3]", []uint32{1, 2, 3}, nil},
		{"[1,2]", []uint32{1, 2}, nil},
		{"[1,2,3]", []uint32{1, 2, 3}, nil},
		{"[1,    2]", []uint32{1, 2}, nil},
		{"[1,    2,    3]", []uint32{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseUint32Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseUint64Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []uint64
		expectedError  error
	}{
		{"[69]", []uint64{69}, nil},
		{"[1, 2]", []uint64{1, 2}, nil},
		{"[1, 2, 3]", []uint64{1, 2, 3}, nil},
		{"[1,2]", []uint64{1, 2}, nil},
		{"[1,2,3]", []uint64{1, 2, 3}, nil},
		{"[1,    2]", []uint64{1, 2}, nil},
		{"[1,    2,    3]", []uint64{1, 2, 3}, nil},
		{"1, 2]", nil, ErrInvalidSyntax},
		{"1, 2, 3]", nil, ErrInvalidSyntax},
		{"[1, 2", nil, ErrInvalidSyntax},
		{"[1, 2, 3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseUint64Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseFloat32Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []float32
		expectedError  error
	}{
		{"[3.1415926535]", []float32{3.1415926535}, nil},
		{"[1.1, 2.2]", []float32{1.1, 2.2}, nil},
		{"[1.1, 2.2, 3.3]", []float32{1.1, 2.2, 3.3}, nil},
		{"[1.1,2.2]", []float32{1.1, 2.2}, nil},
		{"[1.1,2.2,3.3]", []float32{1.1, 2.2, 3.3}, nil},
		{"[1.1,    2.2]", []float32{1.1, 2.2}, nil},
		{"[1.1,    2.2,    3.3]", []float32{1.1, 2.2, 3.3}, nil},
		{"1.1, 2.2]", nil, ErrInvalidSyntax},
		{"1.1, 2.2, 3.3]", nil, ErrInvalidSyntax},
		{"[1.1, 2.2", nil, ErrInvalidSyntax},
		{"[1.1, 2.2, 3.3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseFloat32Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseFloat64Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []float64
		expectedError  error
	}{
		{"[3.1415926535]", []float64{3.1415926535}, nil},
		{"[1.1, 2.2]", []float64{1.1, 2.2}, nil},
		{"[1.1, 2.2, 3.3]", []float64{1.1, 2.2, 3.3}, nil},
		{"[1.1,2.2]", []float64{1.1, 2.2}, nil},
		{"[1.1,2.2,3.3]", []float64{1.1, 2.2, 3.3}, nil},
		{"[1.1,    2.2]", []float64{1.1, 2.2}, nil},
		{"[1.1,    2.2,    3.3]", []float64{1.1, 2.2, 3.3}, nil},
		{"1.1, 2.2]", nil, ErrInvalidSyntax},
		{"1.1, 2.2, 3.3]", nil, ErrInvalidSyntax},
		{"[1.1, 2.2", nil, ErrInvalidSyntax},
		{"[1.1, 2.2, 3.3", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseFloat64Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseComplex64Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []complex64
		expectedError  error
	}{
		{"[69i]", []complex64{69i}, nil},
		{"[1.1i, 2.2i]", []complex64{1.1i, 2.2i}, nil},
		{"[1.1i, 2.2i, 3.3i]", []complex64{1.1i, 2.2i, 3.3i}, nil},
		{"[1.1i,2.2i]", []complex64{1.1i, 2.2i}, nil},
		{"[1.1i,2.2i,3.3i]", []complex64{1.1i, 2.2i, 3.3i}, nil},
		{"[1.1i,    2.2i]", []complex64{1.1i, 2.2i}, nil},
		{"[1.1i,    2.2i,    3.3i]", []complex64{1.1i, 2.2i, 3.3i}, nil},
		{"1.1i, 2.2i]", nil, ErrInvalidSyntax},
		{"1.1i, 2.2i, 3.3i]", nil, ErrInvalidSyntax},
		{"[1.1i, 2.2i", nil, ErrInvalidSyntax},
		{"[1.1i, 2.2i, 3.3i", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseComplex64Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}

func TestParseComplex128Slice(t *testing.T) {
	OpeningBracket = "["
	ClosingBracket = "]"
	Delimiter = ","

	testcases := []struct {
		s              string
		expectedResult []complex128
		expectedError  error
	}{
		{"[69i]", []complex128{69i}, nil},
		{"[1.1i, 2.2i]", []complex128{1.1i, 2.2i}, nil},
		{"[1.1i, 2.2i, 3.3i]", []complex128{1.1i, 2.2i, 3.3i}, nil},
		{"[1.1i,2.2i]", []complex128{1.1i, 2.2i}, nil},
		{"[1.1i,2.2i,3.3i]", []complex128{1.1i, 2.2i, 3.3i}, nil},
		{"[1.1i,    2.2i]", []complex128{1.1i, 2.2i}, nil},
		{"[1.1i,    2.2i,    3.3i]", []complex128{1.1i, 2.2i, 3.3i}, nil},
		{"1.1i, 2.2i]", nil, ErrInvalidSyntax},
		{"1.1i, 2.2i, 3.3i]", nil, ErrInvalidSyntax},
		{"[1.1i, 2.2i", nil, ErrInvalidSyntax},
		{"[1.1i, 2.2i, 3.3i", nil, ErrInvalidSyntax},
	}

	for _, tc := range testcases {
		result, err := ParseComplex128Slice(tc.s, 10)
		if err != tc.expectedError {
			t.Fatalf("expected error %v, got %v", tc.expectedError, err)
		}
		if len(result) != len(tc.expectedResult) {
			t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
		}
		for i := range result {
			if result[i] != tc.expectedResult[i] {
				t.Fatalf("expected result %v, got %v", tc.expectedResult, result)
			}
		}
	}
}
