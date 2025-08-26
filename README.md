# densemap

[![Go Reference](https://pkg.go.dev/badge/github.com/jokruger/densemap.svg)](https://pkg.go.dev/github.com/jokruger/densemap)
[![Go Report Card](https://goreportcard.com/badge/github.com/jokruger/densemap)](https://goreportcard.com/report/github.com/jokruger/densemap)
[![codecov](https://codecov.io/gh/jokruger/densemap/graph/badge.svg?token=Q136FDHBFJ)](https://codecov.io/gh/jokruger/densemap)

Package densemap provides a generic, dense, ID-based mapping structure for fast, contiguous lookups by integer ID. Unlike Go’s built-in map, densemap stores values in a slice indexed by a bounded integer range [minID, maxID]. This design ensures O(1) access, efficient iteration, and predictable memory usage, making it well-suited for cases where IDs are compact and naturally bounded (e.g., enums, small identifier spaces, event codes).

## Install

Run `go get github.com/jokruger/densemap`

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
