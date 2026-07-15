package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptInstallPath = "script/install/install.go"

const scriptInstallStub = `// usage: go run ./script/install
// build via go run ./script/build; then go install
//
// Proposed behavior (sketch):
//   1. Build the project via go run ./script/build.
//   2. Install the module with go install .
//   3. Exit non-zero if either step fails.
package main

import (
	"fmt"
	"os"

	"github.com/xhd2015/xgo/support/cmd"
)

func main() {
	if err := handle(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handle() error {
	fmt.Println("==> Building")
	if err := cmd.Debug().Run("go", "run", "./script/build"); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}
	fmt.Println("==> Installing")
	if err := cmd.Debug().Run("go", "install", "."); err != nil {
		return fmt.Errorf("go install failed: %w", err)
	}
	fmt.Println("install complete")
	return nil
}
`

func FixScriptInstall(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptInstallPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/install",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptInstallPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/install"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptInstallPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptInstallStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptInstallPath)}
	return result, nil
}