package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const doctestTemplate = `# __NAME__ CLI Tests

## Version
0.0.2

# DSN (Domain Specific Notion)

The **__NAME__ CLI** is exercised by doc-style tests.

## How to Run

` + "```sh" + `
doctest vet ./tests/__NAME__-cli
doctest test ./tests/__NAME__-cli
` + "```" + `

` + "```go" + `
type Request struct{}

type Response struct{}

func Run(t *testing.T, req *Request) (*Response, error) {
	return &Response{}, nil
}
` + "```" + `
`

const doctestSetupTemplate = `# Scenario

**Feature**: shared setup for __NAME__ CLI tests

` + "```" + `
# shared doctest harness for __NAME__
__NAME__ CLI -> doctest harness -> assertions
` + "```" + `

` + "```go" + `
func Setup(t *testing.T, req *Request) error {
	t.Helper()
	return nil
}
` + "```" + `
`

func doctestRelPath(name string) string {
	return filepath.Join("tests", name+"-cli", "DOCTEST.md")
}

func doctestSetupRelPath(name string) string {
	return filepath.Join("tests", name+"-cli", "SETUP.md")
}

func LintTestsDoctest(project model.Project) model.RuleResult {
	result := model.RuleResult{
		ID:     "tests/doctest",
		Status: model.RuleOK,
	}
	switch project.Profile {
	case model.ProfileNode, model.ProfileGeneric:
		result.Message = "n/a for profile"
		return result
	}

	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}

	rel := doctestRelPath(meta.Name)
	result.Paths = []string{rel}
	path := filepath.Join(project.Root, rel)
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			result.Status = model.RuleMissing
			result.Message = fmt.Sprintf("missing: %s", rel)
			return result
		}
		result.Status = model.RuleMissing
		result.Message = err.Error()
		return result
	}
	result.Message = "present"
	return result
}

func FixTestsDoctest(project model.Project, dryRun bool) (model.FixResult, error) {
	meta, err := DetectProjectMeta(project.Root)
	if err != nil {
		return model.FixResult{}, err
	}

	doctestRel := doctestRelPath(meta.Name)
	setupRel := doctestSetupRelPath(meta.Name)
	doctestPath := filepath.Join(project.Root, doctestRel)

	if _, err := os.Stat(doctestPath); err == nil {
		return model.FixResult{
			RuleID:  "tests/doctest",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", doctestRel)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}

	result := model.FixResult{RuleID: "tests/doctest"}
	if dryRun {
		result.Actions = []string{
			fmt.Sprintf("dry-run: would create %s", doctestRel),
			fmt.Sprintf("dry-run: would create %s", setupRel),
		}
		return result, nil
	}

	if err := os.MkdirAll(filepath.Dir(doctestPath), 0o755); err != nil {
		return model.FixResult{}, err
	}
	doctestContent := substituteMeta(doctestTemplate, meta)
	if err := os.WriteFile(doctestPath, []byte(doctestContent), 0o644); err != nil {
		return model.FixResult{}, err
	}
	setupPath := filepath.Join(project.Root, setupRel)
	setupContent := substituteMeta(doctestSetupTemplate, meta)
	if err := os.WriteFile(setupPath, []byte(setupContent), 0o644); err != nil {
		return model.FixResult{}, err
	}

	result.Changed = true
	result.Actions = []string{
		fmt.Sprintf("created %s", doctestRel),
		fmt.Sprintf("created %s", setupRel),
	}
	return result, nil
}
