package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const projectLayoutCmdStub = `package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "__NAME__: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	return fmt.Errorf("not implemented")
}
`

func FixProjectLayoutCmd(project model.Project, dryRun bool) (model.FixResult, error) {
	cmdDir := filepath.Join(project.Root, "cmd")
	if _, err := os.Stat(cmdDir); err == nil {
		return model.FixResult{
			RuleID:  "project/layout/cmd",
			Actions: []string{"cmd/ already exists, nothing to do"},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}

	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}

	relPath := filepath.Join("cmd", meta.Name, "main.go")
	path := filepath.Join(project.Root, relPath)

	result := model.FixResult{RuleID: "project/layout/cmd"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", relPath)}
		return result, nil
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	content := substituteMeta(projectLayoutCmdStub, meta)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", relPath)}
	return result, nil
}
