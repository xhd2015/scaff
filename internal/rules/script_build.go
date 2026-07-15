package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptBuildPath = "script/build/build.go"

const scriptBuildStub = `// usage: go run ./script/build (go build -o bin/app)
//
// Proposed behavior (sketch):
//   1. Parse optional flags if any (default: native go build).
//   2. Run go build -o bin/app for the module root.
//   3. Exit non-zero on build failure.
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
`

func FixScriptBuild(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptBuildPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/build",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptBuildPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/build"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptBuildPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptBuildStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptBuildPath)}
	return result, nil
}