# scaff CLI Tests

Doc-style tests for the `scaff` project scaffolding auditor and fixer CLI.

## Version
0.0.2

# DSN (Domain Specific Notion)

The **scaff CLI** amends scaffolding in existing projects. It exposes domain
subcommands `lint` (read-only audit), `fix <rule>` (idempotent repair), and
`rules` (rule inventory), plus a **multi-topic skill** surface under
`scaff skill` only (no top-level `install` or `topics` aliases).

**Project detector** inspects a directory tree, auto-detecting profile from
`go.mod` and `package.json` (`go`, `node`, `polyglot`, or `generic`). The
`--profile` flag overrides detection.

**Lint orchestrator** runs default rules only: `git/ignore` (expected patterns
per profile), `github/testing-workflow` (requires `.github/workflows/test.yml`
specifically), `project/readme` (root `README.md`), `project/license` (root
`LICENSE`), and `tests/doctest` (Go profile: `tests/<name>-cli/DOCTEST.md`).
Opt-in rules (`script/generate`, `project/agents`, `script/dev`, …) are never
reported by lint.

**Fix executor** applies one slash-form rule at a time, merging non-destructively
(append missing `.gitignore` lines, create missing files, scaffold hooks).
`--dry-run` previews changes without writing.

**Rule implementations** each own detection and repair for their artifact:
`.gitignore`, GitHub workflow YAML, root `README.md` / `LICENSE` / `AGENTS.md`,
`cmd/<name>/main.go`, doctest harness under `tests/<name>-cli/`,
`script/generate`, `script/install`, `script/build`, `script/dev`,
`script/bundle/for-linux`, `script/git-hooks`, hook installation into
`.git/hooks/`, GitHub release scripts under `script/github/`, release-assets
helper under `script/github/release-assets/`, and the root `install-via-curl.sh`
curl installer.

Opt-in fix rules (`project/agents`, `project/layout/cmd`, `script/dev`,
`github/release`, `install/via-curl`, `script/github/release-assets`, …) are
not part of default lint. Scaffold templates may substitute project metadata
from `go.mod` (`__NAME__`, `__OWNER__`, `__REPO__`, `__YEAR__`).

**Skill host** embeds a Shape 3 multi-topic skill (`docs/SKILL.md` + nested
`docs/<path>/TOPIC.md` via `docs/embed.go` and `skillcmd.SingleSkill`). Users
retrieve the root index or a nested topic with `scaff skill --show`, list every
topic path with `scaff skill --list`, install the skill tree with
`scaff skill --install`, and discover actions/topics with `scaff skill --help`.
Both flag orders work for show (`--show path` and `path --show`). `--header`
prints YAML frontmatter only. Unknown topic paths error. Install supports
`--dry-run` and a positional target directory.

Tests build the `scaff` binary once per session, materialize temp project
fixtures (for lint/fix), exec the CLI, and assert on exit codes, stdout/stderr,
and filesystem side effects.

## Decision Tree

