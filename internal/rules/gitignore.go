package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xhd2015/scaff/internal/model"
)

const gitignorePath = ".gitignore"

func LintGitIgnore(project model.Project) model.RuleResult {
	result := model.RuleResult{
		ID:     "git.ignore",
		Paths:  []string{gitignorePath},
		Status: model.RuleOK,
	}
	expected := PatternsForProfile(project.Profile)
	content, err := os.ReadFile(filepath.Join(project.Root, gitignorePath))
	if err != nil {
		if os.IsNotExist(err) {
			result.Status = model.RuleMissing
			result.Message = fmt.Sprintf("missing patterns: %s", strings.Join(expected, ", "))
			return result
		}
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}
	missing := missingPatterns(string(content), expected)
	if len(missing) == 0 {
		result.Message = "all patterns present"
		return result
	}
	if len(missing) == len(expected) {
		result.Status = model.RuleMissing
	} else {
		result.Status = model.RulePartial
	}
	result.Message = fmt.Sprintf("missing patterns: %s", strings.Join(missing, ", "))
	return result
}

func FixGitIgnore(project model.Project, dryRun bool) (model.FixResult, error) {
	expected := PatternsForProfile(project.Profile)
	path := filepath.Join(project.Root, gitignorePath)
	content, err := os.ReadFile(path)
	existing := ""
	if err != nil {
		if !os.IsNotExist(err) {
			return model.FixResult{}, err
		}
	} else {
		existing = string(content)
	}
	missing := missingPatterns(existing, expected)
	result := model.FixResult{RuleID: "git.ignore"}
	if len(missing) == 0 {
		result.Actions = []string{"all patterns present, nothing to do"}
		return result, nil
	}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would append to %s: %s", gitignorePath, strings.Join(missing, ", "))}
		return result, nil
	}
	var builder strings.Builder
	if existing != "" {
		builder.WriteString(strings.TrimRight(existing, "\n"))
		builder.WriteString("\n")
	}
	for _, line := range missing {
		builder.WriteString(line)
		builder.WriteString("\n")
	}
	if err := os.WriteFile(path, []byte(builder.String()), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	if existing == "" {
		result.Actions = []string{fmt.Sprintf("created %s", gitignorePath)}
	} else {
		result.Actions = []string{fmt.Sprintf("appended %d line(s) to %s", len(missing), gitignorePath)}
	}
	return result, nil
}

func missingPatterns(content string, expected []string) []string {
	lines := make(map[string]bool)
	for _, line := range strings.Split(content, "\n") {
		lines[strings.TrimSpace(line)] = true
	}
	var missing []string
	for _, pattern := range expected {
		if !lines[pattern] {
			missing = append(missing, pattern)
		}
	}
	return missing
}