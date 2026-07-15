# Scenario

**Feature**: slash-form rule IDs on Catalog and `scaff fix`

```
# dual surface: package Catalog + real scaff CLI
catalog leaves  -> rules.Catalog -> Response.IDs
fix leaves      -> scaff binary  -> Response.ExitCode/Combined
```

## Preconditions

- Module is `github.com/xhd2015/scaff`.
- **Catalog leaves**: `rules.Catalog` is the public slice under test; no binary required.
- **Fix leaves**: build `./cmd/scaff` once per session into a temp cache keyed by
  `DOCTEST_SESSION_ID`; each leaf uses an isolated temp project directory.

## Steps

1. Confirm Catalog is importable (shared).
2. Catalog descendants leave `req.Args` empty so `Run` returns Catalog IDs.
3. Fix descendants set `ScaffBin`, `ProjectDir`, and `Args`, then `Run` execs the CLI.

## Context

- Classic TDD for P2: fix dispatch may still only know dotted IDs; slash accept /
  dotted reject leaves are expected RED until wired.
- P1 catalog leaves stay authoritative for Catalog ID strings.
- Helpers below amortize binary build and fixture writes for the `fix/` branch.

```go
import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/xhd2015/scaff/internal/rules"
)

var (
	scaffBinOnce sync.Once
	scaffBinPath string
	scaffBinErr  error
)

func Setup(t *testing.T, req *Request) error {
	if req == nil {
		return fmt.Errorf("nil request")
	}
	// Catalog must exist as a package-level inventory.
	_ = rules.Catalog
	return nil
}

func repoRoot(t *testing.T) string {
	t.Helper()
	root, err := filepath.Abs(filepath.Join(DOCTEST_ROOT, "..", ".."))
	if err != nil {
		t.Fatalf("repo root: %v", err)
	}
	return root
}

func buildScaffBinary(t *testing.T) string {
	t.Helper()
	scaffBinOnce.Do(func() {
		dir := filepath.Join(os.TempDir(), "scaff-rule-ids-"+DOCTEST_SESSION_ID)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			scaffBinErr = err
			return
		}
		scaffBinPath = filepath.Join(dir, "scaff")
		build := exec.Command("go", "build", "-o", scaffBinPath, "./cmd/scaff")
		build.Dir = repoRoot(t)
		if output, err := build.CombinedOutput(); err != nil {
			scaffBinErr = fmt.Errorf("go build ./cmd/scaff: %w: %s", err, strings.TrimSpace(string(output)))
		}
	})
	if scaffBinErr != nil {
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
```