```
tests/scaff-cli/                              [Command, Rule/Topic, Flags, Fixture]
│
├── lint/                                     Command=lint
│   ├── issues-found/                         exit 1 — scaffold gaps
│   │   ├── empty-go-project/                 go.mod only; 5 default rules; no opt-in
│   │   ├── partial-gitignore/                missing .vscode/ → partial git/ignore
│   │   ├── ci-without-test-yml/              ci.yml present, test.yml missing
│   │   ├── missing-readme/                   go.mod only; lint mentions project/readme
│   │   ├── missing-license/                  go.mod only; lint mentions project/license
│   │   └── missing-doctest/                  go.mod only; lint mentions tests/doctest
│   ├── all-pass/                             exit 0
│   │   └── complete-project/                 full default-lint scaffold
│   ├── json-output/                          --json structured report
│   │   └── issues-report/                    valid JSON on issues
│   ├── profile-override/                     --profile overrides auto-detect
│   │   ├── node/                             node_modules expected
│   │   └── generic/                          only universal gitignore patterns
│   └── target-dir/                           --dir aims at subdirectory
│       └── subdirectory/                     audit nested project root
│
├── fix/                                      Command=fix <rule>
│   ├── unknown-rule/                         exit 2, lists available rules
│   ├── git-ignore/
│   │   ├── create-missing/                   no .gitignore → full pattern set
│   │   ├── append-partial/                   append only missing lines
│   │   ├── idempotent/                       complete → no-op
│   │   ├── dry-run/                          preview append, no write
│   │   └── no-vendor/                        never adds vendor/
│   ├── github-testing-workflow/
│   │   ├── create-missing/                   creates test.yml with go+doctest steps
│   │   ├── idempotent-existing/              test.yml exists → no-op
│   │   ├── ci-only/                          ci.yml present, creates test.yml
│   │   └── dry-run/                          preview create, no write
│   ├── script-generate/
│   │   ├── create-missing/                   creates stub main.go
│   │   └── idempotent/                       exists → no-op
│   ├── script-install/
│   │   ├── create-missing/                   creates install.go stub
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── script-build/
│   │   ├── create-missing/                   creates build.go stub
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── script-dev/                           Rule=script/dev
│   │   ├── create-missing/                   creates script/dev/main.go wrapper
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── project-readme/                       Rule=project/readme
│   │   ├── create-missing/                   creates README.md with install line
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── project-license/                      Rule=project/license
│   │   ├── create-missing/                   creates MIT LICENSE with year/owner
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── tests-doctest/                        Rule=tests/doctest
│   │   ├── create-missing/                   creates tests/<name>-cli doctest tree
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── project-agents/                       Rule=project/agents
│   │   ├── create-missing/                   creates AGENTS.md
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── project-layout-cmd/                   Rule=project/layout/cmd
│   │   ├── create-missing/                   creates cmd/<name>/main.go
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── script-bundle-for-linux/
│   │   ├── create-missing/                   creates for-linux main.go stub
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview create, no write
│   ├── git-hooks/
│   │   ├── scaffold-missing/                 creates install + no-op hooks
│   │   └── idempotent/                       exists → no-op
│   ├── git-hooks-install/
│   │   ├── without-scaffold/                 error + hint to fix git/hooks
│   │   ├── patches-hooks/                    git repo → patches pre-commit/pre-push
│   │   └── non-git/                          no .git → error
│   ├── github-release/                       Rule=github/release
│   │   ├── create-missing/                   go.mod → release main + lib
│   │   ├── partial-scaffold/                 main exists → create lib only
│   │   ├── idempotent/                       both exist → no-op
│   │   └── dry-run/                          preview paths, no write
│   ├── install-via-curl/                     Rule=install/via-curl
│   │   ├── create-missing/                   curl installer at repo root
│   │   ├── idempotent/                       exists → no-op
│   │   └── dry-run/                          preview only
│   └── script-github-release-assets/            Rule=script/github/release-assets
│       ├── create-missing/                   creates main.go + Proposed behavior
│       ├── idempotent/                       exists → no-op
│       └── dry-run/                          preview create, no write
│
└── skill/                                    Command=skill (multi-topic skillcmd)
    ├── list/
    │   └── full-inventory/                   --list name + all topic paths (sorted)
    ├── show/
    │   ├── root/                             --show root SKILL.md index body
    │   ├── header/                           --show --header YAML delimiters only
    │   ├── topic/
    │   │   ├── overview/                     --show overview
    │   │   ├── git-ignore-flag-before/       --show git/ignore
    │   │   ├── git-ignore-path-before/       git/ignore --show
    │   │   ├── github-release/               --show github/release
    │   │   ├── github-upload/                --show github/upload
    │   │   └── install-via-curl/             --show install-via-curl
    │   └── unknown/                          --show missing topic → error
    ├── help/
    │   └── with-topics/                      skill --help + Available topics
    ├── install/
    │   └── dry-run/                          --install --dry-run <tempDir>
    └── no-alias/
        ├── top-level-install/                scaff install → unknown command
        └── top-level-topics/                 scaff topics → unknown command
```

## Test Index

