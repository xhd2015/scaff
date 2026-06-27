package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: git-hooks <install|pre-commit|pre-push>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "install":
		fmt.Println("install hooks")
	case "pre-commit", "pre-push":
		// no-op stub
	default:
		fmt.Fprintln(os.Stderr, "unknown command")
		os.Exit(2)
	}
}