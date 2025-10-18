# Cantonese Romanisation (Go)
[![Go CI](https://github.com/chunlaw/cantonese-romanisation-go/actions/workflows/ci.yml/badge.svg)](https://github.com/chunlaw/cantonese-romanisation-go/actions/workflows/ci.yml)

Port of [`chunlaw/cantonese-romanisation`](https://github.com/chunlaw/cantonese-romanisation) providing a Go API for
looking up Cantonese pronunciations for individual Han characters.

The embedded dictionary exposes three romanisation systems:

- `roman` — raw roman column from the upstream dataset
- `lshk` — Linguistic Society of Hong Kong (Jyutping)
- `yale` — Yale romanisation

## Installation

```bash
go get github.com/AddIcechan/cantonese-romanisation-go
```

## Usage

```go
package main

import (
	"fmt"

	cantonese "github.com/AddIcechan/cantonese-romanisation-go"
)

func main() {
	pronunciations := cantonese.Romanize("香港", cantonese.SchemeLSHK)
	for _, item := range pronunciations {
		fmt.Printf("%s -> %v\n", item.Character, item.Pronunciations)
	}

	// Convenience helper returning the first pronunciation for each rune.
	fmt.Println(cantonese.RomanizeFirst("香港", cantonese.SchemeLSHK))
}
```

## Development

```bash
go test ./...
```

Pull requests are checked by the Go CI workflow, which enforces `gofmt` formatting and runs the full unit test suite.
