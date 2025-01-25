# Bengali Number Converter (banglaconv)
[![GoDoc](https://godoc.org/github.com/sahapranta/banglaconv?status.svg)](https://godoc.org/github.com/sahapranta/banglaconv)
[![Go Report Card](https://goreportcard.com/badge/github.com/sahapranta/banglaconv)](https://goreportcard.com/report/github.com/sahapranta/banglaconv)
![Build workflow](https://github.com/sahapranta/banglaconv/actions/workflows/go.yml/badge.svg)
[![codecov](https://codecov.io/gh/sahapranta/banglaconv/branch/main/graph/badge.svg)](https://codecov.io/gh/sahapranta/banglaconv)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A Go library for converting numbers to Bengali words and Bengali numerals.

## Features

- Convert English numerals to Bengali numerals
- Convert numbers to Bengali word representation
- Supports integers and floating-point numbers
- Handles numbers up to crores (10,000,000)

## Installation

```bash
go get github.com/sahapranta/banglaconv
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/sahapranta/banglaconv"
)

func main() {
    converter := &bengaliconv.NumConverter{}

    // Convert to Bengali numerals
    bengaliNum := converter.ToBengaliNumber(1234)
    fmt.Println(bengaliNum)  // Output: ১২৩৪

    // Convert to Bengali words
    bengaliWord, _ := converter.ToBengaliWord(1234567)
    fmt.Println(bengaliWord)  // Output: বার লক্ষ চৌত্রিশ হাজার পাঁচশত সাতষট্টি

    // Works with floating-point numbers
    floatWord, _ := converter.ToBengaliWord(1234.56)
    fmt.Println(floatWord)  // Output: এক হাজার দুইশত চৌত্রিশ দশমিক পাঁচ ছয়
}
```

## Documentation
Visit the [GoDoc](https://godoc.org/github.com/sahapranta/banglaconv) page for the full documentation.

## Testing

To run tests:

```bash
go test ./...
```

To run benchmarks:

```bash
go test -bench=.
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](https://github.com/sahapranta)