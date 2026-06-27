# Scenario

**Feature**: scaff CLI lint and fix commands

```
# temp project fixture + built scaff binary
Test harness -> scaff binary -> stdout/stderr/exit code -> Response
```

## Preconditions

- The `scaff` CLI binary is built once per doctest session from `./cmd/scaff`.
- Each test case runs against an isolated temporary project directory.

## Steps

1. Allocate a temporary project directory for the test case.
2. Build or reuse the cached `scaff` binary.
3. Descendant `Setup` functions materialize project fixtures and set CLI arguments.
4. `Run` executes `scaff` with `req.Args` from `req.RunDir` (defaults to project root).

## Context

- `DOCTEST_ROOT` points at `tests/scaff-cli`.
- `DOCTEST_SESSION_ID` scopes the cached binary build directory.
- Leaf setups write fixtures under `req.ProjectDir` using shared helpers.

```go
import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func Setup(t *testing.T, req *Request) error {
	req.ProjectDir = t.TempDir()
	req.RunDir = req.ProjectDir
	req.ScaffBin = buildScaffBinary(t)
	return nil
}

func writeFile(root, rel, content string) error {
	path := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, []byte(content), 0o644)
}

func writeGoMod(dir string) error {
	return writeFile(dir, "go.mod", "module example.com/app\n\ngo 1.22\n")
}

func writePackageJSON(dir string) error {
	return writeFile(dir, "package.json", `{"name":"app","version":"1.0.0"}`+"\n")
}

func writePartialGitignore(dir string) error {
	content := strings.Join([]string{
		".DS_Store",
		"bin/",
		"*.test",
		"coverage.out",
	}, "\n") + "\n"
	return writeFile(dir, ".gitignore", content)
}

func writeCompleteGoGitignore(dir string) error {
	content := strings.Join([]string{
		".DS_Store",
		".vscode/",
		"*.swp",
		"*~",
		"bin/",
		"*.test",
		"coverage.out",
	}, "\n") + "\n"
	return writeFile(dir, ".gitignore", content)
}

func writeTestWorkflow(dir string) error {
	content := `name: Test
on:
  push:
  pull_request:
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - run: echo test
`
	return writeFile(dir, ".github/workflows/test.yml", content)
}

func writeCiWorkflow(dir string) error {
	content := `name: CI
on:
  push:
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - run: echo ci
`
	return writeFile(dir, ".github/workflows/ci.yml", content)
}

func readFixture(name string) (string, error) {
	data, err := os.ReadFile(filepath.Join(DOCTEST_ROOT, "testdata", name))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func writeScriptGenerate(dir string) error {
	content, err := readFixture("script-generate-main.go")
	if err != nil {
		return err
	}
	return writeFile(dir, "script/generate/main.go", content)
}

func writeGitHooksMain(dir string) error {
	content, err := readFixture("git-hooks-main.go")
	if err != nil {
		return err
	}
	return writeFile(dir, "script/git-hooks/main.go", content)
}

func initGitRepo(dir string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = dir
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	cfg := exec.Command("git", "config", "user.email", "scaff@test.local")
	cfg.Dir = dir
	if _, err := cfg.CombinedOutput(); err != nil {
		return err
	}
	cfg = exec.Command("git", "config", "user.name", "scaff test")
	cfg.Dir = dir
	if _, err := cfg.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func readProjectFile(t *testing.T, req *Request, rel string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(req.ProjectDir, rel))
	if err != nil {
		if os.IsNotExist(err) {
			return ""
		}
		t.Fatalf("read %s: %v", rel, err)
	}
	return string(data)
}

func fileExists(t *testing.T, req *Request, rel string) bool {
	t.Helper()
	_, err := os.Stat(filepath.Join(req.ProjectDir, rel))
	return err == nil
}
```