| Leaf | Description |
|------|-------------|
| `lint/issues-found/empty-go-project` | Empty Go project reports only default lint rules |
| `lint/issues-found/partial-gitignore` | Partial `.gitignore` yields partial/missing git/ignore |
| `lint/issues-found/ci-without-test-yml` | `ci.yml` does not satisfy github/testing-workflow |
| `lint/issues-found/missing-readme` | Missing `README.md` yields project/readme finding |
| `lint/issues-found/missing-license` | Missing `LICENSE` yields project/license finding |
| `lint/issues-found/missing-doctest` | Missing doctest tree yields tests/doctest finding |
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
| `fix/script-install/create-missing` | Creates `script/install/install.go` stub |
| `fix/script-install/idempotent` | Existing install stub is no-op |
| `fix/script-install/dry-run` | `--dry-run` previews install stub without write |
| `fix/script-build/create-missing` | Creates `script/build/build.go` stub |
| `fix/script-build/idempotent` | Existing build stub is no-op |
| `fix/script-build/dry-run` | `--dry-run` previews build stub without write |
| `fix/script-dev/create-missing` | Creates `script/dev/main.go` go run . --dev wrapper |
| `fix/script-dev/idempotent` | Existing dev stub is no-op |
| `fix/script-dev/dry-run` | `--dry-run` previews dev stub without write |
| `fix/project-readme/create-missing` | Creates `README.md` with go install line and module substitutions |
| `fix/project-readme/idempotent` | Existing README is no-op |
| `fix/project-readme/dry-run` | `--dry-run` previews without writing README |
| `fix/project-license/create-missing` | Creates MIT `LICENSE` with year and owner substitutions |
| `fix/project-license/idempotent` | Existing LICENSE is no-op |
| `fix/project-license/dry-run` | `--dry-run` previews without writing LICENSE |
| `fix/tests-doctest/create-missing` | Creates `tests/<name>-cli/` doctest tree from go.mod |
| `fix/tests-doctest/idempotent` | Existing DOCTEST.md is no-op |
| `fix/tests-doctest/dry-run` | `--dry-run` previews without writing doctest tree |
| `fix/project-agents/create-missing` | Creates `AGENTS.md` with build and doctest test sections |
| `fix/project-agents/idempotent` | Existing AGENTS.md is no-op |
| `fix/project-agents/dry-run` | `--dry-run` previews without writing AGENTS.md |
| `fix/project-layout-cmd/create-missing` | Creates `cmd/myapp/main.go` from go.mod name |
| `fix/project-layout-cmd/idempotent` | Existing `cmd/` layout is no-op |
| `fix/project-layout-cmd/dry-run` | `--dry-run` previews without writing cmd main |
| `fix/script-bundle-for-linux/create-missing` | Creates `script/bundle/for-linux/main.go` stub |
| `fix/script-bundle-for-linux/idempotent` | Existing bundle stub is no-op |
| `fix/script-bundle-for-linux/dry-run` | `--dry-run` previews bundle stub without write |
| `fix/git-hooks/scaffold-missing` | Scaffolds hook runner without sub-check dirs |
| `fix/git-hooks/idempotent` | Existing scaffold is no-op |
| `fix/git-hooks-install/without-scaffold` | Errors with hint when hooks not scaffolded |
| `fix/git-hooks-install/patches-hooks` | Patches `.git/hooks` with scaff marker |
| `fix/git-hooks-install/non-git` | Errors when directory is not a git repo |
| `fix/github-release/create-missing` | Scaffolds release main + lib with go.mod substitutions |
| `fix/github-release/partial-scaffold` | Creates missing lib when release main already exists |
| `fix/github-release/idempotent` | Existing release scaffold is no-op |
| `fix/github-release/dry-run` | `--dry-run` previews without writing release files |
| `fix/install-via-curl/create-missing` | Creates `install-via-curl.sh` with GitHub URL patterns |
| `fix/install-via-curl/idempotent` | Existing installer script is no-op |
| `fix/install-via-curl/dry-run` | `--dry-run` previews without writing installer |
| `fix/script-github-release-assets/create-missing` | Creates `script/github/release-assets/main.go` with Proposed behavior + help |
| `fix/script-github-release-assets/idempotent` | Existing release-assets stub is no-op |
| `fix/script-github-release-assets/dry-run` | `--dry-run` previews release-assets stub without write |
| `skill/list/full-inventory` | `skill --list` prints `scaff` then full sorted topic inventory |
| `skill/show/root` | `skill --show` root body: `name: scaff`, retrieve examples, no install flags |
| `skill/show/header` | `skill --show --header` prints YAML delimiters only |
| `skill/show/topic/overview` | `skill --show overview` → `name: scaff/overview` |
| `skill/show/topic/git-ignore-flag-before` | `skill --show git/ignore` path resolution |
| `skill/show/topic/git-ignore-path-before` | `skill git/ignore --show` alternate flag order |
| `skill/show/topic/github-release` | `skill --show github/release` topic markers |
| `skill/show/topic/github-upload` | `skill --show github/upload` docs-only topic |
| `skill/show/topic/install-via-curl` | `skill --show install-via-curl` topic markers |
| `skill/show/unknown` | Unknown topic path exits non-zero with error signal |
| `skill/help/with-topics` | `skill --help` mentions actions and available topics |
| `skill/install/dry-run` | `skill --install --dry-run <dir>` plans SKILL.md + TOPIC.md, no write |
| `skill/no-alias/top-level-install` | Top-level `install` is unknown (not skill install) |
| `skill/no-alias/top-level-topics` | Top-level `topics` is unknown (not skill list) |

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

	"github.com/xhd2015/doctest/session"
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

// Process-local scaff binary (one-process suite; in-memory mutex, not session cache).
var (
	scaffBinMu   sync.Mutex
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
	dir, err := os.MkdirTemp("", "scaff-doctest-bin-")
	if err != nil {
		scaffBinErr = err
		t.Fatal(err)
	}
	scaffBinPath = filepath.Join(dir, "scaff")
	moduleRoot := filepath.Clean(filepath.Join(d.DOCTEST_ROOT, "..", ".."))
	build := exec.Command("go", "build", "-buildvcs=false", "-o", scaffBinPath, "./cmd/scaff")
	build.Dir = moduleRoot
	if output, err := build.CombinedOutput(); err != nil {
		scaffBinErr = fmt.Errorf("go build ./cmd/scaff: %w: %s", err, strings.TrimSpace(string(output)))
		t.Fatal(scaffBinErr)
	}
	return scaffBinPath
}
```