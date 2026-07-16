package rules

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/xhd2015/scaff/internal/model"
)

const gitPreCommitPath = "script/git/pre-commit/main.go"

const gitPreCommitStub = `// Pre-commit helper: ensure listed paths exist (empty if missing), then git add.
//
//	go run ./script/git/pre-commit
//
// Install (managed hook via git-hooks):
//
//	git-hooks pre-commit add 'script.git.pre-commit' go run ./script/git/pre-commit
//
// Silent on success. Edit ensure paths below for your repo.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Paths relative to repo root. Empty files are created if missing, then staged.
var ensure = []string{
	// "path/to/placeholder.txt",
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	root, err := gitRoot()
	if err != nil {
		return err
	}

	var toAdd []string
	for _, rel := range ensure {
		abs := filepath.Join(root, filepath.FromSlash(rel))
		if err := ensureFile(abs); err != nil {
			return fmt.Errorf("%s: %w", rel, err)
		}
		toAdd = append(toAdd, rel)
	}

	return gitAdd(root, toAdd)
}

func gitRoot() (string, error) {
	out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("git rev-parse --show-toplevel: %w", err)
	}
	root := strings.TrimSpace(string(out))
	if root == "" {
		return "", fmt.Errorf("empty git toplevel")
	}
	return root, nil
}

func ensureFile(abs string) error {
	if st, err := os.Stat(abs); err == nil {
		if st.IsDir() {
			return fmt.Errorf("is a directory")
		}
		return nil
	} else if !os.IsNotExist(err) {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		return err
	}
	return os.WriteFile(abs, nil, 0o644)
}

func gitAdd(root string, relPaths []string) error {
	if len(relPaths) == 0 {
		return nil
	}
	args := append([]string{"add", "--"}, relPaths...)
	cmd := exec.Command("git", args...)
	cmd.Dir = root
	cmd.Stdout = nil
	out, err := cmd.CombinedOutput()
	if err != nil {
		msg := strings.TrimSpace(string(out))
		if msg != "" {
			return fmt.Errorf("git add: %s", msg)
		}
		return fmt.Errorf("git add: %w", err)
	}
	return nil
}
`

func FixGitPreCommit(project model.Project, dryRun bool) (model.FixResult, error) {
	path := filepath.Join(project.Root, gitPreCommitPath)
	if _, err := os.Stat(path); err == nil {
		return model.FixResult{
			RuleID:  "git/pre-commit",
			Actions: []string{fmt.Sprintf("%s already exists, nothing to do", gitPreCommitPath)},
		}, nil
	} else if !os.IsNotExist(err) {
		return model.FixResult{}, err
	}
	result := model.FixResult{RuleID: "git/pre-commit"}
	if dryRun {
		result.Actions = []string{fmt.Sprintf("dry-run: would create %s", gitPreCommitPath)}
		return result, nil
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return model.FixResult{}, err
	}
	if err := os.WriteFile(path, []byte(gitPreCommitStub), 0o644); err != nil {
		return model.FixResult{}, err
	}
	result.Changed = true
	result.Actions = []string{fmt.Sprintf("created %s", gitPreCommitPath)}
	return result, nil
}
