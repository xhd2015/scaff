package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptDevPath = "script/dev/main.go"

const scriptDevStub = `// usage: go run ./script/dev [args...]
//
// Proposed behavior (sketch):
//   1. Forward optional args to the module main.
//   2. Run go run . --dev with those args.
//   3. Exit non-zero on failure.
package main

import (
	"fmt"
	"os"

	"github.com/xhd2015/xgo/support/cmd"
)

func main() {
	err := Handle(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Handle(args []string) error {
	return cmd.Debug().Run("go", append([]string{"run", ".", "--dev"}, args...)...)
}
`

func FixScriptDev(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptDevPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/dev",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptDevPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/dev"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptDevPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptDevStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptDevPath)}
	return result, nil
}
