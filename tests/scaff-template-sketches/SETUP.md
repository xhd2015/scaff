# Scenario

**Feature**: scaffolded Go templates include Proposed behavior sketches

```
# empty project -> scaff fix <rule> -> generated .go contains sketch header
User -> scaff fix -> template write -> file content
```

## Preconditions

- Module is `github.com/xhd2015/scaff`.
- Each leaf uses an isolated temporary project directory.
- The `scaff` binary is built once per process (in-memory mutex + memo) from
  `./cmd/scaff` into `os.MkdirTemp("", "scaff-template-sketches-doctest-bin-")`.

## Steps

1. Allocate a temp project and build (or reuse) the process-local `scaff` binary.
2. Leaf setups write minimal `go.mod` and set `req.Args` for the fix rule.
3. `Run` executes `scaff fix <rule>` from the project directory.
4. Leaf asserts inspect the generated Go file for `Proposed behavior`.

## Context

- Sketch marker is the substring `Proposed behavior` (case-sensitive as written
  in the template comment).
- Expected layout is roughly:

  ```
  // usage: ...
  //
  // Proposed behavior (sketch):
  //   1. ...
  ```

- Leaves only require the phrase; they do not pin exact step text (implementer
  may draft rule-specific sketches).

```go
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/xhd2015/doctest/session"
)

// Process-local scaff binary (one-process suite; in-memory mutex).
var (
	scaffBinMu   sync.Mutex
	scaffBinPath string
	scaffBinErr  error
)

func Setup(t *testing.T, d *session.Doctest, req *Request) error {
	if req == nil {
		return fmt.Errorf("nil request")
	}
	req.ProjectDir = t.TempDir()
	req.RunDir = req.ProjectDir
	req.ScaffBin = buildScaffBinary(t, d)
	return nil
}

func buildScaffBinary(t *testing.T, d *session.Doctest) string {
	t.Helper()
	scaffBinMu.Lock()
	defer scaffBinMu.Unlock()
	if scaffBinPath != "" || scaffBinErr != nil {
		if scaffBinErr != nil {
			t.Fatal(scaffBinErr)
		}
		return scaffBinPath
	}
	dir, err := os.MkdirTemp("", "scaff-template-sketches-doctest-bin-")
	if err != nil {
		scaffBinErr = err
		t.Fatal(err)
	}
	scaffBinPath = filepath.Join(dir, "scaff")
	moduleRoot := filepath.Clean(filepath.Join(d.DOCTEST_ROOT, "..", ".."))
	build := exec.Command("go", "build", "-o", scaffBinPath, "./cmd/scaff")
	build.Dir = moduleRoot
	if output, err := build.CombinedOutput(); err != nil {
		scaffBinErr = fmt.Errorf("go build ./cmd/scaff: %w: %s", err, strings.TrimSpace(string(output)))
		t.Fatal(scaffBinErr)
	}
	return scaffBinPath
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

func writeGoModGitHubScaffold(dir string) error {
	return writeFile(dir, "go.mod", "module github.com/xhd2015/myapp\n\ngo 1.22\n")
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

func assertProposedBehaviorSketch(t *testing.T, content, rel string) {
	t.Helper()
	if content == "" {
		t.Fatalf("%s is empty or missing", rel)
	}
	if !strings.Contains(content, "Proposed behavior") {
		t.Fatalf("%s must contain \"Proposed behavior\" sketch comment, got:\n%s", rel, content)
	}
}
```
