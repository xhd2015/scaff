package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xhd2015/scaff/internal/model"
	"github.com/xhd2015/xgo/support/fileutil"
	"github.com/xhd2015/xgo/support/git"
)

const hookMarker = "# scaff hooks"

func FixGitHooksInstall(project model.Project, dryRun bool) (model.FixResult, error) {
	hooksMain := filepath.Join(project.Root, gitHooksPath)
	if _, err := os.Stat(hooksMain); err != nil {
		if os.IsNotExist(err) {
			return model.FixResult{}, fmt.Errorf("missing %s; run: scaff fix git/hooks", gitHooksPath)
		}
		return model.FixResult{}, err
	}
	gitDir, err := git.GetGitDir(project.Root)
	if err != nil {
		return model.FixResult{}, fmt.Errorf("not a git repository: %w", err)
	}
	result := model.FixResult{RuleID: "git/hooks/install"}
	if dryRun {
		result.Actions = []string{"dry-run: would patch .git/hooks/pre-commit and pre-push"}
		return result, nil
	}
	hooksDir := filepath.Join(gitDir, "hooks")
	if err := os.MkdirAll(hooksDir, 0o755); err != nil {
		return model.FixResult{}, err
	}
	preCommit := filepath.Join(hooksDir, "pre-commit")
	prePush := filepath.Join(hooksDir, "pre-push")
	if err := patchHook(preCommit, "pre-commit"); err != nil {
		return model.FixResult{}, err
	}
	if err := patchHook(prePush, "pre-push"); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{"patched .git/hooks/pre-commit and pre-push"}
	return result, nil
}

func patchHook(hookFile, hookName string) error {
	cmdLine := fmt.Sprintf("go run ./script/git-hooks %s", hookName)
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
			lines = append(lines, hookMarker, cmdLine, "")
		} else {
			endIdx := idx + 1
			for endIdx < len(lines) && strings.TrimSpace(lines[endIdx]) != "" {
				endIdx++
			}
			oldLines := lines
			lines = lines[:idx]
			lines = append(lines, hookMarker, cmdLine, "")
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
