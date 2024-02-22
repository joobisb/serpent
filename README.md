# serpent

[![Go Reference](https://pkg.go.dev/badge/github.com/coder/serpent.svg)](https://pkg.go.dev/github.com/coder/serpent)

`serpent` is a Go CLI configuration framework based on [cobra](https://github.com/spf13/cobra) and used by [coder/coder](https://github.com/coder/coder).
It's designed for large-scale CLIs with dozens of commands and hundreds
of options. If you're building a small, self-contained tool, go with
cobra.

## Basic Usage

See `example/echo`:

```go
package main

import (
	"os"
	"strings"

	"github.com/coder/serpent"
)

func main() {
	var upper bool
	cmd := serpent.Cmd{
		Use:   "echo <text>",
		Short: "Prints the given text to the console.",
		Options: serpent.OptionSet{
			{
				Name:        "upper",
				Value:       serpent.BoolOf(&upper),
				Flag:        "upper",
				Description: "Prints the text in upper case.",
			},
		},
		Handler: func(inv *serpent.Invocation) error {
			if len(inv.Args) == 0 {
				inv.Stderr.Write([]byte("error: missing text\n"))
				os.Exit(1)
			}

			text := inv.Args[0]
			if upper {
				text = strings.ToUpper(text)
			}

			inv.Stdout.Write([]byte(text))
			return nil
		},
	}

	err := cmd.Invoke().WithOS().Run()
	if err != nil {
		panic(err)
	}
}
```

## Design
This Design section assumes you have a good understanding of how `cobra` works.

### Options

Serpent is designed for high-configurability. To us, that means providing
many ways to configure the same value (env, YAML, flags, etc.) and keeping
the code clean and testable as you scale the number of options.

## More coming...
This README is a stub for now. We'll better explain the design and usage
of `serpent` in the future.