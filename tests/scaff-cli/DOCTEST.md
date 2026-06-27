# scaff CLI Tests

Doc-style tests for the `scaff` project scaffolding auditor and fixer CLI.

## Version
0.0.2

# DSN (Domain Specific Notion)

The **scaff CLI** amends scaffolding in existing projects. It exposes two
subcommands: `lint` (read-only audit) and `fix <rule>` (idempotent repair).

**Project detector** inspects a directory tree, auto-detecting profile from
`go.mod` and `package.json` (`go`, `node`, `polyglot`, or `generic`). The
`--profile` flag overrides detection.

**Lint orchestrator** runs default rules only: `git.ignore` (expected patterns
per profile) and `github.testing.workflow` (requires `.github/workflows/test.yml`
specifically). Opt-in rules (`script.generate`, `git.hooks`, `git.hooks.install`)
are never reported by lint.

**Fix executor** applies one dotted rule at a time, merging non-destructively
(append missing `.gitignore` lines, create missing files, scaffold hooks).
`--dry-run` previews changes without writing.

**Rule implementations** each own detection and repair for their artifact:
`.gitignore`, GitHub workflow YAML, `script/generate`, `script/git-hooks`, and
hook installation into `.git/hooks/`.

Tests build the `scaff` binary once per session, materialize temp project
fixtures, exec the CLI, and assert on exit codes, stdout/stderr, and filesystem
side effects.

## Decision Tree

```
tests/scaff-cli/                              [Command, Rule, Flags, Fixture]
│
├── lint/                                     Command=lint
│   ├── issues-found/                         exit 1 — scaffold gaps
│   │   ├── empty-go-project/                 go.mod only; 2 rules; no opt-in rules
│   │   ├── partial-gitignore/                missing .vscode/ → partial git.ignore
│   │   └── ci-without-test-yml/              ci.yml present, test.yml missing
│   ├── all-pass/                             exit 0
│   │   └── complete-project/                 full gitignore + test.yml
│   ├── json-output/                          --json structured report
│   │   └── issues-report/                    valid JSON on issues
│   ├── profile-override/                     --profile overrides auto-detect
│   │   ├── node/                             node_modules expected
│   │   └── generic/                          only universal gitignore patterns
│   └── target-dir/                           --dir aims at subdirectory
│       └── subdirectory/                     audit nested project root
│
└── fix/                                      Command=fix <rule>
    ├── unknown-rule/                         exit 2, lists available rules
    ├── git-ignore/
    │   ├── create-missing/                   no .gitignore → full pattern set
    │   ├── append-partial/                   append only missing lines
    │   ├── idempotent/                       complete → no-op
    │   ├── dry-run/                          preview append, no write
    │   └── no-vendor/                        never adds vendor/
    ├── github-testing-workflow/
    │   ├── create-missing/                   creates test.yml with go+doctest steps
    │   ├── idempotent-existing/              test.yml exists → no-op
    │   ├── ci-only/                          ci.yml present, creates test.yml
    │   └── dry-run/                          preview create, no write
    ├── script-generate/
    │   ├── create-missing/                   creates stub main.go
    │   └── idempotent/                       exists → no-op
    ├── git-hooks/
    │   ├── scaffold-missing/                 creates install + no-op hooks
    │   └── idempotent/                       exists → no-op
    └── git-hooks-install/
        ├── without-scaffold/                 error + hint to fix git.hooks
        ├── patches-hooks/                    git repo → patches pre-commit/pre-push
        └── non-git/                          no .git → error
```

## Test Index

| Leaf | Description |
|------|-------------|
| `lint/issues-found/empty-go-project` | Empty Go project reports only default lint rules |
| `lint/issues-found/partial-gitignore` | Partial `.gitignore` yields partial/missing git.ignore |
| `lint/issues-found/ci-without-test-yml` | `ci.yml` does not satisfy github.testing.workflow |
| `lint/all-pass/complete-project` | Complete scaffold exits 0 |
| `lint/json-output/issues-report` | `--json` emits valid structured lint report |
| `lint/profile-override/node` | `--profile node` checks node_modules pattern |
| `lint/profile-override/generic` | Generic profile checks only universal patterns |
| `lint/target-dir/subdirectory` | `--dir` audits a nested project root |
| `fix/unknown-rule` | Unknown rule exits 2 and lists available rules |
| `fix/git-ignore/create-missing` | Creates `.gitignore` with full Go pattern set |
| `fix/git-ignore/append-partial` | Appends only missing patterns |
| `fix/git-ignore/idempotent` | Second run is no-op |
| `fix/git-ignore/dry-run` | `--dry-run` previews without writing |
| `fix/git-ignore/no-vendor` | Never adds `vendor/` |
| `fix/github-testing-workflow/create-missing` | Creates `test.yml` with go test + doctest |
| `fix/github-testing-workflow/idempotent-existing` | Existing `test.yml` is no-op |
| `fix/github-testing-workflow/ci-only` | Creates `test.yml` when only `ci.yml` exists |
| `fix/github-testing-workflow/dry-run` | `--dry-run` previews create without write |
| `fix/script-generate/create-missing` | Creates `script/generate/main.go` stub |
| `fix/script-generate/idempotent` | Existing stub is no-op |
| `fix/git-hooks/scaffold-missing` | Scaffolds hook runner without sub-check dirs |
| `fix/git-hooks/idempotent` | Existing scaffold is no-op |
| `fix/git-hooks-install/without-scaffold` | Errors with hint when hooks not scaffolded |
| `fix/git-hooks-install/patches-hooks` | Patches `.git/hooks` with scaff marker |
| `fix/git-hooks-install/non-git` | Errors when directory is not a git repo |

## How to Run

```sh
doctest vet ./tests/scaff-cli
doctest test -v ./tests/scaff-cli
```

```go
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

type Request struct {
	Args       []string
	ProjectDir string
	RunDir     string
	ScaffBin   string
}

type Response struct {
	Stdout   string
	Stderr   string
	Combined string
	ExitCode int
}

var (
	scaffBinOnce sync.Once
	scaffBinPath string
	scaffBinErr  error
)

func Run(t *testing.T, req *Request) (*Response, error) {
	runDir := req.ProjectDir
	if req.RunDir != "" {
		runDir = req.RunDir
	}
	cmd := exec.Command(req.ScaffBin, req.Args...)
	cmd.Dir = runDir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	runErr := cmd.Run()
	exitCode := 0
	if runErr != nil {
		if exitErr, ok := runErr.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return nil, runErr
		}
	}
	out := stdout.String()
	errOut := stderr.String()
	combined := out
	if errOut != "" {
		if combined != "" {
			combined += "\n"
		}
		combined += errOut
	}
	return &Response{
		Stdout:   out,
		Stderr:   errOut,
		Combined: combined,
		ExitCode: exitCode,
	}, nil
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
		dir := filepath.Join(os.TempDir(), "scaff-doctest-"+DOCTEST_SESSION_ID)
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
```