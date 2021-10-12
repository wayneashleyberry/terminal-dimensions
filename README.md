[![GoDoc](https://godoc.org/github.com/wayneashleyberry/terminal-dimensions?status.svg)](https://godoc.org/github.com/wayneashleyberry/terminal-dimensions)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/terminal-dimensions)](https://goreportcard.com/report/github.com/wayneashleyberry/terminal-dimensions)
[![GolangCI](https://golangci.com/badges/github.com/wayneashleyberry/terminal-dimensions.svg)](https://golangci.com/r/github.com/wayneashleyberry/terminal-dimensions)

```sh
go get github.com/wayneashleyberry/terminal-dimensions
```

```go
package main

import (
	"fmt"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

func main() {
	x, _ := terminal.Width()
	y, _ := terminal.Height()
	fmt.Printf("Terminal is %d wide and %d high", x, y)
}
```
