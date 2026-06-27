# Implement: scaff CLI

## Context

The scaff CLI amends scaffolding to existing projects. Requirements are in
`REQUIREMENT-DESIGN-scaff-cli.md`. A doctest tree has been designed and sealed.

## Tests are sealed — do not modify

The doctest tree at `./tests/scaff-cli/` is committed and sealed.
**Do NOT modify any file under `./tests/scaff-cli/`.**

25 leaves covering `scaff lint` and `scaff fix <rule>`.

## Feature summary

Implement `cmd/scaff` and supporting packages:

```
scaff lint  [--dir DIR] [--json] [--profile go|node|polyglot]
scaff fix   <rule> [--dir DIR] [--dry-run]
```

### Default lint rules only
- `git.ignore` — check .gitignore has profile-aware patterns (no vendor/)
- `github.testing.workflow` — check `.github/workflows/test.yml` exists

### Fix rules (all available via scaff fix)
- `git.ignore` — create/append patterns
- `github.testing.workflow` — create test.yml with go test + doctest
- `script.generate` — no-op stub at script/generate/main.go
- `git.hooks` — script/git-hooks/main.go with install + no-op pre-commit/pre-push
- `git.hooks.install` — patch .git/hooks/ with # scaff hooks marker

### Module
- Fix go.mod to `github.com/xhd2015/scaff`

### Dependencies
- `github.com/xhd2015/less-flags` for CLI parsing
- May use `github.com/xhd2015/xgo/support/cmd`, `git`, `fileutil` for hooks

### Exit codes
- 0: success
- 1: lint issues or fix failure
- 2: usage error (unknown rule)

## Test tree structure

```
tests/scaff-cli/
├── lint/ (8 leaves)
└── fix/ (17 leaves)
```

Read `tests/scaff-cli/DOCTEST.md` and leaf ASSERT.md files for exact expectations.

## Verify command

```sh
doctest vet ./tests/scaff-cli
doctest test -v ./tests/scaff-cli
```

All 25 tests must pass (GREEN). Do not modify sealed tests.