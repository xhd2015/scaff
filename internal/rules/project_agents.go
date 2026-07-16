package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const agentsPath = "AGENTS.md"

const agentsTemplate = `# AGENTS.md

## Project Overview

__NAME__ is a Go project.

## Build

go run ./script/build

## Test

go test ./...
doctest test ./tests/__NAME__-cli/...
`

func FixProjectAgents(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, agentsPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "project/agents",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", agentsPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}

	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}

	content := substituteMeta(agentsTemplate, meta)
	result := model.FixResult{RuleID: "project/agents"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", agentsPath)}
		return result, nil
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", agentsPath)}
	return result, nil
}
