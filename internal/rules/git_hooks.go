package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const gitHooksPath = "script/git-hooks/main.go"

const gitHooksStub = `// usage: go run ./script/git-hooks <install|pre-commit|pre-push>
//
// Proposed behavior (sketch):
//   1. install: patch .git/hooks/pre-commit and pre-push to invoke this runner.
//   2. pre-commit: run local checks before commit (stub until customized).
//   3. pre-push: run local checks before push (stub until customized).
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xhd2015/xgo/support/fileutil"
	"github.com/xhd2015/xgo/support/git"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: git-hooks <install|pre-commit|pre-push>")
		os.Exit(2)
	}
	switch os.Args[1] {
	case "install":
		if err := install(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "pre-commit", "pre-push":
		// no-op stub
	default:
		fmt.Fprintln(os.Stderr, "unknown command")
		os.Exit(2)
	}
}

const hookMarker = "# scaff hooks"

const preCommitHook = "go run ./script/git-hooks pre-commit"
const prePushHook = "go run ./script/git-hooks pre-push"

func install() error {
	gitDir, err := git.GetGitDir("")
	if err != nil {
		return err
	}
	hooksDir := filepath.Join(gitDir, "hooks")
	if err := os.MkdirAll(hooksDir, 0o755); err != nil {
		return err
	}
	if err := installHook(filepath.Join(hooksDir, "pre-commit"), preCommitHook); err != nil {
		return fmt.Errorf("pre-commit: %w", err)
	}
	if err := installHook(filepath.Join(hooksDir, "pre-push"), prePushHook); err != nil {
		return fmt.Errorf("pre-push: %w", err)
	}
	return nil
}

func installHook(hookFile, cmd string) error {
	var needChmod bool
	err := fileutil.Patch(hookFile, func(data []byte) ([]byte, error) {
		if len(data) == 0 {
			needChmod = true
		}
		content := string(data)
		lines := strings.Split(content, "\n")
		idx := -1
		for i, line := range lines {
			if strings.Contains(line, hookMarker) {
				idx = i
				break
			}
		}
		if idx < 0 {
			lines = append(lines, hookMarker, cmd, "")
		} else {
			endIdx := idx + 1
			for endIdx < len(lines) && strings.TrimSpace(lines[endIdx]) != "" {
				endIdx++
			}
			oldLines := lines
			lines = lines[:idx]
			lines = append(lines, hookMarker, cmd, "")
			if endIdx < len(oldLines) {
				lines = append(lines, oldLines[endIdx:]...)
			}
		}
		return []byte(strings.Join(lines, "\n")), nil
	})
	if err != nil {
		return err
	}
	if needChmod {
		return os.Chmod(hookFile, 0o755)
	}
	return nil
}
`

func FixGitHooks(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, gitHooksPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "git/hooks",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", gitHooksPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "git/hooks"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", gitHooksPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(gitHooksStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", gitHooksPath)}
	return result, nil
}