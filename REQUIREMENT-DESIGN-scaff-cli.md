# Requirement: scaff CLI — project scaffolding auditor & fixer

## Summary

Build `scaff`, a Go CLI that **amends scaffolding to existing projects** (not greenfield
project generation). Two commands:

```
scaff lint  [--dir DIR] [--json] [--profile go|node|polyglot]
scaff fix   <rule> [--dir DIR] [--dry-run]
```

- `lint` — read-only audit; exit 1 if gaps found; exit 0 if all default rules pass
- `fix <rule>` — apply one dotted rule's fix; idempotent; exit 2 on unknown rule

Module path: `github.com/xhd2015/scaff` (fix current `scall` typo in go.mod).

## Philosophy

Detect → explain → fix non-destructively. Never overwrite custom work; merge or skip.

## Default lint rules (only these two)

| Rule ID | What lint checks |
|---------|------------------|
| `git.ignore` | `.gitignore` exists and contains expected common patterns for profile |
| `github.testing.workflow` | `.github/workflows/test.yml` exists specifically (not "any workflow") |

**NOT linted by default:** `script.generate`, `git.hooks`, `git.hooks.install`

## Opt-in fix rules (not in default lint)

| Rule ID | Fix behavior |
|---------|--------------|
| `git.ignore` | Create `.gitignore` or append missing patterns only |
| `github.testing.workflow` | Create `.github/workflows/test.yml` if missing |
| `script.generate` | Create `script/generate/main.go` no-op stub if missing |
| `git.hooks` | Scaffold `script/git-hooks/main.go` (install + no-op pre-commit/pre-push) |
| `git.hooks.install` | Run hook install (patch `.git/hooks/` with `# scaff hooks` marker) |

## git.ignore patterns

**Always (all profiles):**
```
.DS_Store
.vscode/
*.swp
*~
```

**Go profile adds:**
```
bin/
*.test
coverage.out
```

**Node profile adds:**
```
node_modules/
dist/
.env
build/
.next/
```

Polyglot = union of Go + Node. **Do NOT include `vendor/`** — vendored code is OK to commit.

Merge strategy: append only missing lines; dedupe by exact line match; never overwrite file.

## github.testing.workflow template

Create `.github/workflows/test.yml` only when that exact file is missing.
Other workflows (e.g. `ci.yml`) do not satisfy lint.

Template content (use detected go version from go.mod when possible):

```yaml
name: Test

on:
  push:
  pull_request:

jobs:
  test:
    runs-on: ubuntu-latest
    container:
      image: golang:{{GoVersion}}
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
```

## script.generate stub

```go
// usage: go run ./script/generate [targets...]
package main

func main() {
    // add generators as subpackages; wire them here
}
```

Idempotent: if file exists, report nothing to do.

## git.hooks scaffolding

Minimal runner derived from xgo/kool pattern:

- Commands: `install`, `pre-commit`, `pre-push`
- `pre-commit` / `pre-push`: no-op stubs (no sub-check directories like `pre-commit/go-test/`)
- `install`: patches `.git/hooks/pre-commit` and `pre-push` with marker `# scaff hooks`
  and shell snippet calling `go run ./script/git-hooks <hook>`
- Preserve existing hook content; patch by marker block

## git.hooks.install

- Requires `script/git-hooks/main.go` to exist; else error with hint: `scaff fix git.hooks`
- Requires git repo; patches `.git/hooks/`

## Profile detection

Auto-detect from project root:
- `go.mod` → go profile signals
- `package.json` → node profile signals
- both → polyglot
- neither → generic (only always patterns for git.ignore)

`--profile` flag overrides auto-detect.

## CLI design

Use `less-flags` subcommand pattern (no top-level flags except lint may have --dir, --json, --profile).

Suggested layout:
```
cmd/scaff/main.go
internal/audit/       # project detection, lint orchestration
internal/rules/       # one file per rule
internal/fix/         # fix executor
internal/templates/   # go:embed templates
tests/scaff-cli/      # doctest tree (designer creates this)
```

Dependencies:
- `github.com/xhd2015/less-flags`
- `github.com/xhd2015/xgo/support/cmd` (optional, for git hooks install invocation)
- `github.com/xhd2015/xgo/support/git` (optional)
- `github.com/xhd2015/xgo/support/fileutil` (optional, for hook patching)

## Data models (in-memory, no persistent storage)

```go
type Profile string // go, node, polyglot, generic

type Project struct {
    Root    string
    Profile Profile
}

type RuleStatus string // ok, missing, partial

type RuleResult struct {
    ID      string
    Status  RuleStatus
    Message string
    Paths   []string
}

type LintReport struct {
    Project Project
    Results []RuleResult
}

type FixResult struct {
    RuleID  string
    Actions []string // human-readable lines
    Changed bool
}
```

## Exit codes

- 0: success / all rules pass
- 1: lint found issues OR fix rule failed
- 2: usage error (unknown rule, missing args)

## UX examples

```
$ scaff lint
scaff lint: 2 issue(s) in .

  git.ignore                   missing patterns: .vscode/, .DS_Store
  github.testing.workflow      missing: .github/workflows/test.yml

hint: scaff fix <rule>

$ scaff fix git.ignore
appended 2 line(s) to .gitignore

$ scaff fix github.testing.workflow
created .github/workflows/test.yml

$ scaff lint
scaff lint: all good (2/2 rules passing)

$ scaff fix unknown.rule
scaff fix: unknown rule "unknown.rule"
available rules: git.ignore, github.testing.workflow, script.generate, git.hooks, git.hooks.install
exit 2
```

## Test scenarios to cover (doctest tree)

### lint command
- Empty Go project (go.mod only): exit 1, reports git.ignore + github.testing.workflow only
- Does NOT mention script.generate or git.hooks
- Project with complete git.ignore + test.yml: exit 0
- Partial git.ignore (missing .vscode/): partial/missing status for git.ignore
- ci.yml exists but test.yml missing: still reports github.testing.workflow
- --json output: valid structured report
- --profile node: checks node_modules in git.ignore patterns
- --dir points at subdirectory

### fix git.ignore
- No .gitignore: creates with full pattern set
- Partial .gitignore: appends only missing lines
- Second run: no-op, "nothing to do" or "all patterns present"
- --dry-run: prints would-append lines, file unchanged
- Does not add vendor/

### fix github.testing.workflow
- Missing test.yml: creates with go test + doctest steps
- test.yml already exists: no-op
- ci.yml exists, test.yml missing: creates test.yml
- --dry-run: reports would-create, no write

### fix script.generate
- Missing: creates stub
- Exists: no-op
- Not reported by lint

### fix git.hooks
- Missing: creates script/git-hooks/main.go with install + no-op hooks
- No pre-commit/go-test subdirs
- Exists: no-op

### fix git.hooks.install
- Without git.hooks scaffolded: error + hint
- With git.hooks + git init: patches .git/hooks/pre-commit and pre-push
- Non-git dir: error

### fix unknown rule
- exit 2, lists available rules

## How to test

Design a doctest tree under `tests/scaff-cli/` following doctest spec v0.0.2.
Use temp directories and `exec` of built `scaff` binary (or `go run ./cmd/scaff`).
Tests should be RED before implementation exists.

Verify commands:
```sh
doctest vet ./tests/scaff-cli
doctest test -v ./tests/scaff-cli
```

## Constraints

- Do NOT write implementation code — tests only
- Follow MECE decision tree, significance-ordered grouping
- Prefer table-driven Run() in root DOCTEST.md that dispatches to scaff CLI