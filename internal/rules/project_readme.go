package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xhd2015/scaff/internal/model"
)

const readmePath = "README.md"

func LintProjectReadme(project model.Project) model.RuleResult {
	result := model.RuleResult{
		ID:     "project/readme",
		Paths:  []string{readmePath},
		Status: model.RuleOK,
	}
	path := filepath.Join(project.Root, readmePath)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			result.Status = model.RuleMissing
			result.Message = fmt.Sprintf("missing: %s", readmePath)
			return result
		}
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}
	result.Message = "present"
	return result
}

func FixProjectReadme(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, readmePath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "project/readme",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", readmePath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}
	content := readmeTemplate(project.Profile, meta)
	result := model.FixResult{RuleID: "project/readme"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", readmePath)}
		return result, nil
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", readmePath)}
	return result, nil
}

func readmeTemplate(profile model.Profile, meta ProjectMeta) string {
	var builder strings.Builder
	builder.WriteString("# __NAME__\n\n")
	switch profile {
	case model.ProfileNode:
		builder.WriteString("## Install\n\n")
		builder.WriteString("npm install\n")
		builder.WriteString("npm run dev\n\n")
	case model.ProfileGeneric:
		// No install section for generic profile.
	default:
		builder.WriteString("## Install\n\n")
		builder.WriteString("go install __MODULE__@latest\n\n")
	}
	builder.WriteString("## Usage\n\n")
	builder.WriteString("...\n")
	return substituteMeta(builder.String(), meta)
}
