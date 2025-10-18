//go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && filepath.Ext(path) == ".proto" {
			fmt.Println(path)
		}
		return nil
	})
}
