package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const scriptGeneratePath = "script/generate/main.go"

const scriptGenerateStub = `// usage: go run ./script/generate [targets...]
//
// Proposed behavior (sketch):
//   1. Accept optional generator target names as args.
//   2. Dispatch each target to a subpackage generator.
//   3. Exit non-zero if any generator fails; no-op until wired.
package main

func main() {
    // add generators as subpackages; wire them here
}
`

func FixScriptGenerate(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, scriptGeneratePath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "script/generate",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", scriptGeneratePath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "script/generate"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", scriptGeneratePath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(scriptGenerateStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", scriptGeneratePath)}
	return result, nil
}