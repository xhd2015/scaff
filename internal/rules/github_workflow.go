package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/xhd2015/scaff/internal/model"
)

const testWorkflowPath = ".github/workflows/test.yml"

func LintGitHubTestingWorkflow(project model.Project) model.RuleResult {
	result := model.RuleResult{
		ID:     "github.testing.workflow",
		Paths:  []string{testWorkflowPath},
		Status: model.RuleOK,
	}
	path := filepath.Join(project.Root, testWorkflowPath)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			result.Status = model.RuleMissing
			result.Message = fmt.Sprintf("missing: %s", testWorkflowPath)
			return result
		}
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}
	result.Message = "present"
	return result
}

func FixGitHubTestingWorkflow(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, testWorkflowPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "github.testing.workflow",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", testWorkflowPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	content := testWorkflowTemplate(detectGoVersion(project.Root))
	result := model.FixResult{RuleID: "github.testing.workflow"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", testWorkflowPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", testWorkflowPath)}
	return result, nil
}

func detectGoVersion(root string) string {
	data, err := os.ReadFile(filepath.Join(root, "go.mod"))
	if err != nil {
		return "1.22"
	}
	re := regexp.MustCompile(`(?m)^go\s+(\S+)`)
	matches := re.FindStringSubmatch(string(data))
	if len(matches) < 2 {
		return "1.22"
	}
	return matches[1]
}

func testWorkflowTemplate(goVersion string) string {
	return strings.TrimLeft(fmt.Sprintf(`name: Test

on:
  push:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    container:
      image: golang:%s
    steps:
      - uses: actions/checkout@v4
      - name: Go test
        run: go test -v ./...
      - name: Install doctest
        run: |
          if ! command -v doctest >/dev/null 2>&1; then
            go install github.com/xhd2015/doctest/cmd/doctest@latest
          fi
      - name: Doctest
        run: doctest test -v ./...
`, goVersion), "\n")
}