package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "."
	sep := string(os.PathSeparator)

	root, _ = filepath.Abs(root)

	rl := len(strings.Split(root, sep))

	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			fmt.Printf("%s\n", root)
		} else {
			l := len(strings.Split(path, sep))
			l -= rl
			s := ""
			if l != 1 {
				s = strings.Repeat("  ", l-1)
			}
			p := strings.Replace(path, root, "", -1)
			bp := filepath.Base(p)
			fmt.Printf("%s%s\n", s, bp)
		}
		return nil
	})
}
