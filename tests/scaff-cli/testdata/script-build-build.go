// usage: go run ./script/build
package main

import (
	"fmt"
	"os"

	"github.com/xhd2015/xgo/support/cmd"
)

func main() {
	if err := Handle(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Handle(args []string) error {
	fmt.Println("==> Building")
	return cmd.Debug().Run("go", "build", "-o", "bin/app", ".")
}