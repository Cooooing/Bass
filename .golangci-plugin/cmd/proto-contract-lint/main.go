package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cooooing/bass-lint-plugin/internal/protocontract"
)

func main() {
	root := flag.String("root", "", "Proto root directory")
	flag.Parse()
	if *root == "" {
		fmt.Fprintln(os.Stderr, "proto-contract-lint: --root is required")
		os.Exit(2)
	}
	violations, err := protocontract.Check(*root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "proto-contract-lint:", err)
		os.Exit(2)
	}
	for _, violation := range violations {
		fmt.Printf("%s:%d: %s: %s\n", violation.Path, violation.Line, violation.Rule, violation.Message)
	}
	if len(violations) > 0 {
		os.Exit(1)
	}
}
