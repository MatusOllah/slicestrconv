# slicestrconv

[![Go Reference](https://pkg.go.dev/badge/github.com/MatusOllah/slicestrconv.svg)](https://pkg.go.dev/github.com/MatusOllah/slicestrconv) [![Go Report Card](https://goreportcard.com/badge/github.com/MatusOllah/slicestrconv)](https://goreportcard.com/report/github.com/MatusOllah/slicestrconv)

**slicestrconv** is a string to slice converting / parsing library for Go. Like strconv but for slices.

## Basic Usage

```go
package main

import (
    "fmt"

    "github.com/MatusOllah/slicestrconv"
)

func main() {
    boolSlice, err := slicestrconv.ParseBoolSlice("[true, false, true]")
    if err != nil {
        panic(err)
    }

    fmt.Println(boolSlice) // [true false true]

    intSlice, err := slicestrconv.ParseIntSlice("[1, 2, 3, 420, 69]", 10)
    if err != nil {
        panic(err)
    }

    fmt.Println(intSlice) // [1 2 3 420 69]

    floatSlice, err := slicestrconv.ParseFloatSlice("[1.1, 2.2, 3.3, 3.14]", 10)
    if err != nil {
        panic(err)
    }

    fmt.Println(floatSlice) // [1.1 2.2 3.3 3.14]
}
```